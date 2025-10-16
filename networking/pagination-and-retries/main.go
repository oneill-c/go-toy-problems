// Time & Retries
//
// Implements an exponential backoff retry mechanism with ±25% jitter.
// Each retry doubles the delay up to a maximum cap, then applies random
// jitter to prevent synchronized retry storms. Demonstrates use of Go's
// time.Duration, Sleep, and randomization for resilience and rate control.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"
)

const (
	maxRetries    = 3
	jitterDivider = 4 // 25% of base
	base          = 150 * time.Millisecond
	max           = 2 * time.Second
	globalTimeout = 1200 * time.Millisecond
)

// ----------------------- HTTP -----------------------
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Page struct {
	Data []Item `json:"data"`
	Next string `json:"next"` // "" means no more pages
}

func isTransient(status int, err error) bool {
	if err != nil {
		return true
	}
	if status == http.StatusTooManyRequests || status == http.StatusRequestTimeout {
		return true
	}
	return status >= 500 && status < 599
}

func fetchPage(ctx context.Context, url string) (Page, int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Page{}, 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Page{}, 0, err
	}
	defer resp.Body.Close()

	var p Page
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
			return Page{}, 0, err
		}
		return p, resp.StatusCode, nil
	}
	return Page{}, resp.StatusCode, fmt.Errorf("status %d", resp.StatusCode)
}

func FetchAll(ctx context.Context, serverURL string, startingEndpoint string) ([]Item, error) {
	var out []Item
	url := startingEndpoint

	fmt.Println(url)
	for url != "" {
		var pg Page
		var status int

		err := RetryCtx(ctx, maxRetries, base, max, func(ctx context.Context) error {
			p, s, err := fetchPage(ctx, fmt.Sprintf("%s%s", serverURL, url))
			pg, status = p, s
			if err != nil && !isTransient(s, err) {
				return nil
			}
			return err
		})
		if err != nil {
			return nil, err
		}
		if pg.Data == nil && !isTransient(status, nil) && status >= 400 && status < 500 {
			return nil, fmt.Errorf("terminal client error: %d", status)
		}

		out = append(out, pg.Data...)
		fmt.Printf("fetched %d items (total=%d), next=%q\n", len(pg.Data), len(out), pg.Next)
		url = pg.Next
	}
	return out, nil
}

// ----------------------- Mock HTTP Server -----------------------
func NewMockPaginatedServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")

		// Simulate an intermittent failure on page 2
		if page == "2" && time.Now().UnixNano()%2 == 0 {
			http.Error(w, "temporary server error", http.StatusInternalServerError)
			return
		}

		// Data for each page
		pages := map[string][]Item{
			"1": {{ID: 1, Name: "A"}, {ID: 2, Name: "B"}, {ID: 3, Name: "C"}},
			"2": {{ID: 4, Name: "D"}, {ID: 5, Name: "E"}},
			"3": {{ID: 6, Name: "F"}},
		}

		var next string
		switch page {
		case "1":
			next = "/api/items?page=2"
		case "2":
			next = "/api/items?page=3"
		default:
			next = ""
		}

		data := Page{
			Data: pages[page],
			Next: next,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})
	// start server
	return httptest.NewServer(mux)
}

// ----------------------- Timeout and retries -----------------------
func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}

func jitter(d time.Duration) time.Duration {
	// compute a 25% jitter band
	n := int64(d / jitterDivider)

	// safety clamp 1 - skip jittering if too small
	if n <= 0 {
		return d
	}

	// apply symmetric random jitter
	neg := rand.Int63n(n) // [0, n]
	pos := rand.Int63n(n) // [0, n]
	x := d - time.Duration(neg) + time.Duration(pos)

	// safety clamp 2 - return original d if x < 0
	if x < 0 {
		return d
	}
	return x
}

func RetryCtx(ctx context.Context, maxRetries int, base, max time.Duration, op func(context.Context) error) error {
	var err error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		if err = op(ctx); err == nil {
			return nil
		}

		if attempt == maxRetries {
			return fmt.Errorf("max retries of %d reached %w", maxRetries, err)
		}
		delay := min(max, base<<attempt)
		sleep := jitter(delay)
		if sleep > max {
			sleep = max
		}
		t := time.NewTimer(sleep)
		fmt.Printf("attempt=%d error=%v sleeping=%v\n", attempt+1, err, sleep)
		select {
		case <-ctx.Done():
			t.Stop()
			return ctx.Err()
		case <-t.C: // blocks for duration of sleep value, then drops through to next iteration
		}
	}
	return err
}

// ----------------------- Main -----------------------
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), globalTimeout)
	defer cancel()

	srv := NewMockPaginatedServer()
	defer srv.Close()

	page1 := "/api/items?page=1"
	items, err := FetchAll(ctx, srv.URL, page1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("done:", len(items), "items")
}

// 🔐 Exercise: HMAC-Verified JSON Fetch with Retry and Backoff

// In this exercise, you’ll implement a resilient HTTP client in Go that performs the following:
// 	1.	Fetches a JSON payload from an HTTP endpoint using a GET request.
// 	2.	Reads the response body as raw bytes to:
// 		•	Verify the payload’s authenticity using an HMAC-SHA256 signature and a pre-shared secret.
// 		•	Abort if the HMAC does not match.
// 	3.	Parses the verified payload into a typed Go struct.
// 	4.	Retries on network errors or non-2xx responses using:
// 		•	Exponential backoff with jitter.
// 		•	A maximum of 5 retries.
// 		•	A global request timeout of 10 seconds.

// Example payload
// {
// 	"data": {
//  	"event_id": "abc123",
//   	"timestamp": "2025-10-06T15:04:05Z",
//   	"user_id": "user_456",
//   	"action": "login"
// 	}
// }

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type Event struct {
	EventID   string `json:"event_id"`
	Timestamp string `json:"timestamp"`
	UserID    string `json:"user_id"`
	Action    string `json:"action"`
}

type EventPayload struct {
	Data Event `json:"data"`
	Signature string `json:"secret"`
}

// Likely fed from the env, stored in some secure storage and read via the CI pipeline
const (
	maxRetries 				= 5
	requestTimeout 		= 10 * time.Second
	secret						= "some-secret-value"
	baseJitter				= 500 * time.Millisecond
)

func fetchEvents(client *http.Client, parnerUrl string, retries int) (*Event, error) {
	// Iterate until maxRetries
	for i := 0; i < retries; i++ {

		//  call API
		resp, err := client.Get(parnerUrl)
		if err != nil {
			return nil, errors.New("fetchEvents: could not retrieve events")
		}

		//  Handle 4xx (429 separately - if rate limited), handle 5xx
		if resp.StatusCode >= 400 {
			// specific 429 case
			if resp.StatusCode == http.StatusTooManyRequests {
				delay := jitterBackoff(baseJitter, i)
				fmt.Printf("waiting %v before next retry", delay)
				time.Sleep(delay)
				continue
			}
			resp.Body.Close()
			return nil, errors.New("fetchEvents: bad request")
		}

		// 	Read the body as raw bytes
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, errors.New("fetchEvents: could not decode body")
		}

		// extract raw fields
		var raw map[string]json.RawMessage
		if err := json.Unmarshal(body, &raw); err != nil {
			return nil, errors.New("fetchEvents: invalid JSON")
		}

		sigBytes, ok := raw["signature"]
		if !ok {
			return nil, errors.New("fetchEvents: missing signature")
		}

		dataBytes, ok := raw["data"]
		if !ok {
			return nil, errors.New("fetchEvents: missing data")
		}

		var sig string
		if err := json.Unmarshal(sigBytes, &sig); err != nil {
			return nil, errors.New("fetchEvents: invalid signature")
		}

		if !verifyHmac(dataBytes, []byte(secret), sig) {
			return nil, errors.New("fetchEvents: could not verify signature")
		}

		var event Event
		if err := json.Unmarshal(dataBytes, &event); err != nil {
			return nil, errors.New("fetchEvents: could not decode data")
		}
		return &event, nil
	}
	return nil, errors.New("fetchEvents: reties exhausted")
}

func verifyHmac(body []byte, secret []byte, sig string) bool {
	mac := hmac.New(sha256.New, secret)
	mac.Write(body)
	expected := mac.Sum(nil)
	expectedHex := hex.EncodeToString(expected)
	return hmac.Equal([]byte(sig), []byte(expectedHex))
}

func jitterBackoff(baseJ time.Duration, attempt int) time.Duration {
	max := baseJ * (1 << (attempt - 1))
	return time.Duration(rand.Int63n(int64(max)))
}

func main() {
	partnerUrl := "https://partner.example.com/events"
	client := &http.Client{ Timeout: requestTimeout }
	event, err := fetchEvents(client, partnerUrl, maxRetries)
	if err != nil {
		fmt.Printf("main: fetchEvents returned error: %v", err)
		panic(err)
	}

	fmt.Println(event)
}

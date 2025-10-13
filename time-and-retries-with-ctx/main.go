// Time & Retries
//
// Implements an exponential backoff retry mechanism with ±25% jitter.
// Each retry doubles the delay up to a maximum cap, then applies random
// jitter to prevent synchronized retry storms. Demonstrates use of Go's
// time.Duration, Sleep, and randomization for resilience and rate control.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	maxRetries             = 5
	jitterDivider          = 4 // 25% of base
	base                   = 150 * time.Millisecond
	max                    = 2 * time.Second
	perRequestFailDuration = 800 * time.Millisecond
)

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

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Millisecond)
	defer cancel()

	i := 0
	op := func(ctx context.Context) error {
		i++

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(perRequestFailDuration):
			// pretend we made a network call
		}
		if i <= 2 {
			return fmt.Errorf("transient failure on attempt %d", i)
		}
		return nil
	}

	err := RetryCtx(ctx, maxRetries, base, max, op)
	if err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("success")
	}
}

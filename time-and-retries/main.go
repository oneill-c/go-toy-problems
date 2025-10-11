// Time & Retries
//
// Implements an exponential backoff retry mechanism with ±25% jitter.
// Each retry doubles the delay up to a maximum cap, then applies random
// jitter to prevent synchronized retry storms. Demonstrates use of Go's
// time.Duration, Sleep, and randomization for resilience and rate control.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxRetries    = 10
	jitterDivider = 4 // 25% of base
	base          = 150 * time.Millisecond
	max           = 2 * time.Second
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

func Retry(maxRetries int, base, max time.Duration, op func() error) error {
	for attempt := 0; attempt <= maxRetries; attempt++ {
		err := op()
		if err == nil {
			return nil
		}

		if attempt == maxRetries {
			return fmt.Errorf("max retries of %d reached", maxRetries)
		}
		delay := min(max, base<<attempt)
		sleep := jitter(delay)
		if sleep > max {
			sleep = max
		}
		fmt.Printf("attempt=%d error=%v sleeping=%v\n", attempt+1, err, sleep)
		time.Sleep(sleep)
	}
	return nil
}

func main() {
	i := 0
	op := func() error {
		i++
		if i < 3 {
			return fmt.Errorf("attempt %d failed with error", i)
		}
		return nil
	}

	err := Retry(maxRetries, base, max, op)
	if err != nil {
		fmt.Println("error on retry", err)
	} else {
		fmt.Println("success")
	}

	j := 0
	op2 := func() error {
		j++
		return fmt.Errorf("attempt %d failed with error", j)
	}

	err1 := Retry(maxRetries, base, max, op2)
	if err1 != nil {
		fmt.Println("error on retry", err1)
	} else {
		fmt.Println("success")
	}
}

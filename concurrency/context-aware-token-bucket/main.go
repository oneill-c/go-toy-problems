package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	workerCount       = 5
	totalJobsCount    = 12
	tokensPerSecond   = 5
	maxBucketCapacity = 10
	jobBackpressure   = 8
	globalCtxTimeout  = 1 * time.Second
	jobSleepTime      = 200 * time.Millisecond
	// Adjustable dials (comment out above and put these values in instead)
	// globalCtxTimeout = 100 * time.Millisecond // To see less jobs completed
	// jobSleepTime     = 400 * time.Millisecond // Jobs take longer to complete
)

// ---------------------------- Job Store ----------------------------
type JobStore struct {
	data []int
	mu   sync.Mutex
}

func (js *JobStore) Set(v int) {
	js.mu.Lock()
	defer js.mu.Unlock()
	js.data = append(js.data, v)
}

func (js *JobStore) Snapshot() []int {
	js.mu.Lock()
	defer js.mu.Unlock()
	out := make([]int, len(js.data))
	copy(out, js.data)
	return out
}

// ---------------------------- Rate Limiter ----------------------------
type limiter struct {
	rate   int
	burst  int
	tokens chan struct{}
}

type Limiter interface {
	Acquire(ctx context.Context) error
}

func NewLimiter(ctx context.Context, rate, burst int) *limiter {
	// instantiate limiter with token channel with a buffer of 'burst'
	l := &limiter{
		tokens: make(chan struct{}, burst),
		rate:   rate,
		burst:  burst,
	}

	// start ticker, with a token refill event 5 (rate) times per second (every 200ms)
	ticker := time.NewTicker(time.Second / time.Duration(rate))

	// exectute on separate goroutine
	go func() {
		// defer Stop() immediately after creating
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				select {
				case l.tokens <- struct{}{}:
				// default case ensures max is honored as we fall through to it when l.tokens is at capacity (burst)
				default:
				}
			}
		}
	}()

	return l
}

func (l *limiter) Acquire(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-l.tokens:
		return nil
	}
}

// ---------------------------- Worker Pool ----------------------------
func StartWorkerPool(ctx context.Context, jobs <-chan int, workerCount int) []int {
	var wg sync.WaitGroup
	wg.Add(workerCount)

	js := &JobStore{
		data: make([]int, 0, totalJobsCount),
	}

	lim := NewLimiter(ctx, tokensPerSecond, maxBucketCapacity)

	for i := 0; i < workerCount; i++ {
		go func(id int) {
			defer wg.Done()

			for {
				if err := lim.Acquire(ctx); err != nil {
					return
				}
				select {
				case <-ctx.Done():
					fmt.Println("context canceled: ", ctx.Err())
					return
				case j, ok := <-jobs:
					if !ok {
						return
					}
					time.Sleep(jobSleepTime)
					out := j * 2
					js.Set(out)
					fmt.Printf("worker %d processed %d\n", id, j)
				}
			}
		}(i)
	}

	wg.Wait()
	return js.Snapshot()
}

// ---------------------------- Main ----------------------------
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), globalCtxTimeout)
	defer cancel()

	jobs := make(chan int, jobBackpressure)

	go func() {
		for i := 1; i <= totalJobsCount; i++ {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			case jobs <- i:
			}
		}
		close(jobs)
	}()

	results := StartWorkerPool(ctx, jobs, workerCount)
	fmt.Printf("collected %d results before cancel: %v\n", len(results), results)
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

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

const (
	workerCount      = 5
	totalJobsCount   = 12
	globalCtxTimeout = 1 * time.Second
	jobSleepTime     = 200 * time.Millisecond
	// Adjustable dials (comment out above and put these values in instead)
	// globalCtxTimeout = 100 * time.Millisecond // To see less jobs completed
	// jobSleepTime     = 400 * time.Millisecond // Jobs take longer to complete
)

func StartWorkerPool(ctx context.Context, jobs <-chan int, workerCount int) []int {
	var wg sync.WaitGroup
	wg.Add(workerCount)

	js := &JobStore{
		data: make([]int, 0, totalJobsCount),
	}

	for i := 0; i < workerCount; i++ {
		go func(id int) {
			defer wg.Done()

			for {
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

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), globalCtxTimeout)
	defer cancel()

	jobs := make(chan int)

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

package concurrentpipeline

import (
	"context"
)

// Producer
func generator(ctx context.Context, nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}()
	return out
}

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * n:
			}
		}
	}()
	return out
}

func addOne(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n + 1:
			}
		}
	}()
	return out
}

func filterEven(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 != 0 {
				if ctx.Err() != nil {
					return
				}
				continue
			}
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

func collect(ctx context.Context, in <-chan int) []int {
	var out []int
	for {
		select {
		case <-ctx.Done():
			return out
		case n, ok := <-in:
			if !ok {
				return out
			}
			out = append(out, n)
		}
	}
}

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	start := generator(ctx, []int{1, 2, 3, 4})
// 	stage1 := square(ctx, start)
// 	stage2 := addOne(ctx, stage1)
// 	stage3 := filterEven(ctx, stage2)
// 	result := collect(ctx, stage3)
// 	fmt.Println("Pipeline output: ", result)
// }

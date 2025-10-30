package concurrentpipeline

import (
	"context"
	"reflect"
	"testing"
)

func TestConcurrentPipeline_(t *testing.T) {

	got := collect(context.Background(),
		filterEven(context.Background(),
			addOne(context.Background(),
				square(context.Background(),
					generator(context.Background(), []int{1, 2, 3, 4})))))

	want := []int{2, 10}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

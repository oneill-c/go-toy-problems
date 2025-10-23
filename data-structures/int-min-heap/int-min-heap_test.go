package intminheap

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestIntMinHeap_Order(t *testing.T) {
	h := &IntMinHeap{5,1,4}
	heap.Init(h)
	heap.Push(h, 3)
	var got []int
	for h.Len() > 0 {
		got = append(got, heap.Pop(h).(int))
	}
	want := []int{1,3,4,5}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestTopLargest(t *testing.T) {
	nums := []int{5,1,9,3,12,7,2}
	got := TopKLargest(nums, 3)
	want := []int{7,9,12}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
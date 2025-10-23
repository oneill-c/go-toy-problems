package intminheap

import "container/heap"

type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TopKLargest(nums []int, k int) []int {
	if k <= 0 {
		return nil
	}
	h := &IntMinHeap{}
	heap.Init(h)
	for _, x := range nums {
		if h.Len() < k {
			heap.Push(h, x)
		} else if x > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, x)
		}
	}
	out := make([]int, 0, h.Len())
	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(int))
	}
	return out
}
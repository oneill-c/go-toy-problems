package main

import (
	"container/heap"
	"fmt"
)

type User struct {
	ID int
	Name string
	Score int
}

type MergeItem struct {
	User User
	ListID int
	Index int
}

type MergeHeap []MergeItem

func (h MergeHeap) Len() int { return len(h) }
func (h MergeHeap) Less(i, j int) bool { return h[i].User.Score > h[j].User.Score }
func (h MergeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MergeHeap) Push(x any) { *h = append(*h, x.(MergeItem))}
func (h *MergeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func MergAndDedupTopK(sources [][]User, k int) []User {
	h := &MergeHeap{}
	heap.Init(h)
	
	// seed heap with the first element from each list
	for i, list := range sources {
		if len(list) > 0 {
			heap.Push(h, MergeItem{User: list[0], ListID: i, Index: 0})
		}
	}

	seen := make(map[int]bool)
	var result []User

	for h.Len() > 0 && len(result) < k {
		// Repeatidly pop next best item
		item := heap.Pop(h).(MergeItem)

		// Dedupe
		u := item.User
		if !seen[u.ID] {
			seen[u.ID] = true
			result = append(result, u)
		}

		// Push 
		nextIdx := item.Index + 1
		if nextIdx < len(sources[item.ListID]) {
			nextUser := sources[item.ListID][nextIdx]
			heap.Push(h, MergeItem{User: nextUser, ListID: item.ListID, Index: nextIdx})
		}
	}
	return result
 }

func main() {
	listA := []User{{1,"A",90},{2,"B",85},{3,"C",80}}
	listB := []User{{3,"C",95},{4,"D",88},{7,"G",75}}
	listC := []User{{6,"F",99},{2,"B",83},{7,"G",75}}
	merged := MergAndDedupTopK([][]User{listA, listB,listC}, 5)
	fmt.Println("Top 5 merged and deduped:")
	for _, u := range merged {
		fmt.Printf("%s (%d)\n", u.Name, u.Score)
	}
}

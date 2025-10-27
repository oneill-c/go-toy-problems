package topkdedupewithsort

import (
	"container/heap"
	"slices"
	"sort"
)

// User is the record we’re ranking by Score.
type User struct {
	ID    int
	Name  string
	Score int
}

// ---------------- Min-Heap for Users (by Score, tie-break ID) ----------------

// UserMinHeap maintains the *smallest* score at index 0.
// Tie-break: smaller ID is considered "less" to keep ordering deterministic.
type UserMinHeap []User

// TODO: implement heap.Interface methods:
//
// func (h UserMinHeap) Len() int            { /* TODO */ }
// func (h UserMinHeap) Less(i, j int) bool  { /* TODO: return h[i].Score < h[j].Score || (scores equal && h[i].ID < h[j].ID) */ }
// func (h UserMinHeap) Swap(i, j int)       { /* TODO */ }
// func (h *UserMinHeap) Push(x any)         { /* TODO: append x.(User) */ }
// func (h *UserMinHeap) Pop() any           { /* TODO: remove and return last element */ }

// Optional: Peek helper (not required by heap.Interface).
// func (h UserMinHeap) Peek() User { return h[0] }

func (h UserMinHeap) Len() int { return len(h) }
func (h UserMinHeap) Less(i, j int) bool {
	if h[i].Score != h[j].Score {
		return h[i].Score < h[j].Score // min-heap by score
	}
	return h[i].ID > h[j].ID // tie tweak so Reverse() yields ID asc
}
func (h UserMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *UserMinHeap) Push(x any)   { *h = append(*h, x.(User)) }
func (h *UserMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ---------------- Top-K (O(N log K)) ----------------

// TopKUsers returns the K highest-scoring unique users (by ID).
// Contract:
//   - If k <= 0: return nil.
//   - If k >= len(users): return users sorted descending by Score (tie-break ID).
//   - Duplicates by ID should be ignored (keep any one; or keep highest if you choose).
func TopKUsers(users []User, k int) []User {
	if k <= 0 {
		return nil
	}
	if len(users) == 0 {
		return nil
	}

	// OPTIONAL (recommended): deduplicate by ID up front (keep highest score).
	// This keeps the heap smaller and results stable.
	// TODO (optional):
	best := make(map[int]User, len(users))
	for _, u := range users {
		if curr, ok := best[u.ID]; !ok || u.Score > curr.Score || (u.Score == curr.Score && u.ID < curr.ID) {
			best[u.ID] = u
		}
	}
	uniq := make([]User, 0, len(best))
	for _, u := range best {
		uniq = append(uniq, u)
	}
	users = uniq

	if k >= len(users) {
		// Full sort fallback (descending by Score, tie-break ID asc)
		out := append([]User(nil), users...)
		sort.Slice(out, func(i, j int) bool {
			if out[i].Score != out[j].Score {
				return out[i].Score > out[j].Score
			}
			return out[i].ID < out[j].ID
		})
		return out
	}

	h := &UserMinHeap{}
	heap.Init(h)

	// Stream through users maintaining a bounded min-heap of size k.
	for _, u := range users {
		if h.Len() < k {
			heap.Push(h, u)
			continue
		}
		// Compare with the smallest among current top-K.
		min := (*h)[0]
		if u.Score > min.Score || (u.Score == min.Score && u.ID < min.ID) {
			heap.Pop(h)
			heap.Push(h, u)
		}
	}

	// Extract K users (ascending by score because it's a min-heap).
	out := make([]User, 0, h.Len())
	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(User))
	}

	// Reverse to descending, or just sort K items for deterministic order.
	slices.Reverse(out)
	return out
}

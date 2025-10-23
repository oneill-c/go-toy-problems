package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// ------------- Shared -------------
type User struct {
	ID int
	Name string
	Score int
}

func generateUsers(r *rand.Rand, n int) []User {
	users := make([]User, n)
	for i := range users {
		users[i] = User{
			ID: i + 1,
			Name: fmt.Sprintf("User-%d", i+1),
			Score: r.Intn(100),
		}
	}
	return users
}

// ------------- Global Sort (sort.Slice(..)) -------------
func TopKSort(users []User, k int) []User {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Score > users[j].Score
	})
	return users[:k]
}

// ------------- Incremental Sort (MinHeap) -------------
type UserMinHeap []User

func (h UserMinHeap) Len() int { return len(h)}
func (h UserMinHeap) Less(i, j int) bool { return h[i].Score < h[j].Score }
func (h UserMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *UserMinHeap) Push(x any) { *h = append(*h, x.(User))}
func (h *UserMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TopKHeap(users []User, k int) []User {
	h := &UserMinHeap{}
	heap.Init(h)
	for _, u := range users {
		if h.Len() < k {
			heap.Push(h, u)
		} else if u.Score > (*h)[0].Score {
			heap.Pop(h)
			heap.Push(h, u)
		}
	}
	out := make([]User, h.Len())
	for i := len(out)-1; i >= 0; i-- {
		out[i] = heap.Pop(h).(User)
	}
	return out
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	all := generateUsers(r, 10)

	fmt.Println("All users")
	for _, u := range all {
		fmt.Printf("%s => %d\n", u.Name, u.Score)
	}

	k := 3
	fmt.Println("\nTop K (Sort):")
	for _, u := range TopKSort(all, k) {
		fmt.Printf("%s => %d\n", u.Name, u.Score)
	}

	fmt.Println("\nTop K (heap):")
	for _, u := range TopKHeap(all, k) {
		fmt.Printf("%s => %d\n", u.Name, u.Score)
	}
}


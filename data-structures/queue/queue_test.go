package datastructures

import "testing"

func TestQueue_EnqueueDrain(t *testing.T) {
	var q Queue[int]

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Drain queue
	if front, ok := q.Dequeue(); !ok || front != 10 {
		t.Fatalf("Dequeue got (%v, %v), want (10, true)", front, ok)
	}
	if front, ok := q.Dequeue(); !ok || front != 20 {
		t.Fatalf("Dequeue got (%v, %v), want (20, true)", front, ok)
	}
	if front, ok := q.Dequeue(); !ok || front != 30 {
		t.Fatalf("Dequeue got (%v, %v), want (30, true)", front, ok)
	}

	if !q.IsEmpty() {
		t.Fatalf("expected queue to be empty")
	}
}

func TestQueue_EnqueueMoreBeforeDrain(t *testing.T) {
	var q Queue[string]

	// Enqueue some
	q.Enqueue("A")
	q.Enqueue("B")
	q.Enqueue("C")

	// Deque some
	if front, ok := q.Dequeue(); !ok || front != "A" {
		t.Fatalf("Dequeue got (%q, %v), want (\"A\",true)", front, ok)
	}
	if front, ok := q.Dequeue(); !ok || front != "B" {
		t.Fatalf("Dequeue got (%q, %v), want (\"B\",true)", front, ok)
	}

	// Enqueue some more
	q.Enqueue("D")
	q.Enqueue("E")

	// Ensure "C" is still there and will now get dequeued
	if front, ok := q.Dequeue(); !ok || front != "C" {
		t.Fatalf("Dequeue got (%q,%v), want (\"C\",true)", front, ok)
	}

	// Drain the rest
	if front, ok := q.Dequeue(); !ok || front != "D" {
		t.Fatalf("Dequeue got (%q,%v), want (\"D\", true)", front, ok)
	}
	if front, ok := q.Dequeue(); !ok || front != "E" {
		t.Fatalf("Dequeue got (%q,%v), want (\"E\", true)", front, ok)
	}

	if !q.IsEmpty() {
		t.Fatalf("expected queue to be empty")
	}
}
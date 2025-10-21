package datastructures

import "testing"

func TestStack_PushPop(t *testing.T) {
	var s Stack[int]

	// Empty states
	if !s.IsEmpty() {
		t.Fatalf("expected empty stack at start")
	}
	if _, ok := s.Peek(); ok {
		t.Fatalf("expected Peek on empty to return ok=false")
	}
	if _, ok := s.Pop(); ok {
		t.Fatalf("expected Pop on empty to return ok=false")
	}

	// Push entries
	s.Push(10)
	s.Push(20)
	s.Push(30)

	// Peek
	if top, ok := s.Peek(); !ok || top != 30 {
		t.Fatalf("Peek got (%v, %v), want (30, true)", top, ok)
	}
	if s.IsEmpty() {
		t.Fatalf("expected non-empty after pushes")
	}

	// Pop
	if v, ok := s.Pop(); !ok || v != 30 {
		t.Fatalf("Pop got (%v, %v), want (30, true)", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != 20 {
		t.Fatalf("Pop got (%v, %v), want (20, true)", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != 10 {
		t.Fatalf("Pop got (%v, %v), want (10, true)", v, ok)
	}

	// Empty after all elements popped
	if !s.IsEmpty() {
		t.Fatalf("expected empty after popping all elements")
	}
}

func TestStack_PeekDoesNotRemove(t *testing.T) {
	var s Stack[string]
	s.Push("Hello")
	s.Push("World")

	if top, ok := s.Peek(); !ok || top != "World" {
		t.Fatalf("Peek got (%q, %v), want (\"World\", true)", top, ok)
	}

	// Peek again, should be the same
	if top, ok := s.Peek(); !ok || top != "World" {
		t.Fatalf("Peek got (%q, %v), want (\"World\", true)", top, ok)
	}

	// Pop "World"
	if top, ok := s.Pop(); !ok || top != "World" {
		t.Fatalf("Pop got (%q, %v), want (\"World\", true)", top, ok)
	}

	// Peek, should be "Hello" this time
	if top, ok := s.Peek(); !ok || top != "Hello" {
		t.Fatalf("Peek got (%q, %v), want (\"Hello\", true)", top, ok)
	}
}
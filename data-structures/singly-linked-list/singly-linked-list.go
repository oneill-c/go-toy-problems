package datastructures

import "fmt"

type SLNode struct {
	Value int
	Next  *SLNode
}

type LinkedList struct {
	Head *SLNode
	Tail *SLNode
}

// Add a node to the end of the list
func (l *LinkedList) Append(v int) {
	n := &SLNode{}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = n
}

// Add a node to the front of the list
func (l *LinkedList) Prepend(v int) {
	n := &SLNode{Value: v, Next: l.Head}
	l.Head = n
	if l.Tail == nil {
		l.Tail = n
	}
}

// remove the node matching the value v
func (l *LinkedList) Delete(v int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == v {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		return
	}
	prev := l.Head
	for curr := l.Head.Next; curr != nil; curr = curr.Next {
		if curr.Value == v {
			prev.Next = curr.Next
			if curr == l.Tail {
				l.Tail = prev
			}
			return
		}
		prev = curr
	}
}

func (l *LinkedList) Print() {
	for n := l.Head; n != nil; n = n.Next {
		fmt.Printf("%d ", n.Value)
	}
	fmt.Println()
}

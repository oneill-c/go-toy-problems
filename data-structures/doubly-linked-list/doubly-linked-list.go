package datastructures

import "fmt"

type DLNode struct {
	Value int
	Prev  *DLNode
	Next  *DLNode
}

type DoublyLinkedList struct {
	Head *DLNode
	Tail *DLNode
}

func (l *DoublyLinkedList) Append(v int) {
	n := &DLNode{Value: v}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	n.Prev = l.Tail
	l.Tail.Next = n
	l.Tail = n
}

func (l *DoublyLinkedList) Prepend(v int) {
	n := &DLNode{Value: v}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	n.Next = l.Head
	l.Head.Prev = n
	l.Head = n
}

func (l *DoublyLinkedList) Delete(v int) {
	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Value == v {
			if curr.Prev != nil {
				curr.Prev.Next = curr.Next
			} else {
				l.Head = curr.Next
			}
			if curr.Next != nil {
				curr.Next.Prev = curr.Prev
			} else {
				l.Tail = curr.Prev
			}
			return
		}
	}
}

func (l *DoublyLinkedList) PrintForward() {
	for n := l.Head; n != nil; n = n.Next {
		fmt.Printf("%d ", n.Value)
	}
	fmt.Println()
}

func (l *DoublyLinkedList) PrintBackward() {
	for n := l.Tail; n != nil; n = n.Prev {
		fmt.Printf("%d ", n.Value)
	}
	fmt.Println()
}

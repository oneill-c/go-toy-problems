package bst

type BstNode struct {
	Key int
	Left *BstNode
	Right *BstNode
}

type Tree struct {
	Root *BstNode
}

func (t *Tree) Search(key int) bool {
	for n := t.Root; n != nil; {
		switch {
		case key < n.Key:
			n = n.Left
		case key > n.Key:
			n = n.Right
		default:
			return true
		}
	}
	return false
}

func (t *Tree) Insert(key int) {
	if t.Root == nil {
		t.Root = &BstNode{Key: key}
		return
	}
	n := t.Root
	for {
		if key < n.Key {
			if n.Left == nil {
				n.Left = &BstNode{Key: key}
				return
			}
			n = n.Left
		} else if key > n.Key {
			if n.Right == nil {
				n.Right = &BstNode{Key: key}
				return
			}
			n = n.Right
		} else {
			return
		}
	}
}

func (t *Tree) Delete(key int) {
	t.Root = deleteNode(t.Root, key)
}

func deleteNode(n *BstNode, key int) *BstNode {
	if n == nil {
		return nil
	}
	if key < n.Key {
		n.Left = deleteNode(n.Left, key)
		return n
	}
	if key > n.Key {
		n.Right = deleteNode(n.Right, key)
		return n
	}

	if n.Left == nil {
		return n.Right
	}
	if n.Right == nil {
		return n.Left
	}

	succ := n.Right
	for succ.Left != nil {
		succ = succ.Left
	}
	n.Key = succ.Key
	n.Right = deleteNode(n.Right, succ.Key)
	return n
}
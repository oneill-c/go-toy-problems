package main

import "fmt"

type Node struct {
	Val int
	Left *Node
	Right *Node
}

func DFS(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	DFS(root.Left)
	DFS(root.Right)
}

func main() {

	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}

	DFS(root)
}
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Append(l *ListNode, val int) *ListNode {
	if l == nil {
		return &ListNode{
			Val:  val,
			Next: nil,
		}
	}

	base := l
	current := l
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{
		Val:  val,
		Next: nil,
	}

	return base
}

func (l *ListNode) Display() {
	n := l
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

type LinkedList struct {
	head *ListNode
}

// Prepend
// delete

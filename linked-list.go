package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func (l *LinkedList) Append(value int) {
	if l.head == nil {
		l.head = &ListNode{
			Val:  value,
			Next: nil,
		}
		return
	}

	current := l.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &ListNode{
		Val:  value,
		Next: nil,
	}
	return
}

// Prepend
// delete

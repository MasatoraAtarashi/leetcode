package main

import (
	"fmt"
)

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

// Prepend
func Prepend(l *ListNode, val int) *ListNode {
	if l == nil {
		return &ListNode{
			Val:  val,
			Next: nil,
		}
	}

	tmp := &ListNode{
		Val:  l.Val,
		Next: l.Next,
	}
	l.Val = val
	l.Next = tmp
	return l
}

func (l *ListNode) Display() {
	n := l
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func (l *ListNode) String() string {
	s := ""

	n := l
	for n != nil {
		s += fmt.Sprintf("%d -> ", n.Val)
		n = n.Next
	}

	return s
}

func Pop(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return nil
	}

	pre := l
	tmp := l.Next
	for tmp.Next != nil {
		pre = pre.Next
		tmp = tmp.Next
	}

	pre.Next = nil
	return l
}

func Reverse(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}

	var prev *ListNode
	current := &ListNode{
		Val:  l.Val,
		Next: nil,
	}
	next := l.Next

	for current != nil {
		// fmt.Printf("l: %s, prev: %s, current: %s, next: %s\n", l.String(), prev.String(), current.String(), next.String())
		current.Next = prev
		prev = current
		current = next
		// fmt.Printf("l: %s, prev: %s, current: %s, next: %s\n", l.String(), prev.String(), current.String(), next.String())
		if current != nil {
			next = current.Next
		}
	}

	return prev
}

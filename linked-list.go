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
		current.Next = prev
		prev = current
		current = next
		if current != nil {
			next = current.Next
		}
	}

	return prev
}

func Length(l *ListNode) int {
	length := 0
	current := l
	for current != nil {
		length++
		current = current.Next
	}

	return length
}

func Remove(l *ListNode, index int) *ListNode {
	if l == nil {
		return nil
	}
	if index == 0 {
		if l.Next == nil {
			return nil
		} else {
			return &ListNode{
				Val:  l.Next.Val,
				Next: l.Next.Next,
			}
		}
	}
	if index > Length(l)-1 {
		return l
	}

	prev := l
	current := l.Next
	next := current.Next
	for i := 1; i < index; i++ {
		prev = current
		current = next
		next = current.Next
	}

	prev.Next = next
	return l
}

// TODO: Pop-First, Get, Set, Insert

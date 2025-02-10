package main

import "fmt"

type DoublyLinkedList[T comparable] struct {
	Head   *DLLNode[T]
	Tail   *DLLNode[T]
	length int
}

type DLLNode[T comparable] struct {
	Val  T
	Next *DLLNode[T]
	Prev *DLLNode[T]
}

func (d *DoublyLinkedList[T]) Length() int {
	return d.length
}

func (d *DoublyLinkedList[T]) Append(val T) {
	if d.Head == nil {
		node := &DLLNode[T]{
			Val:  val,
			Next: nil,
			Prev: nil,
		}
		d.Head = node
		d.Tail = node
		d.length = 1
		return
	}

	newNode := &DLLNode[T]{
		Val:  val,
		Next: nil,
		Prev: d.Tail,
	}
	d.Tail.Next = newNode
	d.Tail = newNode
	d.length++
}

func (d *DoublyLinkedList[T]) Prepend(val T) {
	if d.Head == nil {
		node := &DLLNode[T]{
			Val:  val,
			Next: nil,
			Prev: nil,
		}
		d.Head = node
		d.Tail = node
		d.length = 1
		return
	}

	newNode := &DLLNode[T]{
		Val:  val,
		Next: d.Head,
		Prev: nil,
	}
	d.Head.Prev = newNode
	d.Head = newNode
	d.length++
}

func (d *DoublyLinkedList[T]) Pop() T {
	if d.Tail == nil {
		var zero T
		return zero
	}

	val := d.Tail.Val
	d.Tail = d.Tail.Prev
	if d.Tail != nil {
		d.Tail.Next = nil
	} else {
		d.Head = d.Tail
	}
	d.length--
	return val
}

func (d *DoublyLinkedList[T]) RemoveVal(val T) bool {
	tmp := d.Head
	for tmp != nil {
		if tmp.Val == val {
			if tmp.Next != nil {
				tmp.Next.Prev = tmp.Prev
			}
			if tmp.Prev != nil {
				tmp.Prev.Next = tmp.Next
			}
			if d.Head == tmp {
				d.Head = tmp.Next
			}
			if d.Tail == tmp {
				d.Tail = tmp.Prev
			}
			tmp.Next = nil
			tmp.Prev = nil
			d.length--
			return true
		}
		tmp = tmp.Next
	}

	return false
}

func (d *DoublyLinkedList[T]) String() string {
	if d.Head == nil {
		return "Empty list"
	}

	var result string
	for node := d.Head; node != nil; node = node.Next {
		result += fmt.Sprintf("%v", node.Val)
		if node.Next != nil {
			result += " <-> "
		}
	}
	return result
}

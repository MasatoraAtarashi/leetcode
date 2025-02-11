package main

type Queue[T comparable] struct {
	linkedList *QueueLinkedList[T]
}

type QueueLinkedList[T comparable] struct {
	First  *QueueNode[T]
	Last   *QueueNode[T]
	length int
}

type QueueNode[T comparable] struct {
	Val  T
	Next *QueueNode[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(val T) {
	if q.linkedList == nil || q.Length() == 0 {
		n := &QueueNode[T]{
			Val:  val,
			Next: nil,
		}
		q.linkedList = &QueueLinkedList[T]{
			First:  n,
			Last:   n,
			length: 1,
		}
		return
	}

	n := &QueueNode[T]{
		Val:  val,
		Next: nil,
	}
	q.linkedList.Last.Next = n
	q.linkedList.Last = n
	q.linkedList.length++
}

func (q *Queue[T]) Dequeue() *QueueNode[T] {
	if q.linkedList == nil || q.linkedList.length == 0 {
		return nil
	}
	if q.linkedList.length == 1 {
		n := q.linkedList.First
		q.linkedList.First = nil
		q.linkedList.Last = nil
		q.linkedList.length = 0
		return n
	}
	n := q.linkedList.First
	q.linkedList.First = q.linkedList.First.Next
	q.linkedList.length--

	return n
}

func (q *Queue[T]) Length() int {
	if q.linkedList == nil {
		return 0
	}
	return q.linkedList.length
}

package main

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Top() T {
	if len(*s) == 0 {
		var z T
		return z
	}

	return (*s)[len(*s)-1]
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}

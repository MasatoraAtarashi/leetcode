package main

import (
	"strconv"
)

const OperatorPlus = "+"
const OperatorMinus = "-"
const OperatorMulti = "*"
const OperatorDivider = "/"

func evalRPN(tokens []string) int {
	s := Stack[int]{}
	for _, t := range tokens {
		var val int
		switch t {
		case OperatorPlus:
			t1 := s.Pop()
			t2 := s.Pop()
			val = t1 + t2
		case OperatorMinus:
			t1 := s.Pop()
			t2 := s.Pop()
			val = t2 - t1
		case OperatorMulti:
			t1 := s.Pop()
			t2 := s.Pop()
			val = t1 * t2
		case OperatorDivider:
			t1 := s.Pop()
			t2 := s.Pop()
			val = t2 / t1
		default:
			i, _ := strconv.Atoi(t)
			s.Push(i)
			continue
		}

		s.Push(val)
		// fmt.Println(s)
	}

	return s.Pop()
}

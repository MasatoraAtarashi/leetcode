package main

import "fmt"

func isValid(s string) bool {
	bucketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack Stack[rune]

	for i := range s {
		if open, isClosed := bucketMap[rune(s[i])]; isClosed {
			if stack.Top() != open {
				return false
			}
			stack.Pop()
		} else {
			stack.Push(rune(s[i]))
		}

		fmt.Println(stack)
	}

	return stack.Empty()
}

func main() {
	isValid("]")
}

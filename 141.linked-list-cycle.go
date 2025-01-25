package main

import "fmt"

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil {
		fmt.Println("slow", slow)
		fmt.Println("fast", fast)

		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}

		if fast == slow {
			return true
		}
		fmt.Println("")
	}

	return false
}

package main

import "fmt"

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 中間ノード見つける
	slow := head
	fast := head
	for fast.Next != nil {
		fmt.Printf("slow: %s\n", slow.String())
		fmt.Printf("fast: %s\n", fast.String())

		slow = slow.Next
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
		fmt.Printf("slow: %s\n", slow.String())
		fmt.Printf("fast: %s\n", fast.String())
	}
	mid := slow.Next
	fmt.Printf("mid: %s\n", mid.String())
	fmt.Printf("head: %s\n", head.String())

	// 2つのリストにして後半をreverse
	slow.Next = nil
	reversedAfter := Reverse(mid)
	fmt.Printf("head: %s\n", head.String())
	fmt.Printf("reversedAfter: %s\n", reversedAfter.String())

	// 順番に新しいリストに追加していく
	originalHead := head
	var result *ListNode
	for head != nil || reversedAfter != nil {
		if head != nil {
			result = Append(result, head.Val)
			head = head.Next
		}
		if reversedAfter != nil {
			result = Append(result, reversedAfter.Val)
			reversedAfter = reversedAfter.Next
		}
	}
	fmt.Printf("result: %s\n", result.String())
	originalHead.Val = result.Val
	originalHead.Next = result.Next
}

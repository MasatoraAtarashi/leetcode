package main

func reverseList(head *ListNode) *ListNode {
	// 反復
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev

		prev = current
		current = next
		if next != nil {
			next = next.Next
		}
	}

	head = prev
	return head
}

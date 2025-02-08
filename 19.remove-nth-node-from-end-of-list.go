package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return Remove(head, Length(head)-n)
}

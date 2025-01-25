package main

import "fmt"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	for list1 != nil || list2 != nil {
		if list1 == nil {
			result = Append(result, list2.Val)
			list2 = list2.Next
		} else if list2 == nil {
			result = Append(result, list1.Val)
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			result = Append(result, list1.Val)
			list1 = list1.Next
		} else {
			result = Append(result, list2.Val)
			list2 = list2.Next
		}
		fmt.Println(result)
	}

	return result
}

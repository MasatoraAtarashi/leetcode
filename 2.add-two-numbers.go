package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	// 繰り上がり
	var carry int
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		num := v1 + v2 + carry
		if num >= 10 {
			num = num - 10
			carry = 1
		} else {
			carry = 0
		}

		result = Append(result, num)
	}

	// 最後の桁が増えるケースに対応
	if carry != 0 {
		result = Append(result, carry)
	}

	return result
}

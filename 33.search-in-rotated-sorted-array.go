package main

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}

		// midの右側がソートされたセグメント
		if nums[mid] < nums[j] {
			if nums[mid] < target && target <= nums[j] {
				i = mid + 1
			} else {
				j = mid - 1
			}
			// midは左側→からの流れのところにいる
		} else {
			if nums[i] <= target && target < nums[mid] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
	}

	return -1
}

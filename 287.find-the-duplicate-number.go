
package main

func findDuplicate(nums []int) int {
	// フロイドの循環検出法を使う
	// Phase1 出会わせる
	rabbit := nums[nums[0]]
	turtle := nums[0]
	for rabbit != turtle {
		rabbit = nums[nums[rabbit]]
		turtle = nums[turtle]
	}

	// Phase2 循環開始点を見つける
	turtle = 0
	for rabbit != turtle {
		rabbit = nums[rabbit]
		turtle = nums[turtle]
	}
  
	return rabbit
}

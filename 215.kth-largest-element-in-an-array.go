package main

func findKthLargest(nums []int, k int) int {
	mh := MinHeap{}
	for i := range nums {
		mh.Push(nums[i])
		if mh.Size() > k {
			mh.Pop()
		}
	}

	return mh.data[0]
}

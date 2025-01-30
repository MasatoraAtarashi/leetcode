package main

type KthLargest struct {
	mh MinHeap
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	mh := MinHeap{}
	for _, num := range nums {
		mh.Push(num)
		if len(mh.data) > k {
			mh.Pop()
		}
	}
	return KthLargest{
		mh: mh,
		k:  k,
	}
}

func (this *KthLargest) Add(val int) int {
	this.mh.Push(val)
	if len(this.mh.data) > this.k {
		this.mh.Pop()
	}
	return this.mh.Peek()
}

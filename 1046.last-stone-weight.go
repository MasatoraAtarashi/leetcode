package main

func lastStoneWeight(stones []int) int {
	mh := MaxHeap{}
	for _, stone := range stones {
		mh.Push(stone)
	}

	for len(mh.data) > 1 {
		x := mh.Pop()
		y := mh.Pop()
		if x != y {
			mh.Push(x - y)
		}
	}

	return mh.Peek()
}

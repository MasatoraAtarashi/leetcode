package main

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	notyet := Stack[int]{}
	for idx, tmp := range temperatures {
		if notyet.Empty() {
			notyet.Push(idx)
			continue
		}
		for temperatures[notyet.Top()] < tmp && !notyet.Empty() {
			y := notyet.Pop()
			answer[y] = idx - y
		}

		notyet.Push(idx)
		// 	fmt.Println(answer, notyet)
	}

	return answer
}

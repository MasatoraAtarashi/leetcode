package main

import (
	"container/heap"
	"fmt"
)

type Task struct {
	taskType byte
	count    int
	nextTime int
}

type TaskHeap []Task

func (p TaskHeap) Len() int {
	return len(p)
}

func (p TaskHeap) Less(i, j int) bool {
	return p[i].count > p[j].count
}

func (p TaskHeap) Swap(i, j int) {
	fmt.Println("p", p)
	fmt.Println("swap", i, j)
	p[i], p[j] = p[j], p[i]
}

func (p *TaskHeap) Push(x any) {
	*p = append(*p, x.(Task))
}

func (p *TaskHeap) Pop() any {
	old := *p
	l := old[len(old)-1]
	*p = old[:len(old)-1]
	return l
}

func leastInterval(tasks []byte, n int) int {
	taskMap := make(map[byte]int, len(tasks))
	for _, t := range tasks {
		taskMap[t] += 1
	}

	wait := &TaskHeap{}
	heap.Init(wait)
	for k, v := range taskMap {
		heap.Push(wait, Task{
			taskType: k,
			count:    v,
			nextTime: 0,
		})
	}

	idleQueue := Queue[Task]{}
	var interval int
	for len(*wait) != 0 || idleQueue.Length() != 0 {
		interval++
		if idleQueue.Length() != 0 && idleQueue.linkedList.First.Val.nextTime == interval {
			t := idleQueue.Dequeue()
			heap.Push(wait, t.Val)
		}
		if len(*wait) == 0 {
			continue
		}

		t := heap.Pop(wait).(Task)

		t.count--
		t.nextTime = interval + n + 1
		if t.count > 0 {
			idleQueue.Enqueue(t)
		}
	}

	return interval
}

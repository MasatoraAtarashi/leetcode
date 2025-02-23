package main

import (
	"container/heap"
	"fmt"
)

type Point struct {
	x, y     int
	distance float64
}

type PointHeap []Point

func (p PointHeap) Len() int {
	return len(p)
}

func (p PointHeap) Less(i, j int) bool {
	return p[i].distance < p[j].distance
}

func (p PointHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PointHeap) Push(x any) {
	*p = append(*p, x.(Point))
}

func (p *PointHeap) Pop() any {
	old := *p
	l := old[len(old)-1]
	*p = old[:len(old)-1]
	return l
}

func kClosest(points [][]int, k int) [][]int {
	h := &PointHeap{}
	heap.Init(h)

	for i := range points {
		heap.Push(h, Point{
			x:        points[i][0],
			y:        points[i][1],
			distance: CalcEuclideanDistance(points[i], []int{0, 0}),
		})
	}
	fmt.Println("heap", h)

	result := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(Point)
		fmt.Println("p", p)
		result = append(result, []int{p.x, p.y})
	}

	return result
}

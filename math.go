package main

import "math"

func CalcEuclideanDistance(x, y []int) float64 {
	if len(x) != 2 || len(y) != 2 {
		return 0
	}
	xd := float64(x[0] - y[0])
	yd := float64(x[1] - y[1])
	return math.Sqrt(math.Pow(xd, 2) + math.Pow(yd, 2))
}

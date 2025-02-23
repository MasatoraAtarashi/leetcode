package main

import "testing"

func TestCalcEuclideanDistance(t *testing.T) {
	tests := []struct {
		name string
		x    []int
		y    []int
		want float64
	}{
		{
			name: "原点との距離",
			x:    []int{3, 4},
			y:    []int{0, 0},
			want: 5, // 3-4-5の直角三角形
		},
		{
			name: "2点間の距離",
			x:    []int{1, 1},
			y:    []int{4, 5},
			want: 5, // √((1-4)² + (1-5)²) = √(9 + 16) = √25 = 5
		},
		{
			name: "同じ点",
			x:    []int{2, 2},
			y:    []int{2, 2},
			want: 0,
		},
		{
			name: "不正なスライスの長さ",
			x:    []int{1},
			y:    []int{2, 2},
			want: 0,
		},
		{
			name: "空のスライス",
			x:    []int{},
			y:    []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalcEuclideanDistance(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("CalcEuclideanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_maxAreaOfIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{grid: [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{grid: [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}},
			want: 0,
		},
		{
			name: "Single cell island",
			args: args{grid: [][]int{{1}}},
			want: 1,
		},
		{
			name: "Empty grid",
			args: args{grid: [][]int{}},
			want: 0,
		},
		{
			name: "Multiple small islands",
			args: args{grid: [][]int{
				{1, 0, 1, 0},
				{0, 1, 0, 1},
				{1, 0, 1, 0},
			}},
			want: 1,
		},
		{
			name: "L shaped island",
			args: args{grid: [][]int{
				{1, 0, 0},
				{1, 0, 0},
				{1, 1, 1},
			}},
			want: 5,
		},
		{
			name: "All ones",
			args: args{grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			}},
			want: 9,
		},
		{
			name: "Two islands",
			args: args{grid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			}},
			want: 4,
		},
	}
	for _, tt := range tests {
		if got := maxAreaOfIsland(tt.args.grid); got != tt.want {
			t.Errorf("%q. maxAreaOfIsland() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

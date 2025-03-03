package main

import "testing"

func Test_orangesRotting(t *testing.T) {
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
			args: args{grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}}, // Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
			want: 4,                                                    // Output: 4
		},
		{
			name: "Example 2",
			args: args{grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}}, // Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
			want: -1,                                                   // Output: -1
		},
		{
			name: "Example 3",
			args: args{grid: [][]int{{0, 2}}}, // Input: grid = [[0,2]]
			want: 0,                           // Output: 0
		},
	}
	for _, tt := range tests {
		if got := orangesRotting(tt.args.grid); got != tt.want {
			t.Errorf("%q. orangesRotting() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

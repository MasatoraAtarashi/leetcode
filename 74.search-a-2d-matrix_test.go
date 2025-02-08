package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "Example 1",
		// 	args: args{
		// 		matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		// 		target: 3,
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "Example 2",
		// 	args: args{
		// 		matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		// 		target: 13,
		// 	},
		// 	want: false,
		// },
		{
			name: "Testcase 1",
			args: args{
				matrix: [][]int{{1}},
				target: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
			t.Errorf("%q. searchMatrix() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "Example 1",
		// 	args: args{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		// 	want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		// },
		// {
		// 	name: "Example 2",
		// 	args: args{matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}},
		// 	want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		// },
		{
			name: "Example 3",
			args: args{matrix: [][]int{{7}, {9}, {6}}},
			want: []int{7, 9, 6},
		},
	}
	for _, tt := range tests {
		if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. spiralOrder() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

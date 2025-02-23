package main

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "Example 2",
			args: args{nums: []int{0, 1}},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "Example 3",
			args: args{nums: []int{1}},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. permute() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

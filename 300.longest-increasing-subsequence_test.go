package main

import (
	"testing"
	"time"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "Example 1",
		// 	args: args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
		// 	want: 4,
		// },
		// {
		// 	name: "Example 2",
		// 	args: args{nums: []int{0, 1, 0, 3, 2, 3}},
		// 	want: 4,
		// },
		// {
		// 	name: "Example 3",
		// 	args: args{nums: []int{7, 7, 7, 7, 7, 7, 7}},
		// 	want: 1,
		// },
		// {
		// 	name: "Example 4",
		// 	args: args{nums: []int{1, 2}},
		// 	want: 2,
		// },
		{
			name: "Example 5",
			args: args{nums: []int{4, 10, 4, 3, 8, 9}},
			want: 3,
		},
	}
	for _, tt := range tests {
		start := time.Now()
		if got := lengthOfLIS(tt.args.nums); got != tt.want {
			t.Errorf("%q. lengthOfLIS() = %v, want %v", tt.name, got, tt.want)
		}
		elapsed := time.Since(start)
		t.Logf("Function took %s", elapsed)
	}
}

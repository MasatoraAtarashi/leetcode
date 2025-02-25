package main

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{[]int{2, 3, 1, 1, 4}}, true},
		{"Example 2", args{[]int{3, 2, 1, 0, 4}}, false},
		{"Example 3", args{[]int{0}}, true},
		{"Example 4", args{[]int{2, 0, 0}}, true},
		{"Example 5", args{[]int{3, 0, 8, 2, 0, 0, 1}}, true},
	}
	for _, tt := range tests {
		if got := canJump(tt.args.nums); got != tt.want {
			t.Errorf("%q. canJump() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

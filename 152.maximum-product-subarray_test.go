package main

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{2, 3, -2, 4}}, 6},
		{"Example 2", args{[]int{-2, 0, -1}}, 0},
	}
	for _, tt := range tests {
		if got := maxProduct(tt.args.nums); got != tt.want {
			t.Errorf("%q. maxProduct() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

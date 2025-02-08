package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := search(tt.args.nums, tt.args.target); got != tt.want {
			t.Errorf("%q. search() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

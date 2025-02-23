package main

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := rob(tt.args.nums); got != tt.want {
			t.Errorf("%q. rob() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

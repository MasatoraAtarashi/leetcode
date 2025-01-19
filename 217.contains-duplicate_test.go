package main

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := containsDuplicate(tt.args.nums); got != tt.want {
			t.Errorf("%q. containsDuplicate() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

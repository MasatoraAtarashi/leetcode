package main

import "testing"

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"abc"}, 3},
		{"Example 2", args{"aaa"}, 6},
	}
	for _, tt := range tests {
		if got := countSubstrings(tt.args.s); got != tt.want {
			t.Errorf("%q. countSubstrings() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

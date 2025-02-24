package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {"Example 1", args{"babad"}, "aba"},
		{"Example 2", args{"cbbd"}, "bb"},
		{"Example 3", args{"ab"}, "a"},
		{"Example 4", args{"a"}, "a"},
		{"Example 5", args{"aaaa"}, "aaaa"},
		{"Example 19", args{"aaa"}, "aaa"},
		{"Example 6", args{"racecarxyz"}, "racecar"},
		{"Example 7", args{"xyzabcdefedcba123"}, "abcdefedcba"},
		{"Example 8", args{"abcdef"}, "a"},
		{"Example 9", args{""}, ""},
	}
	for _, tt := range tests {
		if got := longestPalindrome(tt.args.s); got != tt.want {
			t.Errorf("%q. longestPalindrome() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

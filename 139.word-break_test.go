package main

import "testing"

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{"leetcode", []string{"leet", "code"}},
			want: true,
		},
		{
			name: "Example 2",
			args: args{"applepenapple", []string{"apple", "pen"}},
			want: true,
		},
		{
			name: "Example 3",
			args: args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
			want: false,
		},
		{
			name: "Single character word",
			args: args{"a", []string{"a"}},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
			t.Errorf("%q. wordBreak() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "race a car",
			args: args{
				s: "race a car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := isPalindrome(tt.args.s); got != tt.want {
			t.Errorf("%q. isPalindrome() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

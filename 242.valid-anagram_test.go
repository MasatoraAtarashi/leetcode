package main

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "a",
				t: "ab",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "ab",
				t: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
			t.Errorf("%q. isAnagram() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

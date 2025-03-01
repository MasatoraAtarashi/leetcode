package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{123},
			want: 321,
		},
		{
			name: "Example 2",
			args: args{-123},
			want: -321,
		},
		{
			name: "Example 3",
			args: args{120},
			want: 21,
		},
		{
			name: "Example 4",
			args: args{1534236469},
			want: 0,
		},
	}
	for _, tt := range tests {
		if got := reverse(tt.args.x); got != tt.want {
			t.Errorf("%q. reverse() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

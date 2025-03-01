package main

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{"12"},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{"226"},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{"06"},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{"10"},
			want: 1,
		},
		{
			name: "Example 5",
			args: args{"2101"},
			want: 1,
		},
		{
			name: "Example 6",
			args: args{"1123"},
			want: 5,
		},
	}
	for _, tt := range tests {
		if got := numDecodings(tt.args.s); got != tt.want {
			t.Errorf("%q. numDecodings() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

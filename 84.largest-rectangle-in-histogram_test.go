package main

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{heights: []int{2, 1, 5, 6, 2, 3}},
			want: 10,
		},
		{
			name: "Test Case 2",
			args: args{heights: []int{2, 4}},
			want: 4,
		},
		{
			name: "Test Case 3",
			args: args{heights: []int{1, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		if got := largestRectangleArea(tt.args.heights); got != tt.want {
			t.Errorf("%q. largestRectangleArea() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

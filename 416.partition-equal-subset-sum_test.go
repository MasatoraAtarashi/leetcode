package main

import "testing"

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 5, 11, 5}},
			want: true,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 2, 3, 5}},
			want: false,
		},
		{
			name: "Example 3",
			args: args{nums: []int{1, 2, 5}},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := canPartition(tt.args.nums); got != tt.want {
			t.Errorf("%q. canPartition() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

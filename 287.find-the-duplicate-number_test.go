package main

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1,3,4,2,2}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{nums: []int{3,1,3,4,2}},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{nums: []int{3,3,3,3,3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		if got := findDuplicate(tt.args.nums); got != tt.want {
			t.Errorf("%q. findDuplicate() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

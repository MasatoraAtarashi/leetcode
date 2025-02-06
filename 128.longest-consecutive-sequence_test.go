package main

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1: sequence [1,2,3,4]",
			args: args{
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			want: 4,
		},
		{
			name: "Example 2: sequence [0,1,2,3,4,5,6,7,8]",
			args: args{
				nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			},
			want: 9,
		},
		{
			name: "Empty array",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "Single element",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "No consecutive sequence",
			args: args{
				nums: []int{2, 4, 6, 8},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		if got := longestConsecutive(tt.args.nums); got != tt.want {
			t.Errorf("%q. longestConsecutive() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

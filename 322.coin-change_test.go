package main

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{coins: []int{1, 2, 5}, amount: 11},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{coins: []int{2}, amount: 3},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{coins: []int{1}, amount: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
			t.Errorf("%q. coinChange() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

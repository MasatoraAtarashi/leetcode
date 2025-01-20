package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. twoSum() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

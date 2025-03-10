package main

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "Example 2",
			args: args{root: &TreeNode{Val: 1}},
			want: [][]int{{1}},
		},
		{
			name: "Example 3",
			args: args{root: nil},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. levelOrder() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

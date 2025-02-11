package main

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{n: 3},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "Example 2",
			args: args{n: 1},
			want: []string{"()"},
		},
	}
	for _, tt := range tests {
		if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. generateParenthesis() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

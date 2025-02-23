package main

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{digits: "23"},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "Example 2",
			args: args{digits: ""},
			want: nil,
		},
		{
			name: "Example 3",
			args: args{digits: "2"},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. letterCombinations() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

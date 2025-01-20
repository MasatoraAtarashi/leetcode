package main

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
	}
	for _, tt := range tests {
		if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. groupAnagrams() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

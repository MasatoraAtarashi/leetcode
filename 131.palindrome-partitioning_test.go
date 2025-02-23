package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// {
		// 	name: "Example 1",
		// 	args: args{s: "aab"},
		// 	want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		// },
		// {
		// 	name: "Example 2",
		// 	args: args{s: "a"},
		// 	want: [][]string{{"a"}},
		// },
		{
			name: "Example 3",
			args: args{s: "efe"},
			want: [][]string{{"e", "f", "e"}, {"efe"}},
		},
	}
	for _, tt := range tests {
		if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. partition() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

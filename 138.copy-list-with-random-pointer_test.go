package main

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *LinkedListNode
	}
	tests := []struct {
		name string
		args args
		want *LinkedListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. copyRandomList() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

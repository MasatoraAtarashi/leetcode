package main

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := getSum(tt.args.a, tt.args.b); got != tt.want {
			t.Errorf("%q. getSum() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

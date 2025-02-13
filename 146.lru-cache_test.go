package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LRUCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := ___Constructor(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Constructor() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestLRUCache_Get(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		this *LRUCache
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		this := &LRUCache{}
		if got := this.Get(tt.args.key); got != tt.want {
			t.Errorf("%q. LRUCache.Get() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestLRUCache_Put(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		this *LRUCache
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		this := &LRUCache{}
		this.Put(tt.args.key, tt.args.value)
	}
}

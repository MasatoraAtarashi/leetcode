package main

import (
	"reflect"
	"testing"
)

func TestRemoveIndex(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected []int
	}{
		{
			name:     "基本的な要素の削除",
			slice:    []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "最初の要素の削除",
			slice:    []int{1, 2, 3},
			index:    0,
			expected: []int{2, 3},
		},
		{
			name:     "最後の要素の削除",
			slice:    []int{1, 2, 3},
			index:    2,
			expected: []int{1, 2},
		},
		{
			name:     "test",
			slice:    []int{2, 3},
			index:    1,
			expected: []int{2},
		},
		{
			name:     "1要素のスライスから要素を削除",
			slice:    []int{1},
			index:    0,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveIndex(tt.slice, tt.index)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RemoveIndex() = %v, want %v", result, tt.expected)
			}
		})
	}

	// 文字列スライスでのテスト（ジェネリクスの動作確認）
	strSlice := []string{"a", "b", "c"}
	expectedStrSlice := []string{"a", "c"}
	if result := RemoveIndex(strSlice, 1); !reflect.DeepEqual(result, expectedStrSlice) {
		t.Errorf("RemoveIndex() with strings = %v, want %v", result, expectedStrSlice)
	}
}

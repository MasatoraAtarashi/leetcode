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

func TestSliceContains(t *testing.T) {
	// int型のテスト
	t.Run("int type", func(t *testing.T) {
		tests := []struct {
			name     string
			slice    []int
			value    int
			expected bool
		}{
			{"value exists", []int{1, 2, 3, 4, 5}, 3, true},
			{"value does not exist", []int{1, 2, 3, 4, 5}, 6, false},
			{"empty slice", []int{}, 1, false},
			{"single element slice", []int{42}, 42, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := SliceContains(tt.slice, tt.value)
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			})
		}
	})

	// string型のテスト
	t.Run("string type", func(t *testing.T) {
		tests := []struct {
			name     string
			slice    []string
			value    string
			expected bool
		}{
			{"value exists", []string{"apple", "banana", "cherry"}, "banana", true},
			{"value does not exist", []string{"apple", "banana", "cherry"}, "grape", false},
			{"empty slice", []string{}, "apple", false},
			{"single element slice", []string{"hello"}, "hello", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := SliceContains(tt.slice, tt.value)
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			})
		}
	})

	// float64型のテスト
	t.Run("float64 type", func(t *testing.T) {
		tests := []struct {
			name     string
			slice    []float64
			value    float64
			expected bool
		}{
			{"value exists", []float64{1.1, 2.2, 3.3}, 2.2, true},
			{"value does not exist", []float64{1.1, 2.2, 3.3}, 4.4, false},
			{"empty slice", []float64{}, 1.1, false},
			{"single element slice", []float64{3.14}, 3.14, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := SliceContains(tt.slice, tt.value)
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			})
		}
	})

	// bool型のテスト
	t.Run("bool type", func(t *testing.T) {
		tests := []struct {
			name     string
			slice    []bool
			value    bool
			expected bool
		}{
			{"value exists", []bool{true, false}, true, true},
			{"value does not exist", []bool{true, true}, false, false},
			{"empty slice", []bool{}, true, false},
			{"single element slice", []bool{false}, false, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := SliceContains(tt.slice, tt.value)
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			})
		}
	})
}

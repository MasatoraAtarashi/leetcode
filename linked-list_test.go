package main

import (
	"testing"
)

// リストの値を配列として取得するヘルパー関数
func getValues(l *ListNode) []int {
	var result []int
	current := l
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

// 配列同士を比較するヘルパー関数
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Appendのテスト
func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		val      int
		expected []int
	}{
		{
			name:     "空のリストに追加",
			initial:  []int{},
			val:      1,
			expected: []int{1},
		},
		{
			name:     "1要素のリストに追加",
			initial:  []int{1},
			val:      2,
			expected: []int{1, 2},
		},
		{
			name:     "複数要素のリストに追加",
			initial:  []int{1, 2, 3},
			val:      4,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初期リストの作成
			var list *ListNode
			for _, v := range tt.initial {
				list = Append(list, v)
			}

			// テスト対象の実行
			list = Append(list, tt.val)

			// 結果の検証
			result := getValues(list)
			if !compareSlices(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// Prependのテスト
func TestPrepend(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		val      int
		expected []int
	}{
		{
			name:     "空のリストに追加",
			initial:  []int{},
			val:      1,
			expected: []int{1},
		},
		{
			name:     "1要素のリストに追加",
			initial:  []int{1},
			val:      2,
			expected: []int{2, 1},
		},
		{
			name:     "複数要素のリストに追加",
			initial:  []int{1, 2, 3},
			val:      4,
			expected: []int{4, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初期リストの作成
			var list *ListNode
			for _, v := range tt.initial {
				list = Append(list, v)
			}

			// テスト対象の実行
			list = Prepend(list, tt.val)

			// 結果の検証
			result := getValues(list)
			if !compareSlices(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// Popのテスト
func TestPop(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected []int
	}{
		{
			name:     "空のリストからのPop",
			initial:  []int{},
			expected: []int{}, // 何も変化しない
		},
		{
			name:     "1要素のリストからのPop",
			initial:  []int{1},
			expected: []int{}, // 空になる
		},
		{
			name:     "2要素のリストから削除",
			initial:  []int{1, 2},
			expected: []int{1},
		},
		{
			name:     "複数要素のリストから削除",
			initial:  []int{1, 2, 3, 4},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初期リストの作成
			var list *ListNode
			for _, v := range tt.initial {
				list = Append(list, v)
			}

			// 結果の検証
			result := getValues(Pop(list))
			if !compareSlices(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// Reverseのテスト
func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected []int
	}{
		{
			name:     "空のリスト",
			initial:  []int{},
			expected: []int{},
		},
		{
			name:     "1要素のリスト",
			initial:  []int{1},
			expected: []int{1},
		},
		{
			name:     "2要素のリスト",
			initial:  []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "複数要素のリスト",
			initial:  []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初期リストの作成
			var list *ListNode
			for _, v := range tt.initial {
				list = Append(list, v)
			}

			// 結果の検証
			result := getValues(Reverse(list))
			if !compareSlices(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// Lengthのテスト
func TestLength(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected int
	}{
		{
			name:     "空のリスト",
			initial:  []int{},
			expected: 0,
		},
		{
			name:     "1要素のリスト",
			initial:  []int{1},
			expected: 1,
		},
		{
			name:     "2要素のリスト",
			initial:  []int{1, 2},
			expected: 2,
		},
		{
			name:     "複数要素のリスト",
			initial:  []int{1, 2, 3, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初期リストの作成
			var list *ListNode
			for _, v := range tt.initial {
				list = Append(list, v)
			}

			// 結果の検証
			result := Length(list)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// Removeのテスト
func TestRemove(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		index    int
		expected []int
	}{
		{
			name:     "空のリストから削除",
			initial:  []int{},
			index:    0,
			expected: []int{},
		},
		{
			name:     "先頭の要素を削除",
			initial:  []int{1, 2, 3},
			index:    0,
			expected: []int{2, 3},
		},
		{
			name:     "中間の要素を削除",
			initial:  []int{1, 2, 3},
			index:    1,
			expected: []int{1, 3},
		},
		{
			name:     "最後の要素を削除",
			initial:  []int{1, 2, 3},
			index:    2,
			expected: []int{1, 2},
		},
		{
			name:     "1要素のリストから削除",
			initial:  []int{1},
			index:    0,
			expected: []int{},
		},
		{
			name:     "存在しないインデックスの削除（範囲外）",
			initial:  []int{1, 2, 3},
			index:    5,
			expected: []int{1, 2, 3}, // 変更なし
		},
		{
			name:     "負のインデックスの削除",
			initial:  []int{1, 2, 3},
			index:    -1,
			expected: []int{1, 2, 3}, // 変更なし
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初期リストの作成
			var list *ListNode
			for _, v := range tt.initial {
				list = Append(list, v)
			}

			// テスト対象の実行
			list = Remove(list, tt.index)

			// 結果の検証
			result := getValues(list)
			if !compareSlices(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

package main

import (
	"testing"
)

// Enqueueのテスト
func TestEnqueue(t *testing.T) {
	// 基本的な追加操作
	t.Run("Basic enqueue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)

		if q.linkedList.length != 1 {
			t.Errorf("Expected length 1, got %d", q.linkedList.length)
		}
		if q.linkedList.First.Val != 1 {
			t.Errorf("Expected first value 1, got %d", q.linkedList.First.Val)
		}
		if q.linkedList.Last.Val != 1 {
			t.Errorf("Expected last value 1, got %d", q.linkedList.Last.Val)
		}
	})

	// 複数要素の追加
	t.Run("Multiple enqueues", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		if q.linkedList.length != 3 {
			t.Errorf("Expected length 3, got %d", q.linkedList.length)
		}
		if q.linkedList.First.Val != 1 {
			t.Errorf("Expected first value 1, got %d", q.linkedList.First.Val)
		}
		if q.linkedList.Last.Val != 3 {
			t.Errorf("Expected last value 3, got %d", q.linkedList.Last.Val)
		}
	})

	// 大量のデータ追加
	t.Run("Large number of enqueues", func(t *testing.T) {
		q := NewQueue[int]()
		for i := 0; i < 1000; i++ {
			q.Enqueue(i)
		}
		if q.linkedList.length != 1000 {
			t.Errorf("Expected length 1000, got %d", q.linkedList.length)
		}
		if q.linkedList.First.Val != 0 {
			t.Errorf("Expected first value 0, got %d", q.linkedList.First.Val)
		}
		if q.linkedList.Last.Val != 999 {
			t.Errorf("Expected last value 999, got %d", q.linkedList.Last.Val)
		}
	})

	// 異なる型のテスト
	t.Run("Different types", func(t *testing.T) {
		qString := NewQueue[string]()
		qString.Enqueue("hello")
		qString.Enqueue("world")

		if qString.linkedList.length != 2 {
			t.Errorf("Expected length 2, got %d", qString.linkedList.length)
		}
		if qString.linkedList.First.Val != "hello" {
			t.Errorf("Expected first value 'hello', got %s", qString.linkedList.First.Val)
		}
	})
}

// Dequeueのテスト// Dequeueのテスト
func TestDequeue(t *testing.T) {
	// 基本的な削除操作
	t.Run("Basic dequeue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)
		removed := q.Dequeue()

		if removed == nil {
			t.Error("Expected removed node, got nil")
		}
		if removed.Val != 1 {
			t.Errorf("Expected removed value 1, got %d", removed.Val)
		}
		if q.linkedList.length != 0 {
			t.Errorf("Expected length 0, got %d", q.linkedList.length)
		}
	})

	// 空のキューからの削除
	t.Run("Dequeue from empty queue", func(t *testing.T) {
		q := NewQueue[int]()
		removed := q.Dequeue()

		if removed != nil {
			t.Error("Expected nil, got a node")
		}
	})

	// 複数要素がある場合の削除
	t.Run("Multiple dequeues", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		// 最初の要素を削除
		removed := q.Dequeue()
		if removed == nil || removed.Val != 1 {
			t.Errorf("Expected removed value 1, got %v", removed)
		}
		if q.linkedList.length != 2 {
			t.Errorf("Expected length 2, got %d", q.linkedList.length)
		}
		if q.linkedList.First.Val != 2 {
			t.Errorf("Expected first value 2, got %d", q.linkedList.First.Val)
		}

		// 2番目の要素を削除
		removed = q.Dequeue()
		if removed == nil || removed.Val != 2 {
			t.Errorf("Expected removed value 2, got %v", removed)
		}
		if q.linkedList.length != 1 {
			t.Errorf("Expected length 1, got %d", q.linkedList.length)
		}
		if q.linkedList.First.Val != 3 {
			t.Errorf("Expected first value 3, got %d", q.linkedList.First.Val)
		}
	})

	// 最後の要素の削除
	t.Run("Dequeue last element", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)
		removed := q.Dequeue()

		if removed == nil || removed.Val != 1 {
			t.Errorf("Expected removed value 1, got %v", removed)
		}
		if q.linkedList.length != 0 {
			t.Errorf("Expected length 0, got %d", q.linkedList.length)
		}
		if q.linkedList.First != nil || q.linkedList.Last != nil {
			t.Error("Expected nil First and Last after removing last element")
		}
	})

	// 順序のテスト
	t.Run("Dequeue order test", func(t *testing.T) {
		q := NewQueue[int]()
		expected := []int{1, 2, 3, 4, 5}

		// 要素を追加
		for _, v := range expected {
			q.Enqueue(v)
		}

		// 順番に取り出して確認
		for _, want := range expected {
			removed := q.Dequeue()
			if removed == nil {
				t.Fatal("Expected node, got nil")
			}
			if removed.Val != want {
				t.Errorf("Expected value %d, got %d", want, removed.Val)
			}
		}

		// キューが空になっていることを確認
		if q.linkedList.length != 0 {
			t.Errorf("Expected empty queue, got length %d", q.linkedList.length)
		}
	})

	// 連続したDequeue操作のテスト
	t.Run("Consecutive dequeues", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)

		// 最初のDequeue
		first := q.Dequeue()
		if first == nil || first.Val != 1 {
			t.Errorf("Expected value 1, got %v", first)
		}

		// 空のキューに対する連続したDequeue
		second := q.Dequeue()
		if second != nil {
			t.Errorf("Expected nil, got %v", second)
		}

		third := q.Dequeue()
		if third != nil {
			t.Errorf("Expected nil, got %v", third)
		}
	})
}

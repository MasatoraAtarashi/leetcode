package main

import (
	"fmt"
	"testing"
)

// MinHeap のテスト
func TestMinHeap(t *testing.T) {
	h := &MinHeap{}

	t.Run("Push and Peek", func(t *testing.T) {
		fmt.Println("push 5")
		h.Push(5)
		fmt.Println(h.data)
		fmt.Println("push 3")
		h.Push(3)
		fmt.Println(h.data)
		fmt.Println("push 8")
		h.Push(8)
		fmt.Println(h.data)
		fmt.Println("push 1")
		h.Push(1)
		fmt.Println(h.data)
		fmt.Println("push 6")
		h.Push(6)
		fmt.Println(h.data)

		expectedMin := 1 // 最小値は 1 のはず
		if h.Peek() != expectedMin {
			t.Errorf("Expected min element to be %d, got %d", expectedMin, h.Peek())
		}
	})

	t.Run("Pop", func(t *testing.T) {
		// 1 を削除（最小値のはず）
		fmt.Println("pop 1")
		fmt.Println(h.data)
		popped := h.Pop()
		fmt.Println(h.data)
		if popped != 1 {
			t.Errorf("Expected popped element to be 1, got %d", popped)
		}

		// 次の最小値は 3 のはず
		expectedMin := 3
		if h.Peek() != expectedMin {
			t.Errorf("Expected new min element to be %d, got %d", expectedMin, h.Peek())
		}
		fmt.Println(h.data)
	})

	t.Run("Size", func(t *testing.T) {
		expectedSize := 4 // 5つ入れて1つ削除したので、残りは4
		if h.Size() != expectedSize {
			t.Errorf("Expected heap size to be %d, got %d", expectedSize, h.Size())
		}
	})

	t.Run("Empty", func(t *testing.T) {
		// 残りの要素を全て削除
		for !h.IsEmpty() {
			h.Pop()
			fmt.Println("pop")
		}

		// 空のはず
		if !h.IsEmpty() {
			t.Errorf("Heap should be empty after popping all elements")
		}
	})

	t.Run("Simulate KthLargest operations", func(t *testing.T) {
		h := &MinHeap{}

		fmt.Println("push 4")
		h.Push(4)
		fmt.Println("Heap:", h.data)

		fmt.Println("push 5")
		h.Push(5)
		fmt.Println("Heap:", h.data)

		fmt.Println("push 8")
		h.Push(8)
		fmt.Println("Heap:", h.data)

		fmt.Println("push 2")
		h.Push(2)
		fmt.Println("Heap:", h.data)

		fmt.Println("pop")
		h.Pop()
		fmt.Println("Heap:", h.data)

		fmt.Println("push 3")
		h.Push(3)
		fmt.Println("Heap:", h.data)

		fmt.Println("pop")
		h.Pop()
		fmt.Println("Heap:", h.data)

		fmt.Println("push 5")
		h.Push(5)
		fmt.Println("Heap:", h.data)

		fmt.Println("pop")
		h.Pop()
		fmt.Println("Heap:", h.data)

		fmt.Println("push 10")
		h.Push(10)
		fmt.Println("Heap:", h.data)

		fmt.Println("pop")
		h.Pop()
		fmt.Println("Heap:", h.data)

		fmt.Println("push 9")
		h.Push(9)
		fmt.Println("Heap:", h.data)

		fmt.Println("pop")
		h.Pop()
		fmt.Println("Heap:", h.data)

		fmt.Println("push 4")
		h.Push(4)
		fmt.Println("Heap:", h.data)

		fmt.Println("pop")
		h.Pop()
		fmt.Println("Heap:", h.data)
	})
}

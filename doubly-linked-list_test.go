package main

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		dll := &DoublyLinkedList[int]{}
		if dll.Length() != 0 {
			t.Errorf("expected length 0, got %d", dll.Length())
		}
		if dll.Head != nil || dll.Tail != nil {
			t.Error("head and Tail should be nil in empty list")
		}
	})

	t.Run("append", func(t *testing.T) {
		dll := &DoublyLinkedList[int]{}

		// 1つ目の要素追加
		dll.Append(1)
		if dll.Length() != 1 || dll.Head.Val != 1 || dll.Tail.Val != 1 {
			t.Errorf("after first append: length=%d, head=%d, tail=%d", dll.Length(), dll.Head.Val, dll.Tail.Val)
		}

		// 2つ目の要素追加
		dll.Append(2)
		if dll.Length() != 2 || dll.Head.Val != 1 || dll.Tail.Val != 2 {
			t.Errorf("after second append: length=%d, head=%d, tail=%d", dll.Length(), dll.Head.Val, dll.Tail.Val)
		}

		// 前後の接続確認
		if dll.Head.Next.Val != 2 || dll.Tail.Prev.Val != 1 {
			t.Error("links between nodes are incorrect")
		}
	})

	t.Run("prepend", func(t *testing.T) {
		dll := &DoublyLinkedList[int]{}

		// 1つ目の要素追加
		dll.Prepend(1)
		fmt.Printf("After first prepend: length=%d, head=%v, tail=%v\n", dll.Length(), dll.Head.Val, dll.Tail.Val)
		if dll.Length() != 1 || dll.Head.Val != 1 || dll.Tail.Val != 1 {
			t.Errorf("after first prepend: length=%d, head=%d, tail=%d", dll.Length(), dll.Head.Val, dll.Tail.Val)
		}

		// 2つ目の要素追加
		dll.Prepend(2)
		fmt.Printf("After second prepend: length=%d, head=%v, tail=%v\n", dll.Length(), dll.Head.Val, dll.Tail.Val)
		fmt.Printf("Links: Head.Next=%v, Tail.Prev=%v\n", dll.Head.Next.Val, dll.Tail.Prev.Val)
		if dll.Length() != 2 || dll.Head.Val != 2 || dll.Tail.Val != 1 {
			t.Errorf("after second prepend: length=%d, head=%d, tail=%d", dll.Length(), dll.Head.Val, dll.Tail.Val)
		}

		// 前後の接続確認
		if dll.Head.Next.Val != 1 || dll.Tail.Prev.Val != 2 {
			t.Error("links between nodes are incorrect")
		}
	})

	t.Run("pop", func(t *testing.T) {
		tests := []struct {
			name       string
			setup      []int
			want       int
			wantLength int
			wantHead   *int
			wantTail   *int
		}{
			{
				name:       "empty list",
				wantLength: 0,
				wantHead:   nil,
				wantTail:   nil,
			},
			{
				name:       "single element",
				setup:      []int{1},
				want:       1,
				wantLength: 0,
				wantHead:   nil,
				wantTail:   nil,
			},
			{
				name:       "multiple elements",
				setup:      []int{1, 2},
				want:       2,
				wantLength: 1,
				wantHead:   intPtr(1),
				wantTail:   intPtr(1),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				dll := &DoublyLinkedList[int]{}
				for _, v := range tt.setup {
					dll.Append(v)
				}
				fmt.Printf("Before pop: %v\n", dll.String())

				val := dll.Pop()
				fmt.Printf("After pop: %v, popped: %v\n", dll.String(), val)

				if val != tt.want {
					t.Errorf("Pop() = %v, want %v", val, tt.want)
				}
				if dll.Length() != tt.wantLength {
					t.Errorf("Length = %v, want %v", dll.Length(), tt.wantLength)
				}
				if tt.wantHead == nil {
					if dll.Head != nil {
						t.Error("Head should be nil")
					}
				} else if dll.Head.Val != *tt.wantHead {
					t.Errorf("Head = %v, want %v", dll.Head.Val, *tt.wantHead)
				}
				if tt.wantTail == nil {
					if dll.Tail != nil {
						t.Error("Tail should be nil")
					}
				} else if dll.Tail.Val != *tt.wantTail {
					t.Errorf("Tail = %v, want %v", dll.Tail.Val, *tt.wantTail)
				}
			})
		}
	})

	t.Run("remove val", func(t *testing.T) {
		tests := []struct {
			name     string
			setup    []int
			remove   int
			want     []int
			wantBool bool
		}{
			{
				name:     "remove from middle",
				setup:    []int{1, 2, 3},
				remove:   2,
				want:     []int{1, 3},
				wantBool: true,
			},
			{
				name:     "remove head",
				setup:    []int{1, 2},
				remove:   1,
				want:     []int{2},
				wantBool: true,
			},
			{
				name:     "remove tail",
				setup:    []int{1, 2},
				remove:   2,
				want:     []int{1},
				wantBool: true,
			},
			{
				name:     "remove from empty list",
				setup:    []int{},
				remove:   1,
				want:     []int{},
				wantBool: false,
			},
			{
				name:     "remove non-existent",
				setup:    []int{1, 2},
				remove:   3,
				want:     []int{1, 2},
				wantBool: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				dll := &DoublyLinkedList[int]{}
				for _, v := range tt.setup {
					dll.Append(v)
				}
				fmt.Println("default: ", dll.String())

				got := dll.RemoveVal(tt.remove)
				if got != tt.wantBool {
					t.Errorf("Remove() = %v, want %v", got, tt.wantBool)
				}
				fmt.Println("result: ", dll.String())

				// Check length
				if dll.Length() != len(tt.want) {
					t.Errorf("Length = %v, want %v", dll.Length(), len(tt.want))
				}

				// Check values
				current := dll.Head
				for i, want := range tt.want {
					if current == nil {
						t.Errorf("Node %d: got nil, want %v", i, want)
						break
					}
					if current.Val != want {
						t.Errorf("Node %d = %v, want %v", i, current.Val, want)
					}
					current = current.Next
				}
			})
		}
	})
}

func intPtr(i int) *int {
	return &i
}

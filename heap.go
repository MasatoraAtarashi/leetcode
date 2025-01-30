package main

type MinHeap struct {
	data []int
}

// Push	要素を追加し、ヒープ条件を維持
func (m *MinHeap) Push(val int) {
	m.data = append(m.data, val)
	m.heapifyUp(len(m.data) - 1)
}

// Pop	最小値（ルート）を削除し、ヒープを再構成
func (m *MinHeap) Pop() int {
	if m.IsEmpty() {
		return 0
	}

	min := m.data[0]
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	m.heapifyDown(0)
	return min
}

// Peek	最小値を取得（削除しない）
func (m *MinHeap) Peek() int {
	if m.IsEmpty() {
		return 0
	}

	return m.data[0]
}

// Size	ヒープのサイズを取得
func (m *MinHeap) Size() int {
	return len(m.data)
}

// IsEmpty ヒープが空かを確認
func (m *MinHeap) IsEmpty() bool {
	return len(m.data) == 0
}

func (m *MinHeap) heapifyUp(idx int) {
	parent := (idx - 1) / 2
	for idx > 0 && m.data[parent] > m.data[idx] {
		m.swap(parent, idx)
		idx = parent
		m.heapifyUp(idx)
	}
}

func (m *MinHeap) heapifyDown(idx int) {
	leftChild := 2*idx + 1
	for leftChild < len(m.data) {
		smallest := leftChild
		rightChild := leftChild + 1
		if rightChild < len(m.data) && m.data[rightChild] < m.data[smallest] {
			smallest = rightChild
		}
		if m.data[idx] <= m.data[smallest] {
			break
		}
		m.swap(idx, smallest)
		idx = smallest
	}
}

func (m *MinHeap) swap(idx1, idx2 int) {
	tmp := m.data[idx1]
	m.data[idx1] = m.data[idx2]
	m.data[idx2] = tmp
}

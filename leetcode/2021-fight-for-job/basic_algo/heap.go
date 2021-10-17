package basic_algo

import "sort"

// 最大堆
type MaxHeap []int

func (m *MaxHeap) Len() int {
	return len(*m)
}

func (m *MaxHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *MaxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

// 最大堆基于sort.IntSlice
type MaxHeapByIntSlice struct {
	sort.IntSlice
}

func (m *MaxHeapByIntSlice) Less(i, j int) bool {
	return m.IntSlice[i] > m.IntSlice[j]
}

func (m *MaxHeapByIntSlice) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}

func (m *MaxHeapByIntSlice) Pop() (v interface{}) {
	m.IntSlice, v = (m.IntSlice)[:m.Len()-1], (m.IntSlice)[m.Len()-1]
	return
}

// 最小堆
type MinHeap []int

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	var v int
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

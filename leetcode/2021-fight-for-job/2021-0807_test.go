package _021_fight_for_job

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type minHeap1 []int

func (m *minHeap1) Len() int {
	return len(*m)
}

func (m *minHeap1) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap1) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap1) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap1) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

func (m *minHeap1) Top() (v interface{}) {
	return (*m)[0]
}

// 第一题
// 数据流中第k大的数据
type KthLargest struct {
	k int
	sort.IntSlice
}

func (k1 *KthLargest) Push(x interface{}) {
	k1.IntSlice = append(k1.IntSlice, x.(int))
}

func (k1 *KthLargest) Pop() (v interface{}) {
	k1.IntSlice, v = (k1.IntSlice)[:k1.Len()-1], (k1.IntSlice)[k1.Len()-1]
	return
}

func Constructor3(k int, nums []int) KthLargest {
	k1 := KthLargest{k: k}
	for _, v := range nums {
		k1.Add(v)
	}
	return k1
}

func (k1 *KthLargest) Add(val int) int {
	heap.Push(k1, val)
	if k1.Len() > k1.k {
		heap.Pop(k1)
	}
	return k1.IntSlice[0]
}

// 第三题
// 最后一块石头的重量
func lastStoneWeight(stones []int) int {
	m := &maxHeap1{}
	for _, v := range stones {
		heap.Push(m, v)
	}
	for m.Len() > 1 {
		x := heap.Pop(m).(int)
		y := heap.Pop(m).(int)
		if x != y {
			heap.Push(m, x-y)
		}
	}
	if m.Len() > 0 {
		return heap.Pop(m).(int)
	}
	return 0
}

// 第四题
// 最大值和次大值
func maxProduct(nums []int) int {
	max1, max2 := 0, 0
	for i := range nums {
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}
	return (max1 - 1) * (max2 - 1)
}

func maxProduct1(nums []int) int {
	m := &maxHeap1{nums}
	heap.Init(m)
	max1 := heap.Pop(m).(int)
	max2 := heap.Pop(m).(int)
	return (max1 - 1) * (max2 - 1)
}

type maxHeap1 struct {
	sort.IntSlice
}

func (m *maxHeap1) Less(i, j int) bool {
	return m.IntSlice[i] > m.IntSlice[j]
}

func (m *maxHeap1) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}

func (m *maxHeap1) Pop() (v interface{}) {
	m.IntSlice, v = (m.IntSlice)[:m.Len()-1], (m.IntSlice)[m.Len()-1]
	return
}

// 第五题
func findKthLargest(nums []int, k int) int {
	m := &minHeap{}
	for i := range nums {
		heap.Push(m, nums[i])
		if m.Len() > k {
			heap.Pop(m)
		}
	}
	return heap.Pop(m).(int)
}

// 第六题
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	h := &minHeapPair{}
	for kk, v := range m {
		heap.Push(h, pair{kk, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, 0, k)
	for h.Len() != 0 {
		ans = append(ans, heap.Pop(h).(pair).v)
	}
	return ans
}

func Test_topKFrequent(t *testing.T) {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

type pair struct {
	v     int
	times int
}

type minHeapPair []pair

func (m *minHeapPair) Len() int {
	return len(*m)
}

func (m *minHeapPair) Less(i, j int) bool {
	return (*m)[i].times < (*m)[j].times
}

func (m *minHeapPair) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapPair) Push(x interface{}) {
	*m = append(*m, x.(pair))
}

func (m *minHeapPair) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

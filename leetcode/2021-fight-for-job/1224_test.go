package ojeveryday

import (
	"container/heap"
	"sort"
)

// tag-[贪心]
// leetcode870： 优势洗牌
func advantageCount(nums1 []int, nums2 []int) []int {
	m := &maxheap{}
	for i := range nums2 {
		heap.Push(m, kv{i, nums2[i]})
	}
	sort.Ints(nums1)
	res := make([]int, len(nums1))
	l, r := 0, len(nums1)-1
	for m.Len() != 0 {
		value := heap.Pop(m).(kv)
		if value.v < nums1[r] {
			res[value.i] = nums1[r]
			r--
		} else {
			res[value.i] = nums1[l]
			l++
		}
	}
	return res
}

type kv struct {
	i int
	v int
}

type maxheap []kv

func (m *maxheap) Less(i, j int) bool {
	return (*m)[i].v > (*m)[j].v
}

func (m *maxheap) Len() int {
	return len(*m)
}

func (m *maxheap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxheap) Push(x interface{}) {
	(*m) = append(*m, x.(kv))
}

func (m *maxheap) Pop() interface{} {
	var v interface{}
	v, (*m) = (*m)[m.Len()-1], (*m)[:m.Len()-1]
	return v
}

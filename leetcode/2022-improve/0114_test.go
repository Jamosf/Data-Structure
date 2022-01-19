package _022_improve

import (
	"container/heap"
)

// tag-[堆]
// leetcode373: 查找和最小的K对数字
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	ans := make([][]int, 0, k)
	h := &hp{}
	for i := 0; i < n1; i++{
		heap.Push(h, []int{nums1[i]+nums2[0], i, 0})
	}
	for h.Len() > 0 && len(ans) < k{
		v := heap.Pop(h).([]int)
		v1, v2 := v[1], v[2]
		ans = append(ans, []int{nums1[v1], nums2[v2]})
		if v2+1 < n2{
			heap.Push(h, []int{nums1[v1]+nums2[v2+1], v1, v2+1})
		}
	}
	return ans
}

type hp [][]int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *hp) Pop() interface{} {
	var v interface{}
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return v
}

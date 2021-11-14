package ojeveryday

import (
	"container/heap"
	"fmt"
	"leetcode/leetcode/2021-fight-for-job/basic_algo"
	"sort"
	"testing"
)

// 第一题
// leetcode703: 数据流中的第K大元素
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

func ConstructorKthLargest(k int, nums []int) KthLargest {
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
// leetcode1046: 最后一块石头的重量
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
// leetcode1464: 数组中两元素的最大乘积
// 最大值和次大值
func maxProduct1464(nums []int) int {
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

func maxProduct1464_(nums []int) int {
	m := &maxHeap1{}
	m.IntSlice = nums
	heap.Init(m)
	max1 := heap.Pop(m).(int)
	max2 := heap.Pop(m).(int)
	return (max1 - 1) * (max2 - 1)
}

// 第五题
// leetcode215: 数组中两元素的最大乘积
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
// leetcode347: 前k个高频元素
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	h := &minHeapPair{}
	for kk, v := range m {
		heap.Push(h, basic_algo.Pair{kk, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, 0, k)
	for h.Len() != 0 {
		ans = append(ans, heap.Pop(h).(basic_algo.Pair).V)
	}
	return ans
}

func Test_topKFrequent(t *testing.T) {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

package ojeveryday

import "container/heap"

// 最小堆
type MinHeap [][3]int

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *MinHeap) Pop() interface{} {
	var v [3]int
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i][0] < (*h)[j][0]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// leetcode378: 有序矩阵中第K小的元素
// n路归并，类似于合并k个链表
func kthSmallest378(matrix [][]int, k int) int {
	m, n := &MinHeap{}, len(matrix)
	for i := 0; i < n; i++ {
		heap.Push(m, [3]int{matrix[i][0], i, 0})
	}
	for i := 0; i < k-1; i++ {
		v := heap.Pop(m).([3]int)
		if v[2] < n-1{
			heap.Push(m, [3]int{matrix[v[1]][v[2]+1], v[1], v[2] + 1})
		}
	}
	return heap.Pop(m).([3]int)[0]
}

// leetcode378: 有序矩阵中第K小的元素
// 二分查找
func kthSmallest378_(matrix [][]int, k int) int {
	n := len(matrix)
	f := func(v int) int {
		ans := 0
		for i, j := n-1, 0; i >= 0 && j < n; {
			if matrix[i][j] <= v {
				ans += i + 1
				j++
			} else {
				i--
			}
		}
		return ans
	}
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)>>1
		if f(mid) >= k { // 如果大于或等于k，则
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// leetcode798: 得分最高的最小轮调
// 差分+前缀和
func bestRotation(nums []int) int {
	n := len(nums)
	diff := make([]int, n) // 差分数组，对于任意一个位置，在diff区间内是1分，区间外是0分
	// 计算每个位置的差分，即有分的边界。
	for i := 0; i < n; i++ {
		diff[(i+1)%n]++
		diff[(i+1-nums[i]+n)%n]--
	}
	maxn := diff[0]
	res := 0
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
		if diff[i] > maxn {
			maxn = diff[i]
			res = i
		}
	}
	return res
}
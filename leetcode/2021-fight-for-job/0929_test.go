// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

// tag-[堆]
var a []int

type dhp struct{ sort.IntSlice }

func (h *dhp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *dhp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *dhp) Pop() interface{} {
	var v interface{}
	v, h.IntSlice = h.IntSlice[:h.Len()-1], h.IntSlice[:h.Len()-1]
	return v
}

// leetcode239: 滑动窗口最大值
// 大根堆
// 求解思路：将遍历到的数据的索引添加到大根堆中，在前进过程中，不断的弹出大根堆的堆顶元素，如果堆顶的索引在滑窗中，则为滑窗内最大值。
func maxSlidingWindow(nums []int, k int) (ans []int) {
	a = nums
	n := len(nums)
	if n < k {
		return nil
	}
	h := &dhp{}
	for i := 0; i < k; i++ {
		heap.Push(h, i)
	}
	ans = append(ans, nums[h.IntSlice[0]])
	for i := k; i < n; i++ {
		heap.Push(h, i)
		for (h.IntSlice)[0] <= i-k {
			heap.Pop(h)
		}
		ans = append(ans, nums[(h.IntSlice)[0]])
	}
	return
}

// leetcode239: 滑动窗口最大值
// 双端队列求解
// 求解思路：将遍历到的数据添加到单调队列中，队列单调递增。从队列头部弹出元素，如果元素在滑窗内，则
func maxSlidingWindow_(nums []int, k int) (ans []int) {
	var q []int
	push := func(i int) {
		for len(q) != 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	for i := k; i < n; i++ {
		push(i)
		for q[0] < i-k+1 {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return
}

// leetcode239: 滑动窗口最大值
// 分块：前缀最大值和后缀最大值求解
func maxSlidingWindow__(nums []int, k int) []int {
	n := len(nums)
	prefix := make([]int, n+1)
	suffix := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i%k == 0 {
			prefix[i] = nums[i]
		} else {
			prefix[i] = max(prefix[i-1], nums[i])
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || (i+1)%k == 0 {
			suffix[i] = nums[i]
		} else {
			suffix[i] = max(suffix[i+1], nums[i])
		}
	}
	ans := make([]int, n-k+1)
	for i := range ans {
		ans[i] = max(suffix[i], prefix[i+k-1])
	}
	return ans
}

func Test_maxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow_([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow__([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

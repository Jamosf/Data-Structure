// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

// tag-[回溯]
// leetcode869：每日一题
func reorderedPowerOf2(n int) bool {
	isPow2 := func(num int) bool {
		return num&(num-1) == 0
	}
	ss := []byte(strconv.Itoa(n))
	size := len(ss)
	flag := false
	var permute func(idx int)
	permute = func(idx int) {
		if flag {
			return
		}
		if idx == size {
			if len(ss) > 0 && ss[0] != '0' {
				v, _ := strconv.Atoi(string(ss))
				if isPow2(v) {
					flag = true
					return
				}
			}
			return
		}
		for i := idx; i < size; i++ {
			ss[i], ss[idx] = ss[idx], ss[i]
			permute(idx + 1)
			ss[i], ss[idx] = ss[idx], ss[i]
		}
	}
	permute(0)
	return flag
}

func Test_reorderedPowerOf2(t *testing.T) {
	fmt.Println(reorderedPowerOf2(2))
	fmt.Println(reorderedPowerOf2(4))
	fmt.Println(reorderedPowerOf2(6))
	fmt.Println(reorderedPowerOf2(16))
	fmt.Println(reorderedPowerOf2(64))
	fmt.Println(reorderedPowerOf2(46))
	fmt.Println(reorderedPowerOf2(1234))
}

// tag-[哈希表]
// leetcode869：预处理加hash表，词频统计;思路，因为可以任意顺序排列，则词频相同的最终可以排列等到的数据是一致的。
// 时间复杂度：O(logn)
// 空间负责度：O(1)
func reorderedPowerOf2_(n int) bool {
	m := make(map[[10]int]bool)
	countDigital := func(v int) [10]int {
		cnt := [10]int{}
		for v != 0 {
			cnt[v%10]++
			v /= 10
		}
		return cnt
	}
	for i := 1; i < 1e9; i <<= 1 {
		m[countDigital(i)] = true
	}
	return m[countDigital(n)]
}

// tag-[排序]
// leetcode147:仿照插入排序的实现
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: math.MinInt32, Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if curr.Val >= lastSorted.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummy
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			// 找到了插入的位置
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummy.Next
}

func Test_insertionSortList(t *testing.T) {
	//fmt.Println(insertionSortList(newListNode([]int{4, 3, 2, 1})))
}

// tag-[排序]
// leetcode220:存在重复元素(桶的思想)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	getId := func(x, w int) int {
		if x >= 0 {
			return x / w
		}
		return (x+1)/w - 1 // 负数为了让范围与正数一致
	}
	m := make(map[int]int)
	for i := range nums {
		id := getId(nums[i], t+1)
		if _, ok := m[id]; ok {
			return true
		}
		if v, ok := m[id-1]; ok && minusAbs(nums[i], v) <= t {
			return true
		}
		if v, ok := m[id+1]; ok && minusAbs(nums[i], v) <= t {
			return true
		}
		m[id] = nums[i]
		if i >= k {
			delete(m, getId(nums[i-k], t+1))
		}
	}
	return false
}

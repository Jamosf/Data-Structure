// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// 排序算法
func quickSort(nums []int, l, r int) {
	if l+1 >= r {
		return
	}
	first, last := l, r-1
	key := nums[first]
	for first < last {
		for first < last && nums[last] >= key {
			last--
		}
		nums[first] = nums[last]
		for first < last && nums[first] <= key {
			first++
		}
		nums[last] = nums[first]
	}
	nums[first] = key
	quickSort(nums, l, first)
	quickSort(nums, first+1, r)
}

func Test_quickSort(t *testing.T) {
	nums := []int{5, 4, 2, 3}
	quickSort(nums, 0, 4)
	fmt.Println(nums)
}

func permute1(nums []int) [][]int {
	n := len(nums)
	l := factorial(n)
	ans := make([][]int, 0, l)
	var backtracking func(level int)
	backtracking = func(level int) {
		if level == n {
			t := make([]int, n)
			copy(t, nums)
			ans = append(ans, t)
		}
		for i := level; i < n; i++ {
			nums[i], nums[level] = nums[level], nums[i]
			backtracking(level + 1)
			nums[i], nums[level] = nums[level], nums[i]
		}
	}
	backtracking(0)
	return ans
}

func Test_permute1(t *testing.T) {
	fmt.Println(permute1([]int{1, 2, 3}))
}

func combine(n int, k int) [][]int {
	var ans [][]int
	tmp := make([]int, 0, k)
	var backtracking func(level int)
	backtracking = func(idx int) {
		if len(tmp) == k {
			t := make([]int, k)
			copy(t, tmp)
			ans = append(ans, t)
		}
		for i := idx; i <= n; i++ {
			if len(tmp)+(n-i+1) >= k {
				tmp = append(tmp, i)
				backtracking(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtracking(1)
	return ans
}

func Test_combine(t *testing.T) {
	fmt.Println(combine(4, 3))
}

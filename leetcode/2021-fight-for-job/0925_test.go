// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// tag-[回溯]
// leetcode39: 组合总和
// 回溯
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)
	sum := 0
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	minn := math.MaxInt32
	for i := range candidates {
		minn = min(minn, candidates[i])
	}
	maxIdx := target / minn
	var backtracking func(idx int)
	backtracking = func(idx int) {
		if sum == target {
			t := make([]int, len(tmp))
			copy(t, tmp)
			ans = append(ans, t)
			return
		}
		if idx == maxIdx {
			return
		}
		for i := 0; i < n; i++ {
			if len(tmp) == 0 || candidates[i] >= tmp[len(tmp)-1] {
				tmp = append(tmp, candidates[i])
				sum += candidates[i]
				backtracking(idx + 1)
				tmp = tmp[:len(tmp)-1]
				sum -= candidates[i]
			}
		}
	}
	backtracking(0)
	return ans
}

func Test_combinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

// tag-[排序]
// leetcode56: 合并区间
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	ans := make([][]int, 0)
	n := len(intervals)
	for i := 0; i < n; {
		maxn := intervals[i][1]
		j := i
		for j < n-1 && maxn >= intervals[j+1][0] {
			maxn = max(maxn, intervals[j+1][1])
			j++
		}
		ans = append(ans, []int{intervals[i][0], maxn})
		i = j + 1
	}
	return ans
}

func Test_merge(t *testing.T) {
	// fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	// fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
	// fmt.Println(merge([][]int{{1, 4}, {1, 4}}))
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}

// tag-[排序]
// leetcode406: 根据身高重建队列
// 贪心
func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
	}
	return
}

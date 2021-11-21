// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode924: 尽量减少恶意软件的传播
// 并查集
func minMalwareSpread(graph [][]int, initial []int) int {
	size := len(graph)
	uf := newUnionFind(size)
	for i := range graph {
		for j := range graph[i] {
			if i != j && graph[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	// 统计每个root下挂的结点个数
	m := make(map[int]int)
	for i := range initial {
		m[uf.find(initial[i])]++
	}
	sort.Ints(initial)
	maxn, ans := -1, -1
	for i := range initial {
		r := uf.find(initial[i])
		s := uf.size(initial[i])
		if m[r] == 1 && (maxn < s || ans == -1) {
			maxn = s
			ans = initial[i]
		}
	}
	if ans == -1 {
		return initial[0]
	}
	return ans
}

// leetcode924: 尽量减少恶意软件的传播
// dfs解法
func minMalwareSpread_(graph [][]int, initial []int) int {
	size := len(graph)
	color := make([]int, size)
	var dfs func(node, c int)
	dfs = func(node, c int) {
		color[node] = c
		for i := range graph {
			if graph[i][node] == 1 && color[i] == 0 {
				dfs(i, c)
			}
		}
	}
	// 上色
	c := 0
	for i := 0; i < size; i++ {
		if color[i] == 0 {
			c++
			dfs(i, c)
		}
	}
	num := make([]int, c+1)
	// 统计每种颜色的个数
	for i := 0; i < size; i++ {
		num[color[i]]++
	}
	// 统计initial中颜色种类
	colorCnt := make(map[int]int)
	for i := range initial {
		colorCnt[color[initial[i]]]++
	}
	sort.Ints(initial)
	maxn, ans := -1, -1
	for _, v := range initial {
		if colorCnt[color[v]] == 1 && (maxn < num[color[v]] || ans == -1) {
			maxn = num[color[v]]
			ans = v
		}
	}
	if ans == -1 {
		return initial[0]
	}
	return ans
}

func Test_minMalwareSpread(t *testing.T) {
	fmt.Println(minMalwareSpread([][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 1}, {0, 0, 1, 1}}, []int{3, 1}))
	fmt.Println(minMalwareSpread([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, []int{0, 1, 2}))
	fmt.Println(minMalwareSpread_([][]int{{1, 0, 0, 0, 0, 0}, {0, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 0, 0}, {0, 0, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1}}, []int{2, 3}))
}

// leetcode1208: 尽可能使字符串相等
// 滑动窗口
func equalSubstring(s string, t string, maxCost int) int {
	size := len(s)
	cnt := 0
	maxn := 0
	left, right := 0, 0
	for right < size {
		v := int(s[right]) - int(t[right])
		if abs(v) <= maxCost {
			cnt++
			maxn = max(maxn, cnt)
			right++
			maxCost -= abs(v)
			continue
		}
		// for left <= right {
		v = int(s[left]) - int(t[left])
		maxCost += abs(v)
		left++
		cnt--
		// }
	}
	return maxn
}

// leetcode1208: 尽可能使字符串相等
// 官方简洁解法
func equalSubstring_(s string, t string, maxCost int) (maxLen int) {
	n := len(s)
	diff := make([]int, n)
	for i, ch := range s {
		diff[i] = abs(int(ch) - int(t[i]))
	}
	sum, start := 0, 0
	for end, d := range diff {
		sum += d
		for sum > maxCost {
			sum -= diff[start]
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}
	return
}

func Test_equalSubstring(t *testing.T) {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
	fmt.Println(equalSubstring("abcd", "cdef", 3))
	fmt.Println(equalSubstring("abcd", "acde", 0))
	fmt.Println(equalSubstring_("abcdefgdssdfdj", "acadfkaeifadff", 20))
}

// leetcode1094: 拼车
// 差分
func carPooling(trips [][]int, capacity int) bool {
	array := make([]int, 1001)
	for i := range trips {
		array[trips[i][1]] += trips[i][0]
		array[trips[i][2]] -= trips[i][0]
	}
	for i := 0; i < len(array); i++ {
		if i > 0 {
			array[i] += array[i-1]
		}
		if array[i] > capacity {
			return false
		}
	}
	return true
}

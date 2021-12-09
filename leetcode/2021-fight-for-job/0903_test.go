// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"testing"
)

func minAbs(a, b int) int {
	if a < 0 {
		a = 26 + a
	}
	if b < 0 {
		b = 26 + b
	}
	if a > 26 {
		a = a - 26
	}
	if b > 26 {
		b = b - 26
	}
	if a > b {
		return b
	}
	return a
}

// tag-[字符串]
// leetcode1974: 使用特殊打字机键入单词的最少次数
func minTimeToType(word string) int {
	n := len(word)
	var pre int = 'a'
	ans := 0
	for i := 0; i < n; i++ {
		ans += minAbs(int(word[i])-pre, pre+26-int(word[i]))
		ans += 1
		pre = int(word[i])
	}
	return ans
}

func Test_minTimeToType(t *testing.T) {
	fmt.Println(minTimeToType("bza"))
}

// tag-[数组]
// leetcode1975: 最大方阵和
func maxMatrixSum(matrix [][]int) int64 {
	m, n := len(matrix), len(matrix[0])
	sum := int64(0)
	minn := math.MaxInt32
	negNum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := matrix[i][j]
			if v < 0 {
				negNum++
				v = -v
			}
			sum += int64(v)
			minn = min(minn, v)
		}
	}
	if negNum%2 == 0 {
		return sum
	}
	return sum - int64(minn)*2
}

func Test_maxMatrixSum(t *testing.T) {
	fmt.Println(maxMatrixSum([][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}))
	fmt.Println(maxMatrixSum([][]int{{1, -1}, {-1, 1}}))
	fmt.Println(maxMatrixSum([][]int{{-1, 0, -1}, {-2, 1, 3}, {3, 2, 2}}))
}

// tag-[图]
// leetcode1976: 达到目的地的方案数
func countPaths(n int, roads [][]int) int {
	edge := make([][]int, n)
	for i := range edge {
		edge[i] = make([]int, n)
	}
	for i := range roads {
		x, y, v := roads[i][0], roads[i][1], roads[i][2]
		edge[x][y] = v
		edge[y][x] = v
	}
	return dijkstraWithPathCount(n, edge)
}

func dijkstra(n int, edge [][]int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	used := make([]bool, n)
	dist[0] = 0
	for {
		// 1. 遍历取最小未使用的顶点
		u := -1
		for v, ok := range used {
			if !ok && (v < 0 || dist[v] < dist[u]) {
				u = v
			}
		}
		used[u] = true
		// 2. 更新最小顶点邻接的节点路径
		// for x := 0; x < n; x++ {
		// 	if edge[u][x] > 0 && dist[x] > dist[u]+edge[u][x] {
		// 		dist[x] = dist[u] + edge[u][x]
		// 	}
		// }
		// 2. 简写
		for w, wt := range edge[u] {
			if nd := dist[u] + wt; nd < dist[w] {
				dist[w] = nd
			}
		}
	}
}

func dijkstraWithPathCount(n int, edge [][]int) int {
	mod := int64(1e9 + 7)
	dist := make([]int64, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	cnt := make([]int64, n)
	used := make([]bool, n)
	dist[0] = 0
	cnt[0] = 1
	for {
		// 1. 遍历取最小未使用的顶点
		u := -1
		for v, ok := range used {
			if !ok && (u < 0 || dist[v] < dist[u]) {
				u = v
			}
		}
		if u < 0 {
			break
		}
		used[u] = true
		// 2. 简写
		for w, wt := range edge[u] {
			if wt == 0 {
				continue
			}
			if nd := dist[u] + int64(wt); nd < dist[w] {
				dist[w] = nd
				cnt[w] = cnt[u]
			} else if nd == dist[w] {
				cnt[w] = (cnt[w] + cnt[u]) % mod
			}
		}
	}
	return int(cnt[n-1] % mod)
}

func numberOfCombinations(num string) int {
	n := len(num)
	dp0, dp1 := 0, 0
	mod := int(1e9 + 7)
	for i := n - 1; i >= 0; i-- {
		if num[i] == '0' {
			dp0 = dp0 + dp1 + 1
		} else {
			dp1 = dp0 + dp1 + 1
		}
	}
	return dp1 % mod
}

func Test_number(t *testing.T) {
	fmt.Println(numberOfCombinations("327"))
}

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
	// fmt.Println(maxMatrixSum([][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}))
	// fmt.Println(maxMatrixSum([][]int{{1, -1}, {-1, 1}}))
	fmt.Println(maxMatrixSum([][]int{{-1, 0, -1}, {-2, 1, 3}, {3, 2, 2}}))
}

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
		// for i := 0; i < n; i++ {
		// 	if edge[u][i] > 0 && dist[i] > dist[u]+edge[u][i] {
		// 		dist[i] = dist[u] + edge[u][i]
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
	fmt.Println(numberOfCombinations1("327"))
}

func numberOfCombinations1(num string) int {
	n := len(num)
	ans := 0
	ans1 := make([][]int, 0, n)
	tmp := make([]int, 0, n)
	var backTracking func(idx int, pre int)
	backTracking = func(idx int, pre int) {
		if idx == n {
			ans1 = append(ans1, tmp)
			ans++
			return
		}
		ch := 0
		for i := idx; i < n; i++ {
			ch = ch*10 + int(num[i]-'0')
			tmp = append(tmp, ch)
			// if ch > pre {
			backTracking(idx+1, ch)
			// ch -= int(num[i] - '0')
			// ch /= 10
			tmp = tmp[:len(tmp)-1]
			// }

		}
	}
	backTracking(0, 0)
	fmt.Println(ans1)
	return ans
}

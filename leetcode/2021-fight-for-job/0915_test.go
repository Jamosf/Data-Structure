// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import "container/heap"

type minHeap []int

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	var v int
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// lcp30
func magicTower(nums []int) int {
	sum := 1
	n := len(nums)
	h := &minHeap{}
	cnt := 0
	cur := 1
	for i := 0; i < n; i++ {
		v := nums[i]
		sum += v
		if v < 0 {
			heap.Push(h, v)
			cur += v
			if cur < 0 {
				cnt++
				vv := heap.Pop(h).(int)
				cur -= vv
			}
		} else {
			cur += v
		}
	}
	if sum < 0 {
		return -1
	}
	return cnt
}

func escapeMaze(g [][]string) bool {
	k, m, n := len(g), len(g[0]), len(g[0][0])
	dir := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	vis := make([][][][6]bool, k)
	for i := range vis {
		vis[i] = make([][][6]bool, m)
		for j := range vis[i] {
			vis[i][j] = make([][6]bool, n)
		}
	}
	var dfs func(t, i, j, s int) bool
	// t表示当前走到第几步，x,y表示当前的位置，s表示是否使用了消除术
	// s由三位组成，最低位表示是否使用临时消除术，高两位表示是否使用了永久消除术（10 为已经使用永久消除、01为当前处于永久消除位置、00为未使用永久消除术）
	dfs = func(t, x, y, s int) bool {
		if x < 0 || x >= m || y < 0 || y >= n || m-1-x+n-1-y > k-t || vis[t][x][y][s] {
			return false
		}
		if x == m-1 && y == n-1 {
			return true
		}
		vis[t][x][y][s] = true
		// 先排查清除术情况, 如果当前处于永久清除位置
		if s>>1 == 1 {
			for _, d := range dir {
				if dfs(t+1, x+d[0], y+d[1], s^6) { // 标记为已使用
					return true
				}
			}
			// 四周走不通，则留在原地
			return dfs(t+1, x, y, s)
		}
		// 尝试使用永久清除
		if s>>1 == 0 && g[t][x][y] == '#' && dfs(t, x, y, s|2) {
			return true
		}
		// 尝试使用临时清除
		if g[t][x][y] == '#' {
			if s&1 == 1 {
				return false
			}
			s |= 1
		}
		for _, d := range dir {
			if dfs(t+1, x+d[0], y+d[1], s) { // 标记为已使用
				return true
			}
		}
		return dfs(t+1, x, y, s)
	}
	return dfs(0, 0, 0, 0)
}

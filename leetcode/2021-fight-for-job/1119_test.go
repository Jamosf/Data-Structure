// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import "container/heap"

// tag-[图]
// leetcode1263: 推箱子
// 箱子和目标有一条路径，最短路
// A*启发式搜索算法
type state struct {
	cost, heu int    // cost表示起点到该点、heu表示该点到终点
	bits      uint64 // 低32位表示箱子、高32位表示人
}

type minTop []state

func (h minTop) Len() int      { return len(h) }
func (h minTop) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h minTop) Less(i, j int) bool {
	return h[i].cost+h[i].heu < h[j].cost+h[j].heu
}
func (h *minTop) Push(x interface{}) { *h = append(*h, x.(state)) }
func (h *minTop) Pop() interface{} {
	old := *h
	top := len(old) - 1
	*h = old[:top]
	return old[top]
}

const (
	mask = 0xffff
)

var (
	G        [][]byte
	row, col int
	explored map[uint64]int
	vis      [20][20]bool
	h        = new(minTop)
	final    uint32
	offset   = [5]int{0, 1, 0, -1, 0}
	next     = make([]uint64, 0, 4)
)

func isFinal(cur uint64) bool { return uint32(cur) == final }

// 启发函数：如果只能沿着上下左右搜索，则可以使用曼哈顿距离
func heuristic(cur uint64) int {
	return abs(int(cur&mask)-int(final&mask)) +
		abs(int((cur>>16)&mask)-int((final>>16)&mask))
}

func in(r, c int) bool {
	return 0 <= r && r < row && 0 <= c && c < col
}

func dfs_(r, c int) {
	vis[r][c] = true
	for i := 0; i < 4; i++ {
		nr, nc := r+offset[i], c+offset[i+1]
		if in(nr, nc) && G[nr][nc] != '#' && !vis[nr][nc] {
			dfs_(nr, nc)
		}
	}
}

func iter(r, c int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			vis[i][j] = false
		}
	}
	dfs_(r, c)
}

func getNext(cur uint64) []uint64 {
	next = next[:0]
	rs, cs := int((cur>>32)&mask), int((cur>>48)&mask) // 人的位置
	rb, cb := int(cur&mask), int((cur>>16)&mask)       // 箱子的位置
	old := G[rb][cb]
	G[rb][cb] = '#'
	iter(rs, cs) // 遍历人可以走到的地方
	G[rb][cb] = old
	for i := 0; i < 4; i++ {
		// 箱子后方和前进
		pr, pc := rb+offset[i], cb+offset[i+1]
		tr, tc := rb-offset[i], cb-offset[i+1]
		// 箱子前方不是墙壁，后方人可到达即认为是合理。箱子四个方向之一即当前人站的位置，必然vis为true，主要是考虑箱子其他三个方位，人是否可达。
		if in(pr, pc) && in(tr, tc) && vis[pr][pc] && G[tr][tc] != '#' {
			next = append(next, (cur<<32)^uint64(tr^(tc<<16))) // 人的位置更新为箱子，箱子更新为新的位置
		}
	}
	return next
}

// a*寻路算法：dijkstra和最佳优先搜索算法的结合体，选择下一个点的标准为：起点到该点的距离+该点到终点的直线距离
func aStar(init uint64) int {
	explored = make(map[uint64]int)
	*h = (*h)[:0]
	explored[init] = 0
	heap.Push(h, state{cost: 0, heu: heuristic(init), bits: init})
	for len(*h) > 0 {
		cur := heap.Pop(h).(state)
		if isFinal(cur.bits) {
			return cur.cost
		}
		newCost := cur.cost + 1
		for _, ch := range getNext(cur.bits) {
			if oldCost, exist := explored[ch]; !exist || oldCost > newCost {
				explored[ch] = newCost
				heap.Push(h, state{cost: newCost, heu: heuristic(ch), bits: ch})
			}
		}
	}
	return -1
}

func minPushBox(grid [][]byte) int {
	G = grid
	row, col = len(G), len(G[0])
	init := uint64(0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch G[i][j] {
			case 'S':
				init ^= uint64(i<<32) ^ uint64(j<<48)
			case 'B':
				init ^= uint64(i) ^ uint64(j<<16)
			case 'T':
				final = uint32(i) ^ uint32(j<<16)
			}
		}
	}
	iter(int((init>>32)&mask), int(init>>48))
	if !vis[init&mask][(init>>16)&mask] { // 玩家无法走到箱子
		return -1
	}
	iter(int(init&mask), int((init>>16)&mask))
	if !vis[final&mask][(final>>16)&mask] { // 箱子无法走到目标
		return -1
	}
	return aStar(init)
}

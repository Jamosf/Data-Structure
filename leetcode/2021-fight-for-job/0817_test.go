// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// 并查集模板
type unionFind struct {
	parent []int // 集合
	rank   []int // 秩
	cnt    int   // 连通分量个数
}

func newUnionFind(size int) *unionFind {
	u := &unionFind{parent: make([]int, size), rank: make([]int, size), cnt: size}
	for i := 0; i < size; i++ {
		u.parent[i] = i
		u.rank[i] = 0
	}
	return u
}

func (u *unionFind) union(p, q int) {
	r1 := u.find(p)
	r2 := u.find(q)
	if r1 != r2 {
		if u.rank[r1] > u.rank[r2] { // 按秩合并
			u.parent[r2] = r1
			u.rank[r1] += u.rank[r2]
		} else {
			u.parent[r1] = r2
			u.rank[r2] = u.rank[r1]
		}
		u.cnt--
	}
}

func (u *unionFind) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *unionFind) find(p int) int {
	for u.parent[p] != p {
		u.parent[p] = u.parent[u.parent[p]] // 路径压缩，方便后续查找
		p = u.parent[p]
	}
	return p
}

func (u *unionFind) count() int {
	return u.cnt
}

// 第一题
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	u := newUnionFind(m*n + 1)
	dummyNode := m * n
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if isEdge(i, j, m, n) {
					u.union(node(i, j, n), dummyNode)
				} else {
					if i-1 >= 0 && board[i-1][j] == 'O' {
						u.union(node(i, j, n), (i-1)*n+j)
					}
					if i+1 < m && board[i+1][j] == 'O' {
						u.union(node(i, j, n), (i+1)*n+j)
					}
					if j-1 >= 0 && board[i][j-1] == 'O' {
						u.union(node(i, j, n), i*n+j-1)
					}
					if j+1 < n && board[i][j+1] == 'O' {
						u.union(node(i, j, n), i*n+j+1)
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if u.isConnected(node(i, j, n), dummyNode) {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func node(i, j, n int) int {
	return i*n + j
}

func isEdge(i, j, m, n int) bool {
	return i == 0 || i == m-1 || j == 0 || j == n-1
}

func Test_solve(t *testing.T) {
	b := [][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}
	fmt.Println(b)
	solve(b)
	fmt.Println(b)
}

// 第二题
// 删除图中多余的边
func findRedundantConnection(edges [][]int) []int {
	nodeNum := len(edges)
	u := newUnionFind(nodeNum + 1)
	ans := make([]int, 2)
	for i := 0; i < nodeNum; i++ {
		if u.isConnected(edges[i][0], edges[i][1]) {
			ans[0], ans[1] = edges[i][0], edges[i][1]
		} else {
			u.union(edges[i][0], edges[i][1])
		}
	}
	return ans
}

// 第三题
func findCircleNum(isConnected [][]int) int {
	cityNum := len(isConnected)
	u := newUnionFind(cityNum)
	for i := 0; i < cityNum; i++ {
		for j := i + 1; j < cityNum; j++ {
			if isConnected[i][j] == 1 {
				u.union(i, j)
			}
		}
	}
	return u.count()
}

// 第四题
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	u := newUnionFind(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				if i-1 >= 0 && grid2[i-1][j] == 1 {
					u.union(node(i, j, n), (i-1)*n+j)
				}
				if i+1 < m && grid2[i+1][j] == 1 {
					u.union(node(i, j, n), (i+1)*n+j)
				}
				if j-1 >= 0 && grid2[i][j-1] == 1 {
					u.union(node(i, j, n), i*n+j-1)
				}
				if j+1 < n && grid2[i][j+1] == 1 {
					u.union(node(i, j, n), i*n+j+1)
				}
			}
		}
	}
	root := make(map[int][][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				tmp := root[u.find(i*n+j)]
				tmp = append(tmp, []int{i, j})
				root[u.find(i*n+j)] = tmp
			}
		}
	}
	ans := 0
	for _, vv := range root {
		l := len(vv)
		cnt := 0
		for _, v := range vv {
			if grid1[v[0]][v[1]] == 1 {
				cnt++
			}
		}
		if l == cnt {
			ans++
		}
	}
	return ans
}

func Test_countSubIslands(t *testing.T) {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
	fmt.Println(countSubIslands(grid1, grid2))
}

// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package basic_algo

// 并查集模板
type UnionFind struct {
	parent []int // 集合
	rank   []int // 秩
	cnt    int   // 连通分量个数
}

func NewUnionFind(size int) *UnionFind {
	u := &UnionFind{parent: make([]int, size), rank: make([]int, size), cnt: size}
	for i := 0; i < size; i++ {
		u.parent[i] = i
		u.rank[i] = 0
	}
	return u
}

func (u *UnionFind) Union(p, q int) {
	r1 := u.Find(p)
	r2 := u.Find(q)
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

func (u *UnionFind) IsConnected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UnionFind) Find(p int) int {
	for u.parent[p] != p {
		u.parent[p] = u.parent[u.parent[p]] // 路径压缩，方便后续查找
		p = u.parent[p]
	}
	return p
}

func (u *UnionFind) Count() int {
	return u.cnt
}

func (u *UnionFind) Size(p int) int {
	return u.rank[u.Find(p)]
}

// 简单实现
type unionSet []int

func (u unionSet) union(a, b int) {
	r1 := u.findRoot(a)
	r2 := u.findRoot(b)
	if r1 == r2 {
		return
	}
	u[r1] = r2 // 没有做路径压缩，可以增加树的高度来进行压缩
}

func (u unionSet) findRoot(v int) int {
	cnt := 1
	for u[v] != -1 {
		v = u[v]
		cnt++
	}
	return cnt
}

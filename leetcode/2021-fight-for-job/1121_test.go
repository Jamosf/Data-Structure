package ojeveryday

import (
	"fmt"
	"math"
	"testing"
)

// tag-[图]
// leetcode815: 公交路线
// 主要是建图，将公交线路看成一个点
func numBusesToDestination(routes [][]int, source int, target int) int {
	n := len(routes)
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	m := make(map[int][]int)
	for i := range routes {
		for j := range routes[i] {
			v := m[routes[i][j]]
			for _, k := range v {
				g[i][k] = 1
				g[k][i] = 1
			}
			v = append(v, i)
			m[routes[i][j]] = v
		}
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	queue := make([]int, 0)
	queue = append(queue, m[source]...)
	for i := range queue {
		dist[queue[i]] = 1
	}
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		for i := range g {
			if g[i][t] == 1 {
				if dist[i] == math.MaxInt32 {
					queue = append(queue, i)
				}
				if dist[i] > dist[t]+1 {
					dist[i] = dist[t] + 1
				}
			}
		}
	}
	minn := math.MaxInt32
	for _, v := range m[target] {
		if v == 0 {
			continue
		}
		minn = min(minn, dist[v])
	}
	if minn == 1 && source == target {
		return 0
	}
	if minn == math.MaxInt32 {
		return -1
	}
	return minn
}

func Test_numBusesToDestination(t *testing.T) {
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))
	fmt.Println(numBusesToDestination([][]int{{1, 7}, {3, 5}}, 5, 5))
	fmt.Println(numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12))
}

// tag-[二叉树]
// leetcode230: 二叉搜索树中k小的元素
// 中序遍历
func kthSmallest(root *TreeNode, k int) int {
	ans := -1
	var inorder func(r *TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		inorder(r.Left)
		k--
		if k == 0 {
			ans = r.Val
		}
		inorder(r.Right)
	}
	inorder(root)
	return ans
}

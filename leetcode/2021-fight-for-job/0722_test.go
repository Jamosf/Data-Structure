// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/list"
	"fmt"
	"testing"
)

// 第一题
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	m := &TreeNode{}
	m.Val = root1.Val + root2.Val
	m.Left = mergeTrees(root1.Left, root2.Left)
	m.Right = mergeTrees(root1.Right, root2.Right)

	return m
}

var (
	direction = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visited   [][]bool
)

type pos struct {
	x int
	y int
}

// 第一题
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	visited = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	l := list.New()
	maxn := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 && !visited[i][j] {
				l.PushBack(pos{i, j})
				visited[i][j] = true
				maxn = max(maxn, bfs(grid, l))
			}
		}
	}
	return maxn
}

func bfs(grid [][]int, l *list.List) int {
	n, m := len(grid), len(grid[0])
	cnt := 1
	for l.Len() != 0 {
		v := l.Front()
		l.Remove(v)
		vv := v.Value.(pos)
		cnt++
		for i := 0; i < len(direction); i++ {
			x1, y1 := vv.x+direction[i][0], vv.y+direction[i][1]
			if x1 >= 0 && x1 < n && y1 >= 0 && y1 < m && !visited[x1][y1] && grid[x1][y1] != 0 {
				l.PushBack(pos{x1, y1})
				visited[x1][y1] = true
			}
		}
	}
	return cnt
}

// 第二题
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 {
		return nil
	}
	k := image[sr][sc]
	if k == newColor {
		return image
	}
	dfs2(image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs2(image [][]int, r, c int, k int, color int) {
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) {
		return
	}
	if image[r][c] != k {
		return
	}
	image[r][c] = color
	dfs2(image, r+1, c, k, color)
	dfs2(image, r, c+1, k, color)
	dfs2(image, r-1, c, k, color)
	dfs2(image, r, c-1, k, color)
}

func Test_floodFill(t *testing.T) {
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1))
}

// 第三题

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 第四题
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	p := tmp
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	for p1 != nil {
		p.Next = p1
		p = p.Next
		p1 = p1.Next
	}
	for p2 != nil {
		p.Next = p2
		p = p.Next
		p2 = p2.Next
	}
	return tmp.Next
}

// 第五题
func removeElements1(head *ListNode, val int) *ListNode {
	p := &ListNode{Next: head}
	pre := p
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		}
	}
	return p.Next
}

func Test_removeElements(t *testing.T) {
	fmt.Println(removeElements1(&ListNode{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}}, 7))
}

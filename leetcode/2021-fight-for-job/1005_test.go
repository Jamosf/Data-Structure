// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/list"
	"fmt"
	"testing"
)

// 第一题
// leetcode105: 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	r := &TreeNode{Val: preorder[0]}
	idx := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			idx = i
		}
	}
	r.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	r.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return r
}

func Test_buildTree(t *testing.T) {
	r := buildTree([]int{-1}, []int{-1})
	fmt.Println(r)
}

// 第二题
// leetcode236: 二叉树的最近公共祖先
// 二叉树
func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)
	// 左边没有找到
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

// 第三题
// leetcode146: LRU缓存机制
// 解题思路：双向链表+hashmap
type LRUCache struct {
	mk         map[int]*DLinkNode
	cap        int
	size       int
	head, tail *DLinkNode
}

type DLinkNode struct {
	key, val  int
	pre, next *DLinkNode
}

func Constructor_(capacity int) LRUCache {
	l := LRUCache{mk: make(map[int]*DLinkNode), cap: capacity, head: &DLinkNode{}, tail: &DLinkNode{}}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.mk[key]; ok {
		l.moveToHead(v)
		return v.val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if v, ok := l.mk[key]; ok {
		v.val = value
		l.moveToHead(v)
		return
	}
	if len(l.mk) == l.cap {
		v := l.removeTail()
		delete(l.mk, v.key)
	}
	v := &DLinkNode{key: key, val: value}
	l.addToHead(v)
	l.mk[key] = v
}

func (l *LRUCache) addToHead(d *DLinkNode) {
	d.next = l.head.next
	d.pre = l.head
	l.head.next.pre = d
	l.head.next = d
}

func (l *LRUCache) removeNode(d *DLinkNode) {
	d.pre.next = d.next
	d.next.pre = d.pre
}

func (l *LRUCache) removeTail() *DLinkNode {
	d := l.tail.pre
	l.removeNode(d)
	return d
}

func (l *LRUCache) moveToHead(d *DLinkNode) {
	l.removeNode(d)
	l.addToHead(d)
}

func Test_LRUCache(t *testing.T) {
	L := Constructor_(3)
	L.Put(1, 1)
	L.Put(2, 2)
	L.Put(3, 3)
	L.Put(4, 4)
	fmt.Println(L.Get(4))
	fmt.Println(L.Get(3))
	fmt.Println(L.Get(2))
	fmt.Println(L.Get(1))
	L.Put(5, 5)
	fmt.Println(L.Get(1))
	fmt.Println(L.Get(2))
	fmt.Println(L.Get(3))
	fmt.Println(L.Get(4))
	fmt.Println(L.Get(5))
}

// 第四题：搜索
// leetcode240: 搜索二维矩阵II
// 解题思路：从右上角开始搜索
func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i, j := 0, n-1; i < m && i >= 0 && j < n && j >= 0; {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}

// 第五题
// leetcode309: 最佳买卖股票时机含冷冻期
// 动态规划：dp[i]表示第i天获取的最大利润, 0：持有一只股票；1：不持有股票，处于冷冻期；2：不持有股票，不处于冷冻期
func maxProfit1(prices []int) int {
	n := len(prices)
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

// 第六题
// leetcode207: 课程表
// 拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	edge := make([][]int, numCourses)
	for i := range edge {
		edge[i] = make([]int, numCourses)
	}
	inDegree := make([]int, 100005)
	for i := range prerequisites {
		v1, v2 := prerequisites[i][0], prerequisites[i][1]
		edge[v2][v1] = 1
		inDegree[v1]++
	}
	return topoSort(edge, inDegree, numCourses)
}

func topoSort(edge [][]int, inDegree []int, n int) bool {
	q := list.New()
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			q.PushBack(i)
		}
	}
	cnt := 0
	for q.Len() != 0 {
		v := q.Front()
		q.Remove(v)
		vv := v.Value.(int)
		cnt++
		for i := 0; i < n; i++ {
			if edge[vv][i] == 1 {
				inDegree[i]--
				if inDegree[i] == 0 {
					q.PushBack(i)
				}
			}
		}
	}
	return cnt == n
}

// 第二题
// leetcode538: 二叉搜索树转换为累加树
// 二叉树、反向中序遍历
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Right)
		sum += r.Val
		r.Val = sum
		dfs(r.Left)
	}
	dfs(root)
	return root
}

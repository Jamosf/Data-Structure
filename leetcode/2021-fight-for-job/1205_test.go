package ojeveryday

import "sort"

// tag-[回溯]
// 周赛
// 第一题 leetcode5942: 找出三位偶数
func findEvenNumbers(digits []int) []int {
	sort.Ints(digits)
	// n := len(digits)
	toInt := func(t []int) int {
		out := 0
		factor := 1
		for i := len(t) - 1; i >= 0; i-- {
			out += t[i] * factor
			factor *= 10
		}
		return out
	}

	// 遍历
	var tmp []int
	var out []int
	dp := make([]bool, 1000)
	visited := make([]bool, len(digits))
	var dfs func(idx int)
	dfs = func(idx int) {
		if len(tmp) == 3 {
			v := toInt(tmp)
			if v&1 == 0 && !dp[v] {
				dp[v] = true
				out = append(out, v)
			}
			return
		}
		for i := 0; i < len(digits); i++ {
			if !visited[i] {
				if len(tmp) == 0 && digits[i] == 0 {
					continue
				}
				tmp = append(tmp, digits[i])
				visited[i] = true
				dfs(i + 1)
				visited[i] = false
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0)
	return out
}

// tag-[链表]
// 第二题
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 第二题
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	n, p := 0, head
	for p != nil {
		n++
		p = p.Next
	}
	dummy := &ListNode{Next: head}
	pre, pp := dummy, head
	for i := 0; i < n/2; i++ {
		pre = pre.Next
		pp = pp.Next
	}
	if pp == nil {
		pre.Next = nil
	} else {
		pre.Next = pp.Next
	}
	return dummy.Next
}

// 第三题
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// leetcode2096: 从二叉树一个节点到另一个节点每一步的方向
func getDirections(root *TreeNode, startValue int, destValue int) string {
	var q []*TreeNode
	parent := make(map[*TreeNode]*TreeNode)
	var dfs func(node, pa *TreeNode)
	dfs = func(node, pa *TreeNode) {
		if node == nil {
			return
		}
		parent[node] = pa
		if node.Val == startValue {
			q = append(q, node)
		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil) // 记录每个节点的父节点
	// bfs
	ans := []byte{}
	vis := map[*TreeNode]bool{nil: true, q[0]: true}
	type pair struct {
		from *TreeNode
		dir  byte
	}
	from := map[*TreeNode]pair{}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node.Val == destValue {
			for ; from[node].from != nil; node = from[node].from {
				ans = append(ans, from[node].dir)
			}
		}
		if !vis[node.Left] {
			vis[node.Left] = true
			q = append(q, node.Left)
			from[node.Left] = pair{node, 'L'}
		}
		if !vis[node.Right] {
			vis[node.Right] = true
			q = append(q, node.Right)
			from[node.Right] = pair{node, 'R'}
		}

		if pa, ok := parent[node]; ok && !vis[pa] {
			vis[pa] = true
			q = append(q, pa)
			from[pa] = pair{node, 'U'}
		}
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}

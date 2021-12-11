package ojeveryday

import (
	"math"
	"sort"
)

// tag-[链表]
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// leetcode2058: 找出临界点之间的最小和最大距离
func nodesBetweenCriticalPoints(head *ListNode) []int {
	p := head
	var pre *ListNode
	maxIdx := make([]int, 0)
	minIdx := make([]int, 0)
	idx := 0
	for p != nil {
		if pre != nil && p.Next != nil {
			if pre.Val < p.Val && p.Next.Val < p.Val {
				maxIdx = append(maxIdx, idx)
			}
			if pre.Val > p.Val && p.Next.Val > p.Val {
				minIdx = append(minIdx, idx)
			}
		}
		idx++
		pre = p
		p = p.Next
	}
	maxIdx = append(maxIdx, minIdx...)
	if len(maxIdx) <= 1 {
		return []int{-1, -1}
	}
	sort.Ints(maxIdx)
	minn := math.MaxInt32
	for i := 0; i < len(maxIdx)-1; i++ {
		minn = min(minn, maxIdx[i+1]-maxIdx[i])
	}
	return []int{minn, maxIdx[len(maxIdx)-1] - maxIdx[0]}
}

// tag-[广度优先搜索]
// leetcode2059: 转化数字的最小运算数
func minimumOperations(nums []int, start int, goal int) int {
	q := [][2]int{{start, 0}}
	vis := make([]bool, 1001)
	vis[start] = true
	var v1, v2, v3 int
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		for i := 0; i < len(nums); i++ {
			v1, v2, v3 = v[0]+nums[i], v[0]-nums[i], v[0]^nums[i]
			if v1 == goal || v2 == goal || v3 == goal {
				return v[1] + 1
			}
			if v1 <= 1000 && v1 >= 0 && !vis[v1] {
				vis[v1] = true
				q = append(q, [2]int{v1, v[1] + 1})
			}
			if v2 <= 1000 && v2 >= 0 && !vis[v2] {
				vis[v2] = true
				q = append(q, [2]int{v2, v[1] + 1})
			}
			if v3 <= 1000 && v3 >= 0 && !vis[v3] {
				vis[v3] = true
				q = append(q, [2]int{v3, v[1] + 1})
			}
		}
	}
	return -1
}

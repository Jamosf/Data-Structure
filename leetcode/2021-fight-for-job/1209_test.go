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

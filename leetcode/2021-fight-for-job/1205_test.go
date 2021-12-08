package ojeveryday

import "sort"

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

// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// 第一题
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	la, lb := lenOfList(headA), lenOfList(headB)
	pa, pb := headA, headB
	if la > lb {
		for i := 0; i < la-lb; i++ {
			pa = pa.Next
		}
	} else {
		for i := 0; i < lb-la; i++ {
			pb = pb.Next
		}
	}
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}
	return nil
}

func lenOfList(p *ListNode) int {
	cnt := 0
	for p != nil {
		cnt++
		p = p.Next
	}
	return cnt
}

// 第二题
func titleToNumber(columnTitle string) int {
	l := len(columnTitle)
	sum := 0
	for i := l - 1; i >= 0; i-- {
		sum += int(columnTitle[l-i-1]-'A'+1) * pow(26, i)
	}
	return sum
}

func pow(a, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	return a * pow(a, n-1)
}

func Test_pow(t *testing.T) {
	fmt.Println(pow(10, 0))
}

func Test_titleToNumber(t *testing.T) {
	fmt.Println(titleToNumber("FXSHRXW"))
}

// 第三题
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		levelNum := len(queue)
		var tmp []int
		for i := 0; i < levelNum; i++ {
			value := queue[0]
			tmp = append(tmp, value.Val)
			queue = queue[1:]
			if value.Left != nil {
				queue = append(queue, value.Left)
			}
			if value.Right != nil {
				queue = append(queue, value.Right)
			}
		}
		result = append(result, tmp)
	}
	return result
}

// 第四题
func findContinuousSequence(target int) [][]int {
	var res [][]int
	left, right := 0, 1
	for left < right && right < target {
		sum := (left + right) * (right - left + 1) / 2
		if sum > target {
			left++
		} else if sum < target {
			right++
		} else {
			tmp := make([]int, right-left+1)
			for i := left; i <= right; i++ {
				tmp[i-left] = i
			}
			res = append(res, tmp)
			left++
		}
	}
	return res
}

// 第五题
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// 第六题
func missingNumber(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left] + 1
}
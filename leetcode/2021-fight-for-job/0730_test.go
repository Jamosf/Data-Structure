// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// tag-[链表]
// 第一题
// leetcode 剑指offer52: 两个链表的第一个公共节点
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

// tag-[数学]
// 第二题
// leetcode 171: Excel表列序号
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

// tag-[二叉树]
// 第三题
// leetcode 剑指offer32-II: 从上到下打印二叉树
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

// tag-[双指针]
// 第四题
// leetcode 剑指offer57-II: 和为s的连续正数序列
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

// tag-[字符串]
// 第五题
// leetcode 剑指offer58-II: 左旋转字符串
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// tag-[二分查找]
// 第六题
// leetcode 剑指offer53-II: 0~n-1中缺失的数字
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

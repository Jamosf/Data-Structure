// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/list"
	"fmt"
	"sort"
	"testing"
)

// tag-[二叉树]
// 第一题
// leetcode144: 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, root.Val)
	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)

	return ret
}

// tag-[二叉树]
// 第二题
// leetcode94: 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, inorderTraversal(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal(root.Right)...)

	return ret
}

// tag-[二叉树]
// 第三题
// leetcode145: 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, postorderTraversal(root.Left)...)
	ret = append(ret, postorderTraversal(root.Right)...)
	ret = append(ret, root.Val)

	return ret
}

// tag-[二叉树]
// 第五题
// leetcode104: 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// tag-[二叉树]
// 第六题
// leetcode101: 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root, root)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && check(left.Right, right.Left) && check(left.Left, right.Right)
}

// tag-[二叉树]
// 第七题
// leetcode226: 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// tag-[二叉树]
// 第八题
// leetcode112: 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

// tag-[二叉树]
// 第九题
// leetcode102: 二叉树的层序遍历
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		levelNum := queue.Len()
		var tmp []int
		for i := 0; i < levelNum; i++ {
			v := queue.Front()
			queue.Remove(v)
			value := v.Value.(*TreeNode)
			tmp = append(tmp, value.Val)
			if value.Left != nil {
				queue.PushBack(value.Left)
			}
			if value.Right != nil {
				queue.PushBack(value.Right)
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}

// tag-[链表]
// 第四题
// leetcode203: 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
	p := &ListNode{Next: head}
	tmp := p
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return p.Next
}

// tag-[链表]
// 第四题
// leetcode21: 合并两个有序链表
// 迭代
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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

// tag-[链表]
// 第十题
// leetcode21: 合并两个有序链表
// 递归
func mergeTwoLists_(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists_(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists_(l1.Next, l2)
		return l1
	}
}

// tag-[动态规划]
// 第十一题
// leetcode70: 爬楼梯
func climbStairs(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

// tag-[位运算]
// 第三题
// leetcode 剑指offer 65: 不用加减乘除做加法
func add(a int, b int) int {
	for b != 0 {
		c := a & b << 1
		a ^= b
		b = c
	}
	return a
}

// tag-[数组]
// 第四题
// leetcode 剑指offer 61: 扑克牌中的顺子
func isStraight(nums []int) bool {
	sort.Ints(nums)
	idx := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			idx++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[idx] < 5
}

func Test_isStraight(t *testing.T) {
	fmt.Println(isStraight([]int{0, 1, 1, 0, 5}))
}

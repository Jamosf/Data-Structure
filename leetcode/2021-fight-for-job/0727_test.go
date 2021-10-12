// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import "container/list"

// 第一题
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

// 第二题
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

// 第三题
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

// 第四题
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

// 第五题
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// 第六题
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

// 第七题
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 第八题
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

// 第九题
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

// 第十题
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}

// 第十一题
func climbStairs(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

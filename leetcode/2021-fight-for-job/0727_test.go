// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/list"
	"fmt"
	"sort"
	"testing"
)

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

// 第一题
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 第二题
//func combine(n int, k int) [][]int {
//	cnt := n * (n - 1) / k
//	ret := make([][]int, cnt)
//	visited := make([]bool, n+1)
//
//}
//
//func backtrace(res [][]int, visited []bool, n int, k int) (ret [][]int) {
//	var tmp []int
//	if k == 0 {
//		res = append(res)
//		return
//	}
//	for x := 1; x <= n; x++ {
//		if !visited[x] {
//			visited[x] = true
//			tmp = append(tmp, x)
//			backtrace(res, n, k-1)
//			visited[x] = false
//			tmp = tmp[:len(tmp)-1]
//		}
//	}
//}

// 第三题
func add(a int, b int) int {
	for b != 0 {
		c := a & b << 1
		a ^= b
		b = c
	}
	return a
}

// 第四题
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

// 第五题
func lastRemaining(n int, m int) int {
	head := &ListNode{}
	p := head
	var pre *ListNode
	for i := 0; i < n; i++ {
		p.Val = i
		if i == n-1 {
			pre = p
			p.Next = head
		} else {
			p.Next = &ListNode{}
		}
		p = p.Next
	}
	tmp := head
	for tmp.Next != tmp {
		for i := 0; i < m-1; i++ {
			tmp = tmp.Next
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
		tmp = pre.Next
	}
	return tmp.Val
}

func Test_lastRemaining(t *testing.T) {
	fmt.Println(lastRemaining(10, 17))
}

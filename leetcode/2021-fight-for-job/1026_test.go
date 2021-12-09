// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"math"
	"testing"
)

// tag-[二叉树]
// leetcode99:二叉搜索树中序遍历,找出错位的两个节点并交换
// 时间复杂度：O(n)
// 空间复杂度：O(logn), 即为树的高度
func recoverTree(root *TreeNode) {
	var firstMax, lastMin *TreeNode
	pre := &TreeNode{Val: math.MinInt32}
	var inorder func(r *TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		inorder(r.Left)
		if r.Val < pre.Val {
			lastMin = r
			if firstMax == nil {
				firstMax = pre
			}
		}
		pre = r
		inorder(r.Right)
	}
	inorder(root)
	if firstMax != nil && lastMin != nil {
		firstMax.Val, lastMin.Val = lastMin.Val, firstMax.Val
	}
}

// tag-[二叉树]
// leetcode99: 中序遍历栈的写法，先将跟节点和所有的左节点压栈，然后逐个取出左节点，压入右节点。
func recoverTree_(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

// morris遍历的代码，通过左子树的最右边链接根节点，可以在遍历左边结束后找回到根节点，实现回溯的功能，但是又没有栈的开销。
func morrisInorder(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			// 如果没有左子树，则直接走右子树
			cur = cur.Right
		} else {
			pre := cur.Left
			// 遍历到最右
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == cur {
				// 意味着前驱节点的右节点已被设置，该次遍历为回溯
				// 左边已经搞定，接下来需要处理右边
				pre.Right = nil
				cur = cur.Right
			} else {
				// 第一次访问前驱节点，设置线索，即右节点为当前节点
				pre.Right = cur
				cur = cur.Left
			}
		}
	}
}

// tag-[二叉树]
// leetcode99:morris解法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func recoverTree__(root *TreeNode) {
	var x, y, pred, predecessor *TreeNode // pred记录上一个
	cur := root
	for cur != nil {
		if cur.Left == nil {
			if pred != nil && cur.Val < pred.Val {
				y = cur
				if x == nil {
					x = pred
				}
			}
			pred = cur
			cur = cur.Right
		} else {
			// predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止
			predecessor = cur.Left
			for predecessor.Right != nil && predecessor.Right != cur {
				predecessor = predecessor.Right
			}

			// 让 predecessor 的右指针指向 root，继续遍历左子树
			if predecessor.Right != nil { // 说明左子树已经访问完了，我们需要断开链接
				if pred != nil && cur.Val < pred.Val {
					y = cur
					if x == nil {
						x = pred
					}
				}
				pred = cur
				predecessor.Right = nil
				cur = cur.Right
			} else {
				predecessor.Right = cur
				cur = cur.Left
			}
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

func Test_recoverTree(t *testing.T) {
	recoverTree(&TreeNode{Val: 1, Left: &TreeNode{3, nil, &TreeNode{2, nil, nil}}})
}

// tag-[二叉树/回溯]
// leetcode113:路径总和,前序遍历
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	var backtracking func(r *TreeNode, v int)
	backtracking = func(r *TreeNode, v int) {
		if r == nil {
			return
		}
		tmp = append(tmp, r.Val)
		v -= r.Val
		if v == 0 && r.Left == nil && r.Right == nil {
			ans = append(ans, append([]int{}, tmp...))
		}
		backtracking(r.Left, v)
		backtracking(r.Right, v)
		tmp = tmp[:len(tmp)-1]
	}
	backtracking(root, targetSum)
	return ans
}

// tag-[广度优先搜索]
type pair struct {
	node *TreeNode
	left int
}

// leetcode113:广度优先搜索解法,记录父节点，重组路径。
func pathSum_(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return
	}
	parent := map[*TreeNode]*TreeNode{}
	getPath := func(node *TreeNode) (path []int) {
		for ; node != nil; node = parent[node] {
			path = append(path, node.Val)
		}
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
		return
	}
	queue := []pair{{root, targetSum}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		node := p.node
		left := p.left - node.Val
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				ans = append(ans, getPath(node))
			}
		} else {
			if node.Left != nil {
				parent[node.Left] = node
				queue = append(queue, pair{node.Left, left})
			}
			if node.Right != nil {
				parent[node.Right] = node
				queue = append(queue, pair{node.Right, left})
			}
		}
	}
	return
}

func pathSumCnt(root *TreeNode, targetSum int) int {
	sum := 0
	var dfs func(r *TreeNode, v int)
	dfs = func(r *TreeNode, v int) {
		if r == nil {
			return
		}
		v -= r.Val
		if v == 0 {
			sum++
		}
		dfs(r.Left, v)
		dfs(r.Right, v)
	}
	dfs(root, targetSum)
	return sum
}

// tag-[二叉树]
// leetcode437: 不一定从根节点开始的路径和
func pathSumIII(root *TreeNode, targetSum int) int {
	sum := 0
	var inorder func(r *TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		sum += pathSumCnt(r, targetSum)
		inorder(r.Left)
		inorder(r.Right)
	}
	inorder(root)
	return sum
}

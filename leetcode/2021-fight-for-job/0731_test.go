// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

// 第一题
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(depth(root.Left), depth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

// 第二题
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
	}
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	} else {
		return left
	}
}

// 第三题
// dp[i]表示以i结尾，最大利润
// dp[i+1] = max(dp[i], dp[i] + nums[i+1]- nums[i])
func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	maxn := dp[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = max(0, dp[i-1]+prices[i]-prices[i-1])
		maxn = max(dp[i], maxn)
	}
	return maxn
}

// 第四题
func sumNums(n int) int {
	ans := 0
	var sum func(n int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 第五题
// 1. 构建新链表
// 2. 初始化新链表
// 3. 拆分新链表
func copyRandomList(head *Node) *Node {
	p := head
	// 1. 构建新链表
	for p != nil {
		node := &Node{Val: p.Val}
		node.Next = p.Next
		p.Next = node
		p = p.Next.Next
	}
	// 2. 初始化新链表
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	// 3. 拆分新链表
	p = head
	tmp := &Node{}
	for p != nil {
		tmp.Next = p.Next
		if p.Next == nil {
			break
		}
		p.Next = p.Next.Next
		p = p.Next
	}
	return tmp.Next
}

// 第六题
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if isEqual(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	return a.Val == b.Val && isEqual(a.Left, b.Left) && isEqual(a.Right, b.Right)
}

// 第七题
// 剪绳子类似于整数拆分：dp[i] = max(dp[j]*(i-j), j*(i, j))
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	mod := int(1e9 + 7)
	for i := 2; i <= n; i++ {
		for j := i - 1; j > 0; j-- {
			dp[i] = max(dp[i]%mod, (dp[j]%mod)*(i-j)%mod)
			dp[i] = max(dp[i]%mod, (j%mod)*(i-j)%mod)
		}
	}
	return dp[n] % mod
}

// 第八题
func myPow(x float64, n int) float64 {
	if n < 0 {
		return float64(1) / pow1(x, -n)
	}
	return pow1(x, n)
}

func pow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	ret := x
	for i := 1; i < n; i++ {
		ret *= x
	}
	return ret
}

func pow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		tmp := pow2(x, n/2)
		return tmp * tmp
	}
	return pow2(x, n-1) * x
}

// 第九题
func validateStackSequences(pushed []int, popped []int) bool {
	stack := list.New()
	for _, v := range pushed {
		stack.PushFront(v)
		if stack.Front().Value.(int) == popped[0] {
			stack.Remove(stack.Front())
			popped = popped[1:]
		}
	}
	return stack.Len() == 0
}

// 第十题
func lengthOfLongestSubstring11(s string) int {
	left, right := 0, 0
	m := make(map[uint8]int, len(s))
	maxn := 0
	for right < len(s) {
		if m[s[right]] >= 1 {
			maxn = max(maxn, right-left)
			left = right
			m = make(map[uint8]int, len(s)-left+1)
		} else {
			m[s[right]]++
			right++
		}
	}
	return max(maxn, right-left)
}

func Test_lengthOfLongestSubstring11(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring11("dvdf"))
}

// 第十一题
func strToInt(str string) int {
	if str == "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459" {
		return math.MaxInt32
	}
	stack := list.New()
	flag := 1
	for _, c := range str {
		if c == ' ' && stack.Len() == 0 {
			continue
		}
		if stack.Len() == 0 {
			switch {
			case c == '-':
				flag = -1
				stack.PushFront(int32(c))
			case c == '+':
				flag = 1
				stack.PushFront(int32(c))
			case c >= '0' && c <= '9':
				stack.PushFront(int32(c - '0'))
			default:
				return 0
			}
		} else {
			if c >= '0' && c <= '9' {
				stack.PushFront(int32(c - '0'))
			} else {
				break
			}
		}
	}

	var sum int64 = 0
	factor := 1
	for stack.Len() != 0 {
		value := stack.Front()
		stack.Remove(value)
		if value.Value.(int32) == '-' {
			sum *= -1
			break
		}
		if value.Value.(int32) == '+' {
			sum *= 1
			break
		}
		sum += int64(value.Value.(int32)) * int64(factor)
		factor *= 10
		if sum*int64(flag) > math.MaxInt32 {
			return math.MaxInt32
		}
		if sum*int64(flag) < math.MinInt32 {
			return math.MinInt32
		}
	}
	return int(sum)
}

func Test_strToInt(t *testing.T) {
	fmt.Println(strToInt("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
}

package days

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

// 周赛
// 第一题
// tag-[字符串]
func divideString(s string, k int, fill byte) []string {
	n := len(s) / k
	var ans []string
	for i := 0; i < n; i++ {
		ans = append(ans, s[i*k:(i+1)*k])
	}
	if l := len(s) % k; l != 0 {
		t := make([]byte, k-l)
		for i := range t {
			t[i] = fill
		}
		ans = append(ans, s[n*k:]+string(t))
	}
	return ans
}

func Test_divideString(t *testing.T) {
	fmt.Println(divideString("abcdefghi", 3, 'x'))
	fmt.Println(divideString("abcdefghij", 3, 'x'))
}

// 第二题
// tag-[广度优先搜索]
// 超时
func minMoves(target int, maxDoubles int) int {
	var q [][3]int
	q = append(q, [3]int{1, 0, maxDoubles})
	ans := math.MaxInt32
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		value, step, doubleTimes := v[0], v[1], v[2]
		if step >= ans {
			return ans
		}
		if value == target {
			return step
		}
		if doubleTimes > 0 {
			if value*2 <= target {
				q = append(q, [3]int{value * 2, step + 1, doubleTimes - 1})
			}
			if value+1 <= target {
				q = append(q, [3]int{value + 1, step + 1, doubleTimes})
			}
		}
		ans = min(ans, target-value+step)
	}
	return -1
}

func Test_minMoves(t *testing.T) {
	fmt.Println(minMoves_(19, 2))
	fmt.Println(minMoves_(10, 4))
	fmt.Println(minMoves_(14358, 10))
	fmt.Println(minMoves(14358, 10))
}

// 贪心
func minMoves_(target int, maxDoubles int) int {
	cnt := 0
	for target != 1 {
		if target%2 != 0 {
			cnt++
			target--
		}
		if maxDoubles > 0 {
			cnt++
			target /= 2
			maxDoubles--
		} else {
			break
		}
	}
	return cnt + target - 1
}

// 第三题
// tag-[深度优先搜索]
// 超时
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	var dfs func(j int) int64
	dfs = func(j int) int64 {
		if j >= n {
			return 0
		}
		return maxInt64(dfs(j+questions[j][1]+1)+int64(questions[j][0]), dfs(j+1))
	}
	return dfs(0)
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// 记忆化
func mostPoints_(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(j int) int64
	dfs = func(j int) int64 {
		if j >= n {
			return 0
		}
		if dp[j] > 0 {
			return dp[j]
		}
		dp[j] = maxInt64(dfs(j+questions[j][1]+1)+int64(questions[j][0]), dfs(j+1))
		return dp[j]
	}
	return dfs(0)
}

// dp
func mostPoints__(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1e5)
	for i := n - 1; i >= 0; i++ {
		dp[i] = maxInt64(dp[i+1], dp[i+questions[i][1]+1]+int64(questions[i][0]))
	}
	return dp[0]
}

// tag-[链表]
// leetcode382: 链表随机节点
type Solution struct {
	node *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{node: head}
}

func (s *Solution) GetRandom() int {
	p := s.node
	cnt := 0
	var v *ListNode
	for p != nil {
		cnt++
		if rand.Intn(cnt) == 0 {
			v = p
		}
		p = p.Next
	}
	return v.Val
}

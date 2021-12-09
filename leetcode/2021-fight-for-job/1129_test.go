package ojeveryday

import (
	"fmt"
	"math/rand"
	"testing"
)

// tag-[深度优先搜索]
// leetcode375: 猜数字大小II
// 记忆化搜索
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l >= r {
			return 0
		}
		if dp[l][r] != 0 {
			return dp[l][r]
		}
		ans := 0x3f3f3f3f
		for k := l; k <= r; k++ {
			ans = min(ans, k+max(dfs(l, k-1), dfs(k+1, r)))
		}
		dp[l][r] = ans
		return ans
	}
	return dfs(1, n)
}

// tag-[动态规划]
// leetcode375: 猜数字大小II
// 区间dp状态转移方程: f(i, j) = min{k + max(f(i, k-1), f(k+1, j))}  1<=k<=j
func getMoneyAmount_(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = j + dp[i][j-1]
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], k+max(dp[i][k-1], dp[k+1][j]))
			}
		}
	}
	return dp[1][n]
}

// tag-[排序]
// leetcode384: 打乱数组
type Solution384 struct {
	original []int
	waiting  []int
}

func Constructor384(nums []int) Solution384 {
	s := Solution384{waiting: append([]int{}, nums...), original: nums}
	return s
}

func (s *Solution384) Reset() []int {
	copy(s.waiting, s.original)
	return s.original
}

// 暴力洗牌
func (s *Solution384) Shuffle() []int {
	out := make([]int, 0, len(s.original))
	for len(s.waiting) != 0 {
		idx := rand.Intn(len(s.waiting))
		out = append(out, s.waiting[idx])
		s.waiting = append(s.waiting[:idx], s.waiting[idx+1:]...)
	}
	s.waiting = out
	return out
}

// Fisher-Yates洗牌
func (s *Solution384) Shuffle_() []int {
	n := len(s.waiting)
	for i := range s.waiting {
		idx := i + rand.Intn(n-i)
		s.waiting[i], s.waiting[idx] = s.waiting[idx], s.waiting[i]
	}
	return s.waiting
}

func Test_Solution384(t *testing.T) {
	s := Constructor384([]int{1, 2, 3, 4})
	p := &s
	fmt.Println(p.Shuffle_())
	fmt.Println(p.Reset())
	fmt.Println(p.Shuffle_())
}

// tag-[数学]
// 蓄水池采样算法，保证留下来的概率都是K/N
func ReservoirSampling(nums []int, k int) []int {
	out := make([]int, k)
	for i := 0; i < k; i++ {
		out[i] = nums[i]
	}
	for i := k; i < len(nums); i++ {
		x := rand.Intn(i)
		if x < k {
			out[x] = nums[i]
		}
	}
	return out
}

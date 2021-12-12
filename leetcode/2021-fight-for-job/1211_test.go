package ojeveryday

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// tag-[深度优先搜索]
// leetcode1020: 飞地的数量
func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i >= m || i < 0 || j >= n || j < 0 {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	// 淹没边界
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	// 淹没边界
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

// tag-[深度优先搜索]
// leetcode1905: 统计子岛屿
func countSubIslands_(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i >= m || i < 0 || j >= n || j < 0 {
			return
		}
		if grid2[i][j] == 0 {
			return
		}
		grid2[i][j] = 0
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	// 先把不是子岛屿的排除
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				dfs(i, j)
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

// tag-[双指针]
// leetcode15: 三数之和
func threeSum_(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for i := 0; i < n; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			// 去重
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			v := nums[i] + nums[l] + nums[r]
			if v == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if v > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}

func Test_threeSum_(t *testing.T) {
	fmt.Println(threeSum_([]int{-1, 0, 1, 2, -1, -4}))
}

// tag-[双指针]
// leetcode 剑指 Offer 48. 最长不含重复字符的子字符串
func lengthOfLongestSubstring_(s string) int {
	m, n := make(map[uint8]int), len(s)
	l, r := 0, 0
	ans := 0
	for r < n {
		m[s[r]]++
		for m[s[r]] > 1 {
			m[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}

func Test_lengthOfLongestSubstring_(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring_("pwwkew"))
}

func minSubArrayLen__(target int, nums []int) int {
	n := len(nums)
	l, r := 0, 0
	sum := 0
	ans := math.MaxInt32
	for r < n {
		sum += nums[r]
		for sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func Test_minSubArrayLen__(t *testing.T) {
	fmt.Println(minSubArrayLen__(11, []int{1, 2, 3, 4, 5}))
}

func minWindow_(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	m := make(map[byte]int)
	for i := range t {
		m[t[i]]++
	}
	l, r := 0, 0
	cnt := math.MaxInt32
	var ans string
	for r < len(s) {
		if _, ok := m[s[r]]; ok {
			m[s[r]]--
		}
		for isOk(m) {
			if r-l+1 < cnt {
				cnt = r - l + 1
				ans = s[l : r+1]
			}
			if _, ok := m[s[l]]; ok {
				m[s[l]]++
			}
			l++
		}
		r++
	}
	return ans

}

func isOk(m map[byte]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

func Test_minWindow_(t *testing.T) {
	fmt.Println(minWindow_("ADOBECODEBANC", "ABC"))
}

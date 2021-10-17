// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// dp[x]表示以i结尾的等差数列的个数
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	dp := make([]int, len(nums))
	if nums[2]-nums[1] == nums[1]-nums[0] {
		dp[2] = 1
	}
	for i := 3; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	ans := 0
	for i := range dp {
		ans += dp[i]
	}
	return ans
}

func Test_numberOfArithmeticSlices(t *testing.T) {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 0, 5, 6, 7}))
}

// 双指针方法
func numberOfArithmeticSlices1(nums []int) int {
	n := len(nums)
	left, right := 0, 1
	ans := 0
	for left < n-1 {
		for right < n-1 && nums[right+1]-nums[right] == nums[right]-nums[right-1] {
			right++
		}
		if right-left >= 2 {
			ans += right - left - 1
			left++
		} else {
			left = right
			right++
		}
	}
	return ans
}

func Test_numberOfArithmeticSlices1(t *testing.T) {
	fmt.Println(numberOfArithmeticSlices1([]int{1, 2, 3, 0, 5, 6, 7}))
}

// 第二题
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if j < n-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
				if i < m-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
			}
		}
	}
	return dp
}

// 第三题
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	maxn := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != '0' {
				if i == 0 || j == 0 {
					dp[i][j] = int(matrix[i][j] - '0')
				} else {
					dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				}
				maxn = max(maxn, dp[i][j])
			}
		}
	}
	return maxn * maxn
}

func Test_maximalSquare(t *testing.T) {
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}

// 第四题
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	for i := 1; i < n; i++ {
		if s[i] > '0' && s[i] <= '9' {
			dp[i] += dp[i-1]
		}
		if s[i-1:i+1] >= "10" && s[i-1:i+1] <= "26" {
			if i >= 2 {
				dp[i] += dp[i-2]
			} else {
				dp[i] += 1
			}
		}
	}
	return dp[n-1]
}

// 空间压缩版，内存消耗略优于上面的版本
func numDecodings1(s string) int {
	n := len(s)
	dp := make([]int, 3)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i%3] = 0
		if s[i] > '0' && s[i] <= '9' {
			dp[i%3] += dp[(i-1)%3]
		}
		if s[i-1:i+1] >= "10" && s[i-1:i+1] <= "26" {
			if i >= 2 {
				dp[i%3] += dp[(i-2)%3]
			} else {
				dp[i%3] += 1
			}
		}
	}
	return dp[(n-1)%3]
}

func Test_numDecodings(t *testing.T) {
	fmt.Println(numDecodings1("226"))
}

// 第五题
func wordBreak1(s string, wordDict []string) bool {
	s = " " + s
	n := len(s)
	dp := make([]bool, n)
	dp[0] = true
	for i := 1; i < n; i++ {
		for j := 0; j < len(wordDict); j++ {
			if i >= len(wordDict[j]) && wordDict[j] == s[i+1-len(wordDict[j]):i+1] {
				dp[i] = dp[i] || dp[i-len(wordDict[j])]
			}
		}
	}
	return dp[n-1]
}

// dp[x]表示以i结尾的最长子序列
// 第六题
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	maxn := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxn = max(maxn, dp[i])
	}
	return maxn
}

func Test_lengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}

// 第七题
func isIsomorphic(s string, t string) bool {
	return isIsomorphicExec(s, t) && isIsomorphicExec(t, s)
}

func isIsomorphicExec(s string, t string) bool {
	m := make(map[byte]byte)
	for i := range s {
		if v, ok := m[s[i]]; !ok {
			m[s[i]] = t[i]
		} else {
			if v != t[i] {
				return false
			}
		}
	}
	return true
}

func Test_isIsomorphic(t *testing.T) {
	fmt.Println(isIsomorphic("egt", "add"))
}

func minimumDifference(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	left, right := 0, k-1
	minn := math.MaxInt32
	for right < n {
		minn = min(minn, nums[right]-nums[left])
		left++
		right++
	}
	return minn
}

func Test_minimumDifference(t *testing.T) {
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
}

func kthLargestNumber(nums []string, k int) string {
	sort.SliceStable(nums, func(i, j int) bool {
		if len(nums[i]) > len(nums[j]) {
			return true
		} else if len(nums[i]) < len(nums[j]) {
			return false
		} else {
			return nums[i] > nums[j]
		}
	})
	return nums[k-1]
}

func Test_kthLargestNumber(t *testing.T) {
	fmt.Println(kthLargestNumber([]string{"233", "97"}, 1))
}

type w struct {
	total int
	pos   int
}

func minSessions1(tasks []int, sessionTime int) int {
	n := len(tasks)
	dp := make([]w, sessionTime+1)
	flag := false

	for i := 0; i < n; i++ {
		if tasks[i] == 0 {
			continue
		}
		flag = true
		for j := sessionTime; j >= tasks[i]; j-- {
			if dp[j].total < dp[j-tasks[i]].total+tasks[i] {
				dp[j].total = dp[j-tasks[i]].total + tasks[i]
				dp[j].pos = i
			}
		}
	}
	if !flag {
		return 0
	}
	var ans []int
	for i := 0; i < sessionTime; {
		pos := dp[sessionTime-i].pos
		ans = append(ans, pos)
		i = pos
	}
	for i := range ans {
		tasks[i] = 0
	}
	return 1 + minSessions(tasks, sessionTime)
}

func Test_minSessions(t *testing.T) {
	fmt.Println(minSessions([]int{3, 1, 3, 1, 1}, 8))
}

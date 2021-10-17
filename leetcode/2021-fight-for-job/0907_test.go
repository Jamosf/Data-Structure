// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"testing"
)

func makeFancyString(s string) string {
	n := len(s)
	ans := make([]byte, 0)
	ans = append(ans, s[0])
	for i := 1; i < n; {
		if s[i] != s[i-1] {
			ans = append(ans, s[i])
			i++
			continue
		}
		cnt := 0
		for i < n && s[i] == s[i-1] {
			if cnt < 1 {
				ans = append(ans, s[i])
			}
			cnt++
			i++
		}
	}
	return string(ans)
}

func Test_makeFancyString(t *testing.T) {
	fmt.Println(makeFancyString("leeettttccccoooooddddeeee"))
}

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	m, n := len(board), len(board[0])
	var check func(dx, dy int) bool
	check = func(dx, dy int) bool {
		x, y := rMove+dx, cMove+dy
		step := 1
		for x >= 0 && x < m && y >= 0 && y < n {
			if step == 1 {
				if board[x][y] == color || board[x][y] == '.' {
					return false
				}
			} else {
				if board[x][y] == '.' {
					return false
				}
				if board[x][y] == color {
					return true
				}
			}
			step++
			x += dx
			y += dy
		}
		return false
	}
	direct := [8][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	for i := 0; i < 8; i++ {
		if check(direct[i][0], direct[i][1]) {
			return true
		}
	}
	return false
}

// 区间dp，dp[x][y]表示前i个元素被分成j段的最小值
// 状态转移方程：dp[x][y] = min(dp[x][y], dp[l-1][y-1] + weight[l][x])，其中l范围0...x-1
func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	weight := make([][]int, n)
	for i := range weight {
		weight[i] = make([]int, n)
	}
	// 1. 预处理权值
	for i := 0; i < n; i++ {
		sum := 0
		maxn := nums[i]
		for j := i; j < n; j++ {
			sum += nums[j]
			maxn = max(maxn, nums[j])
			weight[i][j] = maxn*(j-i+1) - sum
		}
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+2)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	// 2. dp计算
	for i := 0; i < n; i++ {
		for j := 1; j < k+2; j++ {
			for l := 0; l <= i; l++ {
				if l == 0 {
					dp[i][j] = min(dp[i][j], 0+weight[0][i])
				} else {
					dp[i][j] = min(dp[i][j], dp[l-1][j-1]+weight[l][i])
				}
			}
		}
	}
	return dp[n-1][k+1]
}

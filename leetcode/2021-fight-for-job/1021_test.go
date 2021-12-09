// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// tag-[数学]
// leetcode66:每日一题
func plusOne66(digits []int) []int {
	n := len(digits)
	carry := 1
	for i := n - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
		if i == 0 && carry > 0 {
			return append([]int{carry}, digits...)
		}
	}
	return digits
}

// tag-[字符串]
// leetcode6:字符串（模拟）
func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 {
		return s
	}
	numCols := n / (4 * (numRows - 1)) * 2 * (numRows - 1)
	left := n % (4 * (numRows - 1))
	if left >= 3*numRows-2 {
		numCols += numRows + (left - 3*numRows + 2)
	} else if left > 2*numRows-2 {
		numCols += numRows
	} else if left >= numRows {
		numCols += 1 + (left - numRows)
	} else {
		numCols += 1
	}
	matrix := make([][]byte, numRows)
	for i := range matrix {
		matrix[i] = make([]byte, numCols)
	}
	row, col := 0, 0
	up := false
	for i := range s {
		matrix[row][col] = s[i]
		if row == numRows-1 {
			up = true
		} else if row == 0 {
			up = false
		}
		if !up {
			row++
		} else {
			row--
			col++
		}
	}
	res := make([]byte, 0, n)
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if matrix[i][j] != 0 {
				res = append(res, matrix[i][j])
			}
		}
	}
	return string(res)
}

// tag-[字符串]
// leetcode6:优化解法，无需计算列的个数
func convert_(s string, numRows int) string {
	n := len(s)
	if numRows == 1 {
		return s
	}
	maxRow := min(numRows, n)
	res := make([][]byte, maxRow)
	row := 0
	up := false
	for i := range s {
		res[row] = append(res[row], s[i])
		if row == numRows-1 || row == 0 {
			up = !up
		}
		if !up {
			row++
		} else {
			row--
		}
	}
	ans := make([]string, 0, n)
	for i := 0; i < numRows; i++ {
		ans = append(ans, string(res[i]))
	}
	return strings.Join(ans, "")
}

func Test_convert_(t *testing.T) {
	fmt.Println(convert("A", 2))
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 5))
	fmt.Println(convert("PAYPALISHIRING", 6))
	fmt.Println(convert("PAYPALISHIRING", 7))
	fmt.Println(convert("PAYPALISHIRING", 8))
	fmt.Println(convert("PAYPALISHIRING", 9))
}

// tag-[字符串]
// leetcode481:神奇字符串（模拟）
func magicalString_(n int) int {
	if n < 3 {
		return 1
	}
	t := make([]byte, 3, n)
	t[0] = '1'
	t[1] = '2'
	t[2] = '2'
	cnt := 1
	fast, slow := 2, 1
	for len(t) < n {
		slow++
		if t[slow] == '2' {
			if t[fast] == '2' {
				t = append(t, []byte{'1', '1'}...)
				if len(t) > n {
					cnt += 1
				} else {
					cnt += 2
				}
			}
			if t[fast] == '1' {
				t = append(t, []byte{'2', '2'}...)
			}
			fast += 2
		} else {
			if t[fast] == '2' {
				t = append(t, '1')
				cnt += 1
			}
			if t[fast] == '1' {
				t = append(t, '2')
			}
			fast += 1
		}
	}
	return cnt
}

func Test_magicalString_(t *testing.T) {
	fmt.Println(magicalString_(100))
}

// tag-[字符串]
// leetcode71: 简化路径
func simplifyPath(path string) string {
	ss := strings.Split(path, "/")
	newPath := make([]string, 0)
	for _, s := range ss {
		if s == "." || s == "" {
			continue
		} else if s == ".." {
			if len(newPath) > 0 {
				newPath = newPath[:len(newPath)-1]
			}
		} else {
			newPath = append(newPath, s)
		}
	}
	return "/" + strings.Join(newPath, "/")
}

func Test_simplifyPath(t *testing.T) {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/../"))
}

// tag-[回溯]
// leetcode93:复原ip地址
func restoreIpAddresses(s string) []string {
	n := len(s)
	if n > 12 {
		return nil
	}
	ans := make([]string, 0)
	tmp := make([]string, 0)
	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == n {
			if len(tmp) == 4 {
				fmt.Println(tmp)
				ans = append(ans, strings.Join(tmp, "."))
			}
			return
		}
		for i := idx + 1; i <= n; i++ {
			if isValidIp(s[idx:i]) {
				tmp = append(tmp, s[idx:i])
				backtrace(i)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrace(0)
	return ans
}

func isValidIp(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return v >= 0 && v <= 255
}

func Test_isValidIp(t *testing.T) {
	fmt.Println(isValidIp("265"))
}

func Test_restoreIpAddresses(t *testing.T) {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("101023"))
}

// tag-[字符串/回溯]
// leetcode97:交错字符串(记忆化搜索)
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n3 != n1+n2 {
		return false
	}
	dp := make([][]bool, n1) // 记忆化，存储中间过程数据
	for i := range dp {
		dp[i] = make([]bool, n2)
	}
	var backtrace func(idx1, idx2, idx3 int) bool
	backtrace = func(idx1, idx2, idx3 int) bool {
		if idx3 == n3 {
			return true
		}
		if idx1 < n1 && idx2 < n2 && dp[idx1][idx2] {
			return false
		}
		if idx1 < n1 && s3[idx3] == s1[idx1] && backtrace(idx1+1, idx2, idx3+1) {
			return true
		}
		if idx2 < n2 && s3[idx3] == s2[idx2] && backtrace(idx1, idx2+1, idx3+1) {
			return true
		}
		if idx1 < n1 && idx2 < n2 {
			dp[idx1][idx2] = true
		}
		return false
	}
	return backtrace(0, 0, 0)
}

// tag-[动态规划]
// leetcode97:交错字符串(记忆化搜索)
func isInterleave_(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n3 != n1+n2 {
		return false
	}
	dp := make([][]bool, n1+1) // 记忆化，存储中间过程数据
	for i := range dp {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s1[i-1] == s3[p] // i-1对应数组的索引是第i个，和dp刚好差一
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[n1][n2]
}

func Test_isInterleave(t *testing.T) {
	fmt.Println(isInterleave_("aabcc", "dbbca", "aadbbcbcac"))
}

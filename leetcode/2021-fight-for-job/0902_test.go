// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// dp[i]的含义是：i表示的二进制对应的索引全部选中的情况下，最小的组合数
func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	m := 1 << n
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 20 // 以普遍理论而论，状态压缩的数据范围会小于20
	}
	// 1. 计算哪些子集可以在sessionTime内完成，标记为1
	for i := 0; i < m; i++ {
		state, idx := i, 0
		spend := 0
		for state > 0 {
			if state&1 == 1 {
				spend += tasks[idx]
			}
			state >>= 1
			idx++
		}
		if spend <= sessionTime {
			dp[i] = 1
		}
	}
	// 2. 从第一步计算出子集推到全集的最小值，枚举二进制子集并完成状态转移。
	// 一个111的子集可以由110 + 1构成。即i = 111， j = 110， i^j = 1。i^j，有等于i-j的语义。
	// 另外还有一个更新j的技巧：j = (j-1) &i， j: 110->101->100->011->010->001。这样会出现一个重复计算的问题，因为j减少的时候，会等于i^j。
	// 原则是，j与i^j一定是i的子集。
	for i := 1; i < m; i++ {
		if dp[i] == 1 {
			continue
		}
		for j := i; j > 0; j = (j - 1) & i {
			dp[i] = min(dp[i], dp[j]+dp[i^j])
		}
	}
	return dp[m-1]
}

// dfs解这个题
func minSessions2(tasks []int, sessionTime int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})
	ans := 20
	n := len(tasks)
	times := make([]int, n)
	var dfs func(u, k int)
	dfs = func(u, k int) {
		if k >= ans { // 再往后搜索已经不可能出现更小的解
			return
		}
		if u == n { // 搜索到了最后一个数
			ans = k
			return
		}
		// 1. 先尝试往老的时间段里面塞塞看
		for i := 0; i < k; i++ {
			if times[i]+tasks[u] <= sessionTime {
				times[i] += tasks[u]
				dfs(u+1, k)
				times[i] -= tasks[u]
			}
		}
		// 2. 使用新的时间段
		times[k] = tasks[u]
		dfs(u+1, k+1)
		times[k] = 0
	}
	dfs(0, 0)
	return ans
}

func Test_minSessions2(t *testing.T) {
	fmt.Println(minSessions2([]int{1, 1, 1, 3, 3, 1}, 8))
}

//
func numberOfUniqueGoodSubsequences(binary string) int {
	n := len(binary)
	dp0, dp1 := 0, 0
	has0, mod := 0, int(1e9+7)
	for i := n - 1; i >= 0; i-- {
		if binary[i] == '0' {
			has0 = 1
			dp0 = (dp0 + dp1 + 1) % mod
		} else {
			dp1 = (dp0 + dp1 + 1) % mod
		}
	}
	return (dp1 + has0) % mod
}

// 分割字符串
func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	t1 := make([]int, len(s1))
	s2 := strings.Split(version2, ".")
	t2 := make([]int, len(s2))
	for i := range s1 {
		v := strings.TrimLeft(s1[i], "0")
		if v != "" {
			t1[i], _ = strconv.Atoi(v)
		} else {
			t1[i] = 0
		}
	}
	for i := range s2 {
		v := strings.TrimLeft(s2[i], "0")
		if v != "" {
			t2[i], _ = strconv.Atoi(v)
		} else {
			t2[i] = 0
		}
	}
	n := min(len(t1), len(t2))
	for i := 0; i < n; i++ {
		if t1[i] > t2[i] {
			return 1
		} else if t1[i] < t2[i] {
			return -1
		}
	}
	if len(t1) > len(t2) {
		for i := len(t2); i < len(t1); i++ {
			if t1[i] > 0 {
				return 1
			}
		}
	}
	if len(t1) < len(t2) {
		for i := len(t1); i < len(t2); i++ {
			if t2[i] > 0 {
				return -1
			}
		}
	}
	return 0
}

func Test_compareVersion(t *testing.T) {
	fmt.Println(compareVersion1("1.2", "1.10"))
}

// 双指针解法
func compareVersion1(v1 string, v2 string) int {
	n1, n2 := len(v1), len(v2)
	pre1, pre2 := 0, 0
	pos1, pos2 := 0, 0
	for pos1 < n1 && pos2 < n2 {
		for pos1 < n1 && v1[pos1] != '.' {
			pos1++
		}
		for pos2 < n2 && v2[pos2] != '.' {
			pos2++
		}
		for pre1 < pos1 && v1[pre1] == '0' {
			pre1++
		}
		for pre2 < pos2 && v2[pre2] == '0' {
			pre2++
		}
		if idx := compare(v1[pre1:pos1], v2[pre2:pos2]); idx != 0 {
			return idx
		}
		pos1++
		pos2++
		pre1, pre2 = pos1, pos2
	}
	for pos1 < n1 {
		if v1[pos1] > '0' && v1[pos1] != '.' {
			return 1
		}
		pos1++
	}
	for pos2 < n2 {
		if v2[pos2] > '0' && v2[pos2] != '.' {
			return -1
		}
		pos2++
	}
	return 0
}

func compare(s, t string) int {
	if len(s) > len(t) {
		return 1
	}
	if len(s) < len(t) {
		return -1
	}
	if s > t {
		return 1
	}
	if s < t {
		return -1
	}
	return 0
}

// leetcode 100%的典范代码
func compareVersion3(version1 string, version2 string) int {
	v1 := NewVersionIterator(version1)
	v2 := NewVersionIterator(version2)
	for {
		r1, ok1 := v1.NextRevision()
		r2, ok2 := v2.NextRevision()
		if r1 > r2 {
			return 1
		}

		if r1 < r2 {
			return -1
		}

		if !ok1 && !ok2 {
			break
		}
	}

	return 0
}

type VersionIterator struct {
	version string
	index   int
}

func NewVersionIterator(version string) *VersionIterator {
	return &VersionIterator{
		version: version,
	}
}

func (v *VersionIterator) NextRevision() (int, bool) {
	if v.index == len(v.version) {
		return 0, false
	}

	revision := 0
	for {
		ch := v.version[v.index]
		v.index++
		if ch == '.' {
			break
		}

		revision = revision*10 + int(ch-'0')
		if v.index == len(v.version) {
			break
		}
	}

	return revision, true
}

// 输入：s = "aacecaaa"
// 输出："aaacecaaa"
// 解法：将字符串翻转，然后去掉中间重叠的部分即可
func shortestPalindrome(s string) string {
	n := len(s)
	s1 := []byte(s)
	for i := 0; i < n/2; i++ {
		s1[i], s1[n-i-1] = s1[n-i-1], s1[i]
	}
	ss1 := string(s1)
	// 技巧：逆序遍历，可以最大化的删除重复元素，保证最终得到的字符串最小
	for i := n; i >= 0; i-- {
		if s[:i] == ss1[n-i:] {
			return ss1[:n-i] + s
		}
	}
	return ""
}

func Test_shortestPalindrome(t *testing.T) {
	fmt.Println(shortestPalindrome("aacecaaa"))
}

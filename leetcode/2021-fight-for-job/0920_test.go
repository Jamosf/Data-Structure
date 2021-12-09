// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

// tag-[字符串]
// leetcode2011：执行操作后的变量值
func finalValueAfterOperations(operations []string) int {
	ans := 0
	for i := range operations {
		if operations[i] == "X++" || operations[i] == "++X" {
			ans++
		} else {
			ans--
		}
	}
	return ans
}

// tag-[单调栈]
// leetcode2012: 数组美丽值求和
// 双向单调栈
func sumOfBeauties(nums []int) int {
	n := len(nums)
	// 1. 正向的递减单调栈
	s := make([]int, 0)
	t := make([]int, n)
	for i := range t {
		t[i] = -1
	}
	for i := range nums {
		for len(s) != 0 && nums[s[len(s)-1]] >= nums[i] {
			t[s[len(s)-1]] = i
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	// 2. 反向的单调递增栈
	k := make([]int, n)
	for i := range k {
		k[i] = -1
	}
	s = s[:0]
	for i := n - 1; i >= 0; i-- {
		for len(s) != 0 && nums[s[len(s)-1]] <= nums[i] {
			k[s[len(s)-1]] = i
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	ans := 0
	for i := range t {
		if i > 0 && i < n-1 {
			if t[i] == -1 && k[i] == -1 {
				ans += 2
			} else {
				if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
					ans += 1
				}
			}
		}
	}
	return ans
}

// tag-[前缀和]
// leetcode2012: 数组美丽值求和
// 前缀最大值和后缀最小值
func sumOfBeauties_(nums []int) int {
	n := len(nums)
	f1 := make([]int, n)
	f1[0] = nums[0]
	for i := 1; i < n; i++ {
		f1[i] = max(f1[i-1], nums[i])
	}
	f2 := make([]int, n)
	f2[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		f2[i] = min(f2[i+1], nums[i])
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		if nums[i] > f1[i-1] && nums[i] < f2[i+1] {
			ans += 2
			continue
		}
		if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			ans += 1
		}
	}
	return ans
}

func Test_sumOfBeauties(t *testing.T) {
	fmt.Println(sumOfBeauties([]int{5, 5, 10, 4, 6}))
	fmt.Println(sumOfBeauties_([]int{5, 5, 10, 4, 6}))
}

// tag-[矩阵]
// leetcode2013: 检测正方形
type DetectSquares struct {
	p map[int]map[int]int
}

func ConstructorDetectSquares() DetectSquares {
	return DetectSquares{make(map[int]map[int]int)}
}

func (d *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if _, ok := d.p[x]; !ok {
		d.p[x] = make(map[int]int)
	}
	d.p[x][y]++
}

func (d *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	if _, ok := d.p[x]; !ok {
		return 0
	}
	ans := 0
	for y1, c := range d.p[x] {
		if y != y1 {
			ans += c * d.p[x+minusAbs(y, y1)][y] * d.p[x+minusAbs(y, y1)][y1]
			ans += c * d.p[x-minusAbs(y, y1)][y] * d.p[x-minusAbs(y, y1)][y1]
		}
	}
	return ans
}

func Test_DetectSquares(t *testing.T) {
	d := ConstructorDetectSquares()
	d.Add([]int{3, 10})
	d.Add([]int{3, 10})
	d.Add([]int{11, 2})
	d.Add([]int{3, 2})
	d.Add([]int{3, 2})
	d.Add([]int{3, 2})
	fmt.Println(d.Count([]int{11, 10}))
	fmt.Println(d.Count([]int{14, 8}))
	d.Add([]int{11, 2})
	fmt.Println(d.Count([]int{11, 10}))
}

// tag-[字符串]
func longestSubsequenceRepeatedK(s string, k int) (ans string) {
	n := len(s)
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}
	nxt := make([][26]int, n)
	cnt := [26]int{}
	for i := n - 1; i >= 0; i-- {
		nxt[i] = pos
		pos[s[i]-'a'] = i
		cnt[s[i]-'a']++
	}

	// 计算所有可能出现在答案中的字符，包括重复的
	// 倒着统计，这样下面计算排列时的第一个合法方案就是答案，从而提前退出
	a := []byte{}
	for i := 25; i >= 0; i-- {
		for c := cnt[i]; c >= k; c -= k {
			a = append(a, 'a'+byte(i))
		}
	}

	for m := len(a); m > 0 && ans == ""; m-- { // 从大到小枚举答案长度 m
		permutations(len(a), m, func(ids []int) bool { // 枚举长度为 m 的所有排列
			t := make([]byte, m)
			for i, id := range ids {
				t[i] = a[id]
			}
			i, j := 0, 0
			if t[0] == s[0] {
				j = 1
			}
			for {
				i = nxt[i][t[j%m]-'a']
				if i == n {
					break
				}
				j++
			}
			if j >= k*m {
				ans = string(t)
				return true // 提前退出
			}
			return false
		})
	}
	return
}

// 模板：生成 n 选 r 的排列
func permutations(n, r int, do func(ids []int) bool) {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	if do(ids[:r]) {
		return
	}
	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}
	for {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				tmp := ids[i]
				copy(ids[i:], ids[i+1:])
				ids[n-1] = tmp
				cycles[i] = n - i
			} else {
				j := cycles[i]
				ids[i], ids[n-j] = ids[n-j], ids[i]
				if do(ids[:r]) {
					return
				}
				break
			}
		}
		if i == -1 {
			return
		}
	}
}

// tag-[哈希表]
// leetcode2006：差的绝对值为k的数对数目
func countKDifference(nums []int, k int) int {
	m := make(map[int]int)
	ans := 0
	for i := range nums {
		ans += m[nums[i]-k]
		ans += m[nums[i]+k]
		m[nums[i]]++
	}
	return ans
}

// tag-[数组]
// leetcode2007: 从双倍数组中还原数组
func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	cnt := make(map[int]int)
	ans := make([]int, 0)
	for _, v := range changed {
		if cnt[v] == 0 {
			cnt[v*2]++
			ans = append(ans, v)
		} else {
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
		}
	}
	if len(cnt) == 0 {
		return ans
	}
	return nil
}

func Test_findOriginalArray(t *testing.T) {
	fmt.Println(findOriginalArray([]int{1, 3, 4, 2, 6, 8}))
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// tag-[动态规划]
// leetcode2008: 出租车的最大盈利
func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool {
		a, b := rides[i], rides[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	m := len(rides)
	dp := make([]int64, m)
	dp[0] = int64(rides[0][1] - rides[0][0] + rides[0][2])
	maxn := dp[0]
	for i := 1; i < m; i++ {
		v := int64(rides[i][1] - rides[i][0] + rides[i][2])
		flag := false
		for j := i - 1; j >= 0; j-- {
			if rides[i][0] >= rides[j][1] {
				flag = true
				dp[i] = maxInt64(dp[i], dp[j]+v)
			}
		}
		if !flag {
			dp[i] = v
		}
		maxn = maxInt64(maxn, dp[i])
	}
	return maxn
}

// leetcode2008: 出租车的最大盈利
func maxTaxiEarnings_(n int, rides [][]int) int64 {
	g := make([][][2]int, n+1)
	for _, v := range rides {
		start, end, trips := v[0], v[1], v[2]
		g[end] = append(g[end], [2]int{start, trips})
	}
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[i-1]
		for _, e := range g[i] {
			f[i] = max(f[i], f[e[0]]+i-e[0]+e[1])
		}
	}
	return int64(f[n])
}

func Test_maxTaxiEarnings(t *testing.T) {
	fmt.Println(maxTaxiEarnings(10, [][]int{{9, 10, 2}, {4, 5, 6}, {6, 8, 1}, {1, 5, 5}, {4, 9, 5}, {1, 6, 5}, {4, 8, 3}, {4, 7, 10}, {1, 9, 8}, {2, 3, 5}}))
	fmt.Println(maxTaxiEarnings_(10, [][]int{{9, 10, 2}, {4, 5, 6}, {6, 8, 1}, {1, 5, 5}, {4, 9, 5}, {1, 6, 5}, {4, 8, 3}, {4, 7, 10}, {1, 9, 8}, {2, 3, 5}}))
}

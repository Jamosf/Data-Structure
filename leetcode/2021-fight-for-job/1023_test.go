// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

// leetcode638:每日一题,dfs记忆化搜索（非回溯）
func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	filterSpecial := make([][]int, 0)
	totalCnt, totalPrice := 0, 0
	// 先过滤出优先的大礼包
	for i := range special {
		totalCnt, totalPrice = 0, 0
		for j, v := range special[i][:n] {
			totalCnt++
			totalPrice += v * price[j]
		}
		if totalCnt > 0 && totalPrice > special[i][n] {
			filterSpecial = append(filterSpecial, special[i])
		}
	}
	dp := make(map[string]int)
	var dfs func(needs []byte) int
	dfs = func(needs []byte) int {
		if v, ok := dp[string(needs)]; ok {
			return v
		}
		ans := 0
		for i := range needs {
			ans += int(needs[i]) * price[i]
		}
		nextNeeds := make([]byte, len(needs))
	outer:
		for i := range filterSpecial {
			for j, v := range filterSpecial[i][:n] {
				if v > int(needs[j]) {
					continue outer
				}
				nextNeeds[j] = byte(int(needs[j]) - v)
			}
			ans = min(ans, dfs(nextNeeds)+filterSpecial[i][n])
		}
		dp[string(needs)] = ans
		return ans
	}
	needs_ := make([]byte, len(needs))
	for i := range needs {
		needs_[i] = byte(needs[i])
	}
	return dfs(needs_)
}

// leetcode638:回溯法思想
func shoppingOffers_(price []int, special [][]int, needs []int) int {
	n := len(price)
	m := len(special)
	sum := 0
	minn := math.MaxInt32
	smaller := func(a, b []int) bool {
		for i := range a {
			if a[i] > b[i] {
				return false
			}
		}
		return true
	}
	var backtracking func(idx int, left []int)
	backtracking = func(idx int, left []int) {
		if idx == m {
			extra := 0
			for i := range left {
				extra += left[i] * price[i]
			}
			minn = min(minn, sum+extra)
			return
		}
		for i := idx; i < m; i++ { // 先尝试在大礼包中选择，同步更新left和最小值
			if smaller(special[i][:n], left) {
				sum += special[i][n]
				for j := range left {
					left[j] -= special[i][j]
				}
				backtracking(i, left) // 有条件的递归，i有可能无法达到m
				sum -= special[i][n]
				for j := range left {
					left[j] += special[i][j]
				}
			}
		}
		backtracking(m, left) // i不一定能达到m，因此需要最后用m来收尾处理。大礼包选择完后，还有多余的，则按照价格购买。
	}
	backtracking(0, needs)
	return minn
}

func Test_shoppingOffers(t *testing.T) {
	// fmt.Println(shoppingOffers_([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}))
	fmt.Println(shoppingOffers_([]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}))
}

// leetcode周赛第一题
func countValidWords(sentence string) int {
	ss := strings.Split(sentence, " ")
	isChar := func(v byte) bool {
		return v >= 'a' && v <= 'z'
	}
	isPunctuation := func(v byte) bool {
		return v == '!' || v == '.' || v == ','
	}
	isValid := func(s string) bool {
		cnt1 := 0
		cnt2 := 0
		for j := range s {
			v := s[j]
			if !isChar(v) && v != '-' && !isPunctuation(v) {
				return false
			}
			if v == '-' {
				if cnt1 > 0 {
					return false
				}
				cnt1++
				if j > 0 && j < len(s)-1 && isChar(s[j-1]) && isChar(s[j+1]) {
					continue
				}
				return false
			}
			if isPunctuation(v) {
				if cnt2 > 0 {
					return false
				}
				cnt2++
				if j != len(s)-1 {
					return false
				}
			}
		}
		return true
	}
	cnt := 0
	for i := range ss {
		if ss[i] == "" || ss[i] == " " {
			continue
		}
		if isValid(ss[i]) {
			cnt++
		}
	}
	return cnt
}

func Test_countValidWords(t *testing.T) {
	fmt.Println(countValidWords("alice and  bob are playing stone-game10"))
	fmt.Println(countValidWords("he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."))
	fmt.Println(countValidWords("!this  1-s b8d!"))
	fmt.Println(countValidWords("cat and  dog"))
	fmt.Println(countValidWords("!this  a-s- bad!"))
}

// leetcode周赛第二题
func nextBeautifulNumber(n int) int {
	v := n
	cnt := 0
	for v != 0 {
		cnt++
		v = v / 10
	}
	isBeautiful := func(t int) bool {
		if t == 0 {
			return false
		}
		numCnt := [9]int{}
		for t != 0 {
			numCnt[t%10]++
			t = t / 10
		}
		for i := range numCnt {
			if numCnt[i] != 0 && numCnt[i] != i {
				return false
			}
		}
		return true
	}
	convert := func(b []int) int {
		res := 0
		multi := 1
		for i := len(b) - 1; i >= 0; i-- {
			res += b[i] * multi
			multi *= 10
		}
		return res
	}
	ans := 0
	flag := false
	maxn := 6
	tmp := make([]int, 0)
	var backtrace func(idx int, delta int)
	backtrace = func(idx int, delta int) {
		if flag {
			return
		}
		if idx >= cnt {
			val := convert(tmp)
			fmt.Println(tmp)
			if val > n && isBeautiful(val) {
				flag = true
				ans = val
				return
			}
		}
		if idx > cnt+delta {
			return
		}
		for i := 1; i <= maxn; i++ {
			tmp = append(tmp, i)
			backtrace(idx+1, delta)
			tmp = tmp[:len(tmp)-1]
		}
	}
	for i := 1; i <= 6; i++ {
		backtrace(1, i-1)
		if flag {
			break
		}
	}

	return ans
}

// 暴力即可通过版本
func nextBeautifulNumber_(n int) int {
	isBeautiful := func(t int) bool {
		if t == 0 {
			return false
		}
		numCnt := [10]int{}
		for t != 0 {
			numCnt[t%10]++
			t = t / 10
		}
		for i := range numCnt {
			if numCnt[i] != 0 && numCnt[i] != i {
				return false
			}
		}
		return true
	}
	for {
		n++
		if isBeautiful(n) {
			return n
		}
	}
}

func Test_nextBeautifulNumber(t *testing.T) {
	// fmt.Println(nextBeautifulNumber(1))
	fmt.Println(nextBeautifulNumber_(3000))
}

// leetcode周赛第三题
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	// 建树
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		v := parents[i]
		g[v] = append(g[v], i)
	}
	ans := 0
	maxScore := 0
	var dfs func(int) int
	dfs = func(v int) int {
		size, score := 1, 1
		for _, w := range g[v] {
			sz := dfs(w)
			size += sz
			score *= sz // 左右相乘
		}
		if v > 0 {
			score *= n - size
		}
		if score > maxScore {
			maxScore, ans = score, 1
		} else if score == maxScore {
			ans++
		}
		return size
	}
	dfs(0)
	return ans
}

// 周赛第三题参考
func countHighestScoreNodes_(parents []int) (ans int) {
	n := len(parents)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		v := parents[w]
		g[v] = append(g[v], w) // 建树
	}

	maxScore := 0
	var dfs func(int) int
	dfs = func(v int) int {
		size, score := 1, 1
		for _, w := range g[v] {
			sz := dfs(w)
			size += sz
			score *= sz // 由于是二叉树所以 score 最大约为 (1e5/3)^3，在 64 位整数范围内
		}
		if v > 0 {
			score *= n - size
		}
		if score > maxScore {
			maxScore, ans = score, 1
		} else if score == maxScore {
			ans++
		}
		return size
	}
	dfs(0)
	return
}

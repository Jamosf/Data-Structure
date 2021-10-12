// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

func minimumMoves(s string) int {
	n := len(s)
	idx := 0
	cnt := 0
	for idx < n {
		if s[idx] == 'X' {
			cnt++
			idx += 3
		} else {
			idx++
		}
	}
	return cnt
}

// dfs求解超时
func missingRolls(rolls []int, mean int, n int) []int {
	t := mean * (n + len(rolls))
	sum := 0
	for i := range rolls {
		sum += rolls[i]
	}
	diff := t - sum
	total := 0
	flag := false
	var ans []int
	var res []int
	var dfs func(depth int)
	dfs = func(depth int) {
		if flag {
			return
		}
		if depth == n {
			if total == diff {
				res = make([]int, len(ans))
				copy(res, ans)
				flag = true
			}
			return
		}
		for i := 1; i <= 6; i++ {
			if flag {
				break
			}
			total += i
			ans = append(ans, i)
			dfs(depth + 1)
			ans = ans[:len(ans)-1]
			total -= i
		}
	}
	dfs(0)
	return res
}

func Test_missingRolls(t *testing.T) {
	fmt.Println(missingRolls1([]int{6, 3, 4, 3, 5, 3}, 1, 6))
}

// 求解
func missingRolls1(rolls []int, mean int, n int) []int {
	t := mean * (n + len(rolls))
	sum := 0
	for i := range rolls {
		sum += rolls[i]
	}
	diff := t - sum
	if diff <= 0 {
		return nil
	}
	v := diff / n
	p := diff % n
	if v == 0 || v > 6 || (v == 6 && p != 0) {
		return nil
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = v
		if i < p {
			ans[i]++
		}
	}
	return ans
}

// 第三题
func stoneGameIX(stones []int) bool {
	c := [3]int{}
	for _, v := range stones {
		c[v%3]++
	}
	return checkW(c) || checkW([3]int{c[0], c[2], c[1]})
}

func checkW(c [3]int) bool {
	if c[1] == 0 { // 如果余数为1的个数等于0，则直接看余数为2的个数
		return false
	}
	c[1]--
	turn := 1 + min(c[1], c[2])*2 + c[0]
	if c[1] > c[2] { // 如果以1开头，序列末尾可以再加个1，和也不能被3整除
		turn++
		c[1]--
	}
	return turn%2 == 1 && c[1] != c[2] // 回合为奇数，且还有石子剩余，轮到bob出，则alice胜出
}

// 第四题
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	ans := make([]int, m+n)
	var t uint8
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			t = (num2[i] - '0') * (num1[j] - '0')
			ans[i+j+1] += int(t % 10)
			c := 0
			if ans[i+j+1] >= 10 {
				ans[i+j+1] %= 10
				c = 1
			}
			carry := int(t/10) + c
			idx := i + j
			for {
				if ans[idx]+carry >= 10 {
					v := ans[idx] + carry
					ans[idx] = v % 10
					carry = v / 10
					idx--
				} else {
					ans[idx] += carry
					break
				}
			}
		}
	}
	b := make([]byte, m+n)
	for i := range ans {
		b[i] = byte(ans[i] + '0')
	}
	i := 0
	for b[i] == '0' {
		i++
	}
	return string(b[i:])
}

func Test_multiply(t *testing.T) {
	fmt.Println(multiply("1234", "4567"))
}

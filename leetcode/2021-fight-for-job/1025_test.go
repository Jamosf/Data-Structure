// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

func intToBytes(n int64) []byte {
	ans := make([]byte, 0)
	for n != 0 {
		ans = append(ans, byte(n%10)+'0')
		n /= 10
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func Test_intToBytes(t *testing.T) {
	fmt.Println(string(intToBytes(12345)))
}

// leetcode166
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	num, den := int64(numerator), int64(denominator)
	isNegtive := false
	if num < 0 {
		isNegtive = !isNegtive
		num = -num
	}
	if den < 0 {
		isNegtive = !isNegtive
		den = -den
	}
	ans := make([]byte, 0)
	m := make(map[int64][2]int64)
	first := true
	cycle := [2]int{}
	for num != 0 {
		quotient := num / den
		reminder := num % den
		if v, ok := m[reminder]; ok && reminder != 0 && v[1] == quotient {
			cycle[0] = int(v[0])
			cycle[1] = len(ans)
			break
		} else if !first {
			m[reminder] = [2]int64{int64(len(ans)), quotient}
		}
		if quotient < 1 {
			if first {
				first = !first
				if len(ans) == 0 {
					ans = append(ans, '0')
				}
				ans = append(ans, '.')
			} else {
				ans = append(ans, '0')
			}
			num *= 10
		} else {
			ans = append(ans, intToBytes(quotient)...)
			if first && reminder != 0 {
				first = !first
				ans = append(ans, '.')
			}
			num = reminder * 10
		}
	}
	if cycle[0] != 0 || cycle[1] != 0 {
		ans = append(ans[:cycle[0]], append([]byte{'('}, ans[cycle[0]:]...)...)
		ans = append(ans, ')')
	}
	if isNegtive {
		ans = append([]byte{'-'}, ans...)
	}
	return string(ans)
}

func Test_fractionToDecimal(t *testing.T) {
	fmt.Println(fractionToDecimal(20, 3))
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(1, 5))
	fmt.Println(fractionToDecimal(45, 698))
	fmt.Println(fractionToDecimal(100, 9))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(0, 3))
	fmt.Println(fractionToDecimal(500, 10))
	fmt.Println(fractionToDecimal(140898435, 17))
	fmt.Println(fractionToDecimal(-50, 8))
}

// leetcode263
func isUgly(n int) bool {
	factors := []int{2, 3, 5}
	if n <= 0 {
		return false
	}
	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}

// leetcode264: 多路归并
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	p2, p3, p5 := 1, 1, 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if x2 == dp[i] {
			p2++
		}
		if x3 == dp[i] {
			p3++
		}
		if x5 == dp[i] {
			p5++
		}
	}
	return dp[n]
}

func Test_nthUglyNumber(t *testing.T) {
	fmt.Println(nthUglyNumber(100))
}

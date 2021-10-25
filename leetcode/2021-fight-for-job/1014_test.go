package ojeveryday

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"testing"
)

func fourSum1(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				v := nums[i] + nums[j] + nums[l] + nums[r]
				if v > target {
					r--
				} else if v < target {
					l++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}
		}
	}
	return ans
}

func Test_fourSum(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

// leetcode8
func myAtoi(s string) int {
	ss := strings.TrimLeft(s, " ")
	ans := make([]byte, 0)
	for i := range ss {
		v := ss[i]
		if v >= '0' && v <= '9' {
			ans = append(ans, v)
		}
		if v == '-' || v == '+' {
			if len(ans) == 0 {
				ans = append(ans, v)
			} else {
				break
			}
		}
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || v == '.' || v == ' ' {
			break
		}
	}
	var factor int
	if len(ans) > 0 {
		if ans[0] == '-' {
			factor = -1
			ans = ans[1:]
		} else if ans[0] == '+' {
			factor = 1
			ans = ans[1:]
		} else {
			factor = 1
		}
	}
	res := 0
	t := 1
	for i := range ans {
		res += int(ans[len(ans)-i-1]-'0') * t
		if factor == 1 && (res > math.MaxInt32 || t > math.MaxInt32) {
			return math.MaxInt32
		}
		if factor == -1 && (res < math.MinInt32 || t < math.MinInt32) {
			return math.MinInt32
		}
		t *= 10
	}
	r := factor * res
	if r > math.MaxInt32 {
		return math.MaxInt32
	}
	if r < math.MinInt32 {
		return math.MinInt32
	}
	return r
}

func Test_myAtoi(t *testing.T) {
	fmt.Println(myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
}

// leetcode8 优化解法
func myAtoi1(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	//丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	//标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')  //字节 byte '0' == 48
		if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}

func Test_forsum(t *testing.T) {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

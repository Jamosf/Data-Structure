// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"math"
	"sort"
)

// tag-[数学]
// leetcode29
// 快速乘
// x 和 y 是负数，z 是正数
// 判断 z * y >= x 是否成立
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 { // 不能使用除法
		if z&1 > 0 {
			// 需要保证 result + add >= x
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			// 需要保证 add + add >= x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

// tag-[数学]
func divide(dividend, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 // 注意溢出，并且不能使用除法
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 { // 注意溢出
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans
}

// leetcode319 脑筋急转弯
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

// tag-[数组]
// leetcode1033
func numMovesStones(a int, b int, c int) []int {
	v := []int{a, b, c}
	sort.Ints(v)
	minn := v[0]
	maxn := v[2]
	minv := 0
	if maxn-v[1] > 1 {
		minv++
	}
	if v[1]-minn > 1 {
		minv++
	}
	if maxn-v[1] == 2 || v[1]-minn == 2 {
		minv = 1
	}
	return []int{minv, v[1] - minn + maxn - v[1] - 2}
}

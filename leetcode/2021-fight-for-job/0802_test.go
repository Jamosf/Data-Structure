// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// 第二题
// leetcode 剑指offer56-I：数组中数字出现的次数
// 数组中两个出现一次的数，分开异或
func singleNumbers(nums []int) []int {
	m, n := 0, 1
	x, y := 0, 0
	for _, v := range nums {
		m ^= v
	}
	for m&n == 0 {
		n <<= 1
	}
	for _, v := range nums {
		if v&n == 0 {
			x ^= v
		} else {
			y ^= v
		}
	}
	return []int{x, y}
}

func Test_singleNumbers(t *testing.T) {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}

// 第三题
// leetcode剑指offer56-II: 数组中数字出现的次数II
func singleNumberII(nums []int) int {
	m := make([]int, 32)
	for _, v := range nums {
		for i := 0; i < 32; i++ {
			if v&(1<<i) != 0 {
				m[i]++
			}
		}
	}
	ans := 0
	for i, v := range m {
		if v%3 != 0 {
			ans += 1 << i
		}
	}
	return ans
}

// 第四题
// leetcode剑指offer 66： 构建乘积数组
func constructArr(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	b := make([]int, len(a))
	b[0] = 1
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	tmp := 1
	for j := len(a) - 2; j >= 0; j-- {
		b[j] *= tmp * a[j+1]
		tmp *= a[j+1]
	}
	return b
}

func Test_constructArr(t *testing.T) {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}

// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// leetcode287: 寻找重复数
// 利用数字出现次数的的前缀和
func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

// leetcode287: 寻找重复数
// 利用快慢指针，找环的入口
// 由于题目的数据范围在1~n之间，可以把这些元素组织成链表
func findDuplicate_(nums []int) int {
	fast, slow := nums[0], 0
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	p1, p2 := 0, nums[slow]
	for p1 != p2 {
		p1 = nums[p1]
		p2 = nums[p2]
	}
	return p1
}

// leetcode287: 寻找重复数
// 二进制解法
// 解题思路：如果重复的数字在第i位为1，那么第i位上1的个数大于1~n所有数字第i位上1的个数。
func findDuplicate__(nums []int) int {
	n := len(nums)
	bit_max := 31
	for (n-1)>>bit_max == 0 {
		bit_max--
	}
	ans := 0
	for bit := 0; bit <= bit_max; bit++ {
		x, y := 0, 0
		for i := 0; i < n; i++ {
			if nums[i]&(1<<bit) > 0 {
				x++
			}
			if i >= 1 && (i&(1<<bit)) > 0 {
				y++
			}
		}
		if x > y {
			ans |= 1 << bit
		}
	}
	return ans
}

func Test_findDuplicate1(t *testing.T) {
	fmt.Println(findDuplicate([]int{1, 3, 2, 2, 4}))
	fmt.Println(findDuplicate_([]int{1, 3, 2, 2, 4}))
	fmt.Println(findDuplicate__([]int{1, 3, 2, 2, 4}))
}

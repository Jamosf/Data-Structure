// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// tag-[双指针]
// 第一题
// leetcode80: 删除有序数组中的重复项II
func removeDuplicates(nums []int) int {
	sort.Ints(nums)
	first, second := 0, 1
	l := len(nums)
	for second < l {
		if nums[first] == nums[second] {
			first++
			second++
			for second < l && nums[second-1] == nums[second] {
				if second+1 < l {
					nums = append(nums[:second], nums[second+1:]...)
				} else {
					nums = nums[:second]
				}

				l--
			}
		}
		first++
		second++
	}
	return l
}

func Test_remove(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 1, 1}))
}

// tag-[数组]
// 第三题
// leetcode18: 四数之和
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-1; i++ {
		res := threesumT(append(nums[:i], nums[i+1:]...), target-nums[i])
		for k := range res {
			flag := false
			for j := range ans {
				if ans[j][0] == nums[i] && ans[j][1] == res[k][0] && ans[j][2] == res[k][1] && ans[j][3] == res[k][3] {
					flag = true
					break
				}
			}
			if !flag {
				ans = append(ans, []int{nums[i], res[k][0], res[k][1], res[k][2]})
			}
		}

	}
	return ans
}

func threesumT(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 1; i < len(nums)-1; i++ {
		l, r := i-1, i+1
		for l >= 0 && r <= len(nums)-1 {
			if nums[l]+nums[i]+nums[r] == target {
				flag := false
				for j := range ans {
					if ans[j][0] == nums[l] && ans[j][1] == nums[i] && ans[j][2] == nums[r] {
						flag = true
						break
					}
				}
				if !flag {
					ans = append(ans, []int{nums[l], nums[i], nums[r]})
				}
				l--
				r++
			} else if nums[l]+nums[i]+nums[r] > target {
				l--
			} else {
				r++
			}
		}
	}
	return ans
}

func Test_four(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

// tag-[双指针]
// 第四题
// leetcode16: 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := math.MaxInt32
	ans := 0
	for i := 1; i < len(nums)-1; i++ {
		l, r := i-1, i+1
		for l >= 0 && r <= len(nums)-1 {
			sum := nums[l] + nums[i] + nums[r]
			if sum == target {
				return sum
			} else if sum > target {
				diff := sum - target
				if diff < min {
					min = diff
					ans = sum
				}
				l--
				for l >= 0 && nums[l+1] == nums[l] {
					l--
				}
			} else {
				diff := target - sum
				if diff < min {
					min = diff
					ans = sum
				}
				r++
				for r <= len(nums)-1 && nums[r-1] == nums[r] {
					r++
				}
			}
		}
	}
	return ans
}

func Test_close(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

// tag-[双指针]
// 第五题
// leetcode15: 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 1; i < len(nums)-1; i++ {
		l, r := i-1, i+1
		for l >= 0 && r <= len(nums)-1 {
			if nums[l]+nums[i]+nums[r] == 0 {
				flag := false
				for j := range ans {
					if ans[j][0] == nums[l] && ans[j][1] == nums[i] && ans[j][2] == nums[r] {
						flag = true
						break
					}
				}
				if !flag {
					ans = append(ans, []int{nums[l], nums[i], nums[r]})
				}
				l--
				r++
			} else if nums[l]+nums[i]+nums[r] > 0 {
				l--
			} else {
				r++
			}
		}
	}
	return ans
}

func Test_threeNum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 0}))
}

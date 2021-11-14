// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// 二分查找
// 第一题
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	ans := make([]int, 0, 2)
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			right--
		}
	}
	if nums[left] == target {
		ans = append(ans, left)
	} else {
		ans = append(ans, -1)
	}
	left, right = 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1 + 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid
		} else {
			left++
		}
	}
	if nums[right] == target {
		ans = append(ans, right)
	} else {
		ans = append(ans, -1)
	}
	return ans
}

func Test_searchRange(t *testing.T) {
	fmt.Println(searchRange([]int{7}, 7))
}

// 第二题
func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func Test_mySqrt(t *testing.T) {
	fmt.Println(mySqrt(0))
}

// 第三题
func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	left, right := 0, row*col-1
	for left <= right {
		mid := left + (right-left)>>1
		i, j := mid/col, mid%col
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	i, j := left/col, left%col
	if i < row && j < col {
		return matrix[i][j] == target
	}
	return false
}

func Test_searchMatrix(t *testing.T) {
	fmt.Println(searchMatrix([][]int{{1, 1}}, 2))
}

// 第四题
// 单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	for i := range nums1 {
		m[nums1[i]] = i
	}
	stack := make([]int, 0, len(nums2))
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	for i := range nums2 {
		for len(stack) != 0 && nums2[stack[len(stack)-1]] < nums2[i] {
			if idx, ok := m[nums2[stack[len(stack)-1]]]; ok {
				ans[idx] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func Test_nextGreaterElement(t *testing.T) {
	fmt.Println(nextGreaterElement([]int{4}, []int{4}))
}

// 第六题
func nextGreaterElements(nums []int) []int {
	tmp := append(nums, nums...)
	stack := make([]int, 0, len(nums)*2)
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	for i := range tmp {
		for len(stack) != 0 && tmp[stack[len(stack)-1]] < tmp[i] {
			if stack[len(stack)-1] < len(nums) {
				ans[stack[len(stack)-1]] = tmp[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func Test_nextGreaterElements(t *testing.T) {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

// 第七题
func nextGreaterElementIII(n int) int {
	var nums []int
	num := n
	for n != 0 {
		nums = append(nums, n%10)
		n = n / 10
	}
	sort.Ints(nums)
	var dfs func(level int)
	var all []int
	var tmp []int
	visited := make([]bool, len(nums))
	dfs = func(level int) {
		if level == len(nums) {
			sum := 0
			t := 1
			for j := len(tmp) - 1; j >= 0; j-- {
				sum += tmp[j] * t
				t *= 10
			}
			all = append(all, sum)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				tmp = append(tmp, nums[i])
				visited[i] = true
				dfs(level + 1)
				tmp = tmp[:len(tmp)-1]
				visited[i] = false
			}
		}
	}
	dfs(0)
	idx := sort.SearchInts(all, num+1)
	for i := idx; i < len(all); i++ {
		if all[i] > num && all[i] <= math.MaxInt32 {
			return all[i]
		}
	}
	return -1
}

func Test_nextGreaterElementIII(t *testing.T) {
	fmt.Println(nextGreaterElementIII(1234))
}

// 第八题
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))
	for i := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// 第九题
func find132pattern(nums []int) bool {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	ans, stack := make([]int, len(nums)), make([]int, 0, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	for i := range nums {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			ans[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < len(ans); i++ {
		if ans[i] > 0 {
			for j := ans[i] + 1; j < len(nums); j++ {
				if nums[j] < nums[i] {
					return true
				}
			}
		}
	}
	return false
}

func Test_find132pattern(t *testing.T) {
	fmt.Println(find132pattern([]int{3, 5, 0, 3, 4}))
}

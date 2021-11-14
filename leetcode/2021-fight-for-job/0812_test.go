// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// 第一题
// leetcode209: 长度最小的子数组
// 双指针
func minSubArrayLen(target int, nums []int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	minn := math.MaxInt32
	left, right := 0, 1
	for right < len(nums) {
		if sum[right]-sum[left] < target {
			right++
		} else {
			minn = min(minn, right-left)
			left++
		}
	}
	if minn == math.MaxInt32 {
		return 0
	}
	return minn
}

// leetcode209: 长度最小的子数组
// 二分搜索
func minSubArrayLen_(target int, nums []int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	minn := math.MaxInt32
	for i := range sum {
		idx := binarySearch(sum, sum[i]+target)
		if idx > 0 {
			minn = min(minn, idx-i)
		}
	}
	if minn == math.MaxInt32 {
		return 0
	}
	return minn
}

func binarySearch(sum []int, target int) int {
	left, right := 0, len(sum)-1
	for left < right { // 查找左边界
		mid := (left + right) >> 1
		if sum[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if sum[left] >= target {
		return left
	}
	return -1
}

func Test_sortSearch(t *testing.T) {
	fmt.Println(sort.SearchInts([]int{1, 2, 3}, 5))
}

// 第二题
// leetcode238: 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	tmp := 1
	for i := len(nums) - 2; i >= 0; i-- {
		tmp *= nums[i+1]
		ans[i] *= tmp
	}
	return ans
}

// 第三题
// leetcode304: 二维区域和检索-矩阵不可变
type NumMatrix struct {
	sum    [][]int
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row, col := len(matrix), len(matrix[0])
	sum := make([][]int, row)
	for i := 0; i < row; i++ {
		sum[i] = make([]int, col)
	}
	numMatrix := NumMatrix{sum: sum, matrix: matrix}
	for i := 0; i < row; i++ {
		numMatrix.sum[i][0] = matrix[i][0]
		for j := 1; j < col; j++ {
			numMatrix.sum[i][j] = numMatrix.sum[i][j-1] + matrix[i][j]
		}
	}
	return numMatrix
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += n.sum[i][col2] - n.sum[i][col1] + n.matrix[i][col1]
	}
	return sum
}

// 第四题
// leetcode523: 连续的子数组和
// 方法1 前缀和+hashmap
func checkSubarraySum(nums []int, k int) bool {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := range nums {
		sum[i+1] = sum[i] + nums[i]
	}
	m := make(map[int]int)
	for i := 0; i < len(sum); i++ {
		v := sum[i] % k
		idx, ok := m[v]
		if ok {
			if i-idx > 1 {
				return true
			}
		} else {
			m[v] = i
		}
	}
	return false
}

func Test_checkSubarraySum(t *testing.T) {
	fmt.Println(checkSubarraySum([]int{5, 0, 0, 0}, 3))
}

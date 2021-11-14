// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// leetcode85: 矩阵中的最大矩形面积，单调栈
func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	height := make([]int, n)
	maxn := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != '0' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		maxn = max(maxn, largestRectangleArea(height))
	}
	return maxn
}

// leetcode84: 最大的矩形面积，单调递减栈
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0, n)
	right := make([]int, n)
	left := make([]int, n)
	for i := range right {
		right[i] = n
	}
	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	maxn := 0
	for i := range right {
		maxn = max(maxn, (right[i]-left[i]-1)*heights[i])
	}
	return maxn
}

func Test_largestRectangleArea(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
	fmt.Println(largestRectangleArea([]int{2, 5}))
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
}

// leetcode42：接雨水, 单调栈解法
// 思路：如果栈内元素超过两个，并且当前元素大于栈顶元素，那么栈顶元素处可以积水
func trap_(height []int) (ans int) {
	n := len(height)
	stack := make([]int, 0, n)
	ans = 0
	for i := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			h := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			ans += (min(height[i], height[stack[len(stack)-1]]) - h) * (i - stack[len(stack)-1] - 1)
		}
		stack = append(stack, i)
	}
	return ans
}

func Test_trap(t *testing.T) {
	fmt.Println(trap_([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

// leetcode901: 股票价格跨度, 单调递增栈
type StockSpanner struct {
	prices []int
	stack  []int
	res    []int
}

func ConstructorStockSpanner() StockSpanner {
	return StockSpanner{}
}

func (s *StockSpanner) Next(price int) int {
	tmp := 1
	s.prices = append(s.prices, price)
	for len(s.stack) > 0 && price >= s.prices[s.stack[len(s.stack)-1]] {
		tmp += s.res[s.stack[len(s.stack)-1]]
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, len(s.prices)-1)
	s.res = append(s.res, tmp)
	return tmp
}

func Test_StockSpanner(t *testing.T) {
	s := ConstructorStockSpanner()
	s.Next(100)
	s.Next(80)
	s.Next(60)
	s.Next(70)
	s.Next(60)
}

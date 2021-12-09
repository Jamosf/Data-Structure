// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"testing"
)

// tag-[单调栈]
// leetcode316和1081
func removeDuplicateLetters(s string) string {
	var count [26]int
	for i := range s {
		count[s[i]-'a']++
	}
	stack := make([]byte, 0, len(s))
	instack := [26]bool{}
	for i := range s {
		if !instack[s[i]-'a'] {
			for len(stack) > 0 && stack[len(stack)-1] > s[i] {
				last := stack[len(stack)-1] - 'a'
				if count[last] == 0 { // 如果这个字符在后面没有，则不能弹出
					break
				}
				stack = stack[:len(stack)-1]
				instack[last] = false
			}
			stack = append(stack, s[i])
			instack[s[i]-'a'] = true
		}
		count[s[i]-'a']-- // 记录剩下的字符
	}
	return string(stack)
}

func Test_removeDuplicateLetters(t *testing.T) {
	fmt.Println(removeDuplicateLetters("abcfdbcgthsiidbbcxxwwsxxxxkkkl"))
}

// tag-[单调栈]
// leetcode321
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	pickNum := func(nums []int, k int) []int {
		n := len(nums)
		stack := make([]int, 0, n)
		drop := n - k
		for i := range nums {
			for len(stack) > 0 && stack[len(stack)-1] < nums[i] && drop > 0 {
				stack = stack[:len(stack)-1]
				drop--
			}
			stack = append(stack, nums[i])
		}
		return stack[:k]
	}
	maxSlice := func(n1 []int, n2 []int) []int {
		for i := range n1 {
			if n1[i] > n2[i] {
				return n1
			} else if n1[i] < n2[i] {
				return n2
			}
		}
		return n1
	}
	isMax := func(n1 []int, n2 []int) bool {
		l1, l2 := len(n1), len(n2)
		l := min(l1, l2)
		for i := 0; i < l; i++ {
			if n1[i] > n2[i] {
				return true
			} else if n1[i] < n2[i] {
				return false
			}
		}
		return l1 > l2
	}
	merge := func(n1 []int, n2 []int) []int {
		l1, l2 := len(n1), len(n2)
		r := make([]int, 0, l1+l2)
		var i, j int
		for i < l1 && j < l2 {
			if n1[i] > n2[j] {
				r = append(r, n1[i])
				i++
			} else if n1[i] < n2[j] {
				r = append(r, n2[j])
				j++
			} else {
				if isMax(n1[i:], n2[j:]) {
					r = append(r, n1[i])
					i++
				} else {
					r = append(r, n2[j])
					j++
				}
			}
		}
		if i == l1 {
			r = append(r, n2[j:]...)
		} else {
			r = append(r, n1[i:]...)
		}
		return r
	}

	m, n := len(nums1), len(nums2)
	maxn := make([]int, k)
	for i := 0; i <= k; i++ {
		if i <= m && k-i <= n {
			maxn = maxSlice(maxn, merge(pickNum(nums1, i), pickNum(nums2, k-i)))
		}
	}
	return maxn
}

func Test_maxNumber(t *testing.T) {
	// fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	// fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
	// fmt.Println(maxNumber([]int{6, 5}, []int{6, 7, 4}, 5))
	fmt.Println(maxNumber([]int{5, 6, 8}, []int{6, 4, 0}, 3))
}

// tag-[数组]
// leetcode334
func increasingTriplet(nums []int) bool {
	small, mid := math.MaxInt32, math.MaxInt32
	for i := range nums {
		if nums[i] <= small {
			small = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}

func Test_increasingTriplet(t *testing.T) {
	fmt.Println(increasingTriplet([]int{2, 1, 4, 2, 4, 3}))
}

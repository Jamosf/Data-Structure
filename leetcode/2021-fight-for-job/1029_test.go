// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// leetcode986:双指针
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int, 0)
	m, n := len(firstList), len(secondList)
	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		v1, v2 := firstList[p1], secondList[p2]
		t1, t2 := max(v1[0], v2[0]), min(v1[1], v2[1])
		if t1 <= t2 {
			ans = append(ans, []int{t1, t2})
		} else {
			if t2 == v1[1] {
				p1++
			} else {
				p2++
			}
			continue
		}
		if t2 == v1[1] && t2 == v2[1] {
			p1++
			p2++
			continue
		}
		if t2 < v1[1] {
			p2++
		} else if t2 < v2[1] {
			p1++
		}
	}
	return ans
}

// leetcode986:双指针-精简逻辑
func intervalIntersection_(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int, 0)
	m, n := len(firstList), len(secondList)
	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		v1, v2 := firstList[p1], secondList[p2]
		t1, t2 := max(v1[0], v2[0]), min(v1[1], v2[1])
		if t1 <= t2 {
			ans = append(ans, []int{t1, t2})
		}
		if v1[1] > v2[1] {
			p2++
		} else {
			p1++
		}
	}
	return ans
}

func Test_intervalIntersection(t *testing.T) {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
	fmt.Println(intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}))
	fmt.Println(intervalIntersection([][]int{{1, 7}}, [][]int{{3, 10}}))
}

// leetcode443:双指针
func compress(chars []byte) int {
	n := len(chars)
	write, left := 0, 0
	for read, ch := range chars {
		if read == n-1 || ch != chars[read+1] {
			chars[write] = ch
			write++
			num := read - left + 1
			if num > 1 {
				anchor := write
				for ; num > 0; num /= 10 {
					chars[write] = '0' + byte(num%10)
					write++
				}
				s := chars[anchor:write]
				for i, n := 0, len(s); i < n/2; i++ {
					s[i], s[n-1-i] = s[n-1-i], s[i]
				}
			}
			left = read + 1
		}
	}
	return write
}

func Test_compress(t *testing.T) {
	// fmt.Println(compress([]byte{'a'}))
	// fmt.Println(compress([]byte{'a', 'a'}))
	fmt.Println(compress([]byte{'a', 'b', 'c'}))
	// fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
	// fmt.Println(compress([]byte{'a', 'a', '2', '2', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
}

// leetcode457: 快慢指针
func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(idx int) int {
		return ((idx+nums[idx])%n + n) % n
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}
		slow, fast := i, next(i)
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[next(fast)] > 0 {
			if slow == fast {
				if slow == next(slow) {
					break
				}
				return true
			}
			slow = next(slow)
			fast = next(next(fast))
		}
		// 非循环的标记位0，防止再次遍历，提升效率
		mark := i
		for nums[mark]*nums[next(mark)] > 0 {
			tmp := mark
			mark = next(mark)
			nums[tmp] = 0
		}
	}
	return false
}

func Test_circularArrayLoop(t *testing.T) {
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
	fmt.Println(circularArrayLoop([]int{-1, 2}))
	fmt.Println(circularArrayLoop([]int{-2, 1, -1, -2, -2}))
	fmt.Println(circularArrayLoop([]int{1}))
	fmt.Println(circularArrayLoop([]int{1, 1}))
	fmt.Println(circularArrayLoop([]int{-1, -2, -3, -4, -5}))
}

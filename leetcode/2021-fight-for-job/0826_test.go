// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

// tag-[双指针]
// 第一题
// leetcode45: 分发饼干
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i, j int
	for i, j = 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			i++
			j++
		} else {
			j++
		}
	}
	return i
}

// tag-[数组]
// 第二题
// leetcode135：分发糖果
func candy(ratings []int) int {
	ans := make([]int, len(ratings))
	for i := range ans {
		ans[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			ans[i] = ans[i-1] + 1
		}
	}
	sum := ans[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			ans[i] = max(ans[i], ans[i+1]+1)
		}
		sum += ans[i]
	}
	return sum
}

func Test_candy(t *testing.T) {
	fmt.Println(candy([]int{1, 0, 2}))
}

// tag-[数组]
// 第三题
// leetcode435: 无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans := 0
	for i := 0; i < len(intervals); i++ {
		v := intervals[i][1]
		for i < len(intervals)-1 && v > intervals[i+1][0] {
			ans++
			i++
		}
	}
	return ans
}

func Test_eraseOverlapIntervals(t *testing.T) {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
}

// tag-[数组]
// 第四题
// leetcode605：种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	ans := 0
	pre := -1
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			if pre != -1 {
				ans += (i - pre - 2) >> 1
			} else {
				ans += i >> 1
			}
			pre = i
		}
	}
	if pre == -1 {
		ans += (len(flowerbed) + 1) >> 1
	} else {
		ans += (len(flowerbed) - 1 - pre) >> 1
	}
	return ans >= n
}

func Test_canPlaceFlowers(t *testing.T) {
	fmt.Println(canPlaceFlowers([]int{0}, 1))
}

// tag-[数组]
// 第五题
// leetcode763: 划分字母区间
func partitionLabels(s string) []int {
	tmp := make([][2]int, 26)
	for i := range tmp {
		tmp[i][0], tmp[i][1] = -1, -1
	}
	for i := range s {
		if tmp[s[i]-'a'][0] == -1 {
			tmp[s[i]-'a'][0] = i
		} else {
			tmp[s[i]-'a'][1] = i
		}
	}
	var ans []int
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		v := tmp[s[i]-'a']
		if v[0] > right {
			ans = append(ans, right-left+1)
			left, right = i, i
		}
		if i == len(s)-1 {
			ans = append(ans, right-left+1)
			left, right = i, i
		}
		if v[0] != -1 && v[0] < left {
			left = v[0]
		}
		if v[1] != -1 && v[1] > right {
			right = v[1]
		}
	}
	return ans
}

func Test_partitionLabels(t *testing.T) {
	fmt.Println(partitionLabels("aaaaaaaaaaa"))
}

// tag-[双指针]
// 第六题
// leetcode76: 最小覆盖子串
func minWindow(s string, t string) string {
	ms, mt := make(map[byte]int), make(map[byte]int)
	for i := range t {
		mt[t[i]]++
	}
	left, right := 0, 0
	var ans string
	for right < len(s) {
		ms[s[right]]++
		for isInclude(ms, mt) {
			if len(ans) == 0 || right-left+1 < len(ans) {
				ans = s[left : right+1]
			}
			ms[s[left]]--
			left++
		}
		right++
	}
	return ans
}

func isInclude(s, t map[byte]int) bool {
	for k, v := range t {
		if v > s[k] {
			return false
		}
	}
	return true
}

func Test_minWindow(t *testing.T) {
	fmt.Println(minWindow("a", "aa"))
}

// tag-[二分查找]
// 第七题
// leetcode633：平方数之和
func judgeSquareSum(c int) bool {
	left, right := 0, sqrt(c)
	for left <= right {
		v := left*left + right*right
		if v > c {
			right--
		} else if v < c {
			left++
		} else {
			return true
		}
	}
	return false
}

func sqrt(n int) int {
	if n == 0 {
		return 0
	}
	left, right := 0, n
	for left < right {
		mid := left + (right-left)>>1
		if mid == 0 {
			return right
		}
		v := n / mid
		if v == mid {
			return mid
		} else if v < mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func Test_judgeSquareSum(t *testing.T) {
	fmt.Println(sqrt(1))
	fmt.Println(judgeSquareSum(5))
}

// tag-[双指针]
// 第八题
// leetcode680: 验证回文字符串II
func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

func Test_validPalindrome(t *testing.T) {
	fmt.Println(validPalindrome("ab"))
}

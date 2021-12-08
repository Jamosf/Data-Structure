// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// tag-[字符串]
// 第一题
// leetcode3: 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	left, right := 0, 0
	maxn := 0
	for left = 0; left < len(s); left++ {
		if left != 0 {
			delete(m, s[left-1])
		}
		for right < len(s) && m[s[right]] == 0 {
			m[s[right]]++
			right++
		}
		maxn = max(maxn, right-left)
	}
	return maxn
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring(" "))
}

// tag-[字符串]
// 第二题
// leetcode567: 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	m := make(map[uint8]int)
	for i := range s1 {
		m[s1[i]]++
	}
	left, right := 0, len(s1)
	for i := left; i <= right; i++ {
		if m[s2[i]] != 0 {
			m[s2[i]]--
		}
	}
	if isMapEmpty(m) {
		return true
	}
	for right < len(s2) {
		left++
		right++
		if m[s2[left-1]] >= 0 {

		}
	}
	return false
}

func isMapEmpty(m map[uint8]int) bool {
	cnt := 0
	for _, v := range m {
		cnt += v
	}
	return cnt == 0
}

// tag-[字符串]
// 第三题
// leetcode387: 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	var m [26]int
	for _, v := range s {
		m[v-'a']++
	}
	for i, v := range s {
		if m[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

// tag-[字符串]
// 第四题
// leetcode383: 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	var m [26]int
	for _, v := range magazine {
		m[v-'a']++
	}
	for _, v := range ransomNote {
		m[v-'a']--
	}
	for _, v := range m {
		if v < 0 {
			return false
		}
	}
	return true
}

// tag-[字符串]
// 第五题
// leetcode242: 有效的字母异位词
func isAnagram(s string, t string) bool {
	var m [26]int
	for _, v := range s {
		m[v-'a']++
	}
	var n [26]int
	for _, v := range t {
		n[v-'a']++
	}
	return m == n
}

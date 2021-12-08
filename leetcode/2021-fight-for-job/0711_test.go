// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

// tag-[字符串]
// leetcode718: 最长重复子数组
func findLength(nums1 []int, nums2 []int) int {
	lenA, lenB := len(nums1), len(nums2)
	ret := 0
	for i := 0; i < lenA; i++ {
		k := 0
		for j := 0; j < min(lenA-i, lenB); j++ {
			if nums1[i+j] == nums2[j] {
				k++
			} else {
				k = 0
			}
			ret = max(ret, k)
		}
	}

	for i := 0; i < lenB; i++ {
		k := 0
		for j := 0; j < min(lenB-i, lenA); j++ {
			if nums1[j] == nums2[j+i] {
				k++
			} else {
				k = 0
			}
			ret = max(ret, k)
		}
	}
	return ret
}

// leetcode187: 重复DNA序列
func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]uint8, len(s)-10)
	for i := 0; i < len(s)-10+1; i++ {
		m[s[i:i+10]]++
	}
	var result []string
	for k, v := range m {
		if v > 1 {
			result = append(result, k)
		}
	}
	return result
}

// leetcode1461: 检查一个字符串是否包含所有长度为K的二进制子串
func hasAllCodes(s string, k int) bool {
	m := make(map[string]uint8, len(s)-k)
	for i := 0; i < len(s)-k+1; i++ {
		m[s[i:i+k]]++
	}
	result := 0
	for _, v := range m {
		if v >= 1 {
			result++
		}
	}
	return result == 1<<k
}

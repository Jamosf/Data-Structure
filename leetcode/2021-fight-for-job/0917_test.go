// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"strings"
	"testing"
)

func reversePrefix(word string, ch byte) string {
	idx := strings.Index(word, string(ch))
	if idx == -1 {
		return word
	}
	b := []byte(word[:idx+1])
	for i := 0; i <= idx/2; i++ {
		b[i], b[idx-i] = b[idx-i], b[i]
	}
	return string(b) + word[idx+1:]
}

func interchangeableRectangles(rectangles [][]int) int64 {
	m := make(map[float64]int64)
	for i := 0; i < len(rectangles); i++ {
		m[float64(rectangles[i][0])/float64(rectangles[i][1])]++
	}
	cnt := int64(0)
	for _, v := range m {
		if v > 1 {
			cnt += v * (v - 1) / 2
		}
	}
	return cnt
}

func maxProduct(s string) int {
	n := len(s)
	m := map[int]int{}
	for i := 1; i < 1<<n-1; i++ {
		t := make([]byte, 0)
		for idx := 0; idx < n; idx++ {
			if 1<<idx&i == 1<<idx {
				t = append(t, s[n-idx-1])
			}
		}
		if isPlalindrome(t) {
			m[i] = len(t)
		}
	}
	maxn := 0
	for i := 1; i < 1<<n-1; i++ {
		for j := i - 1; j >= 0; j-- {
			if i&j == 0 && m[i] != 0 && m[j] != 0 {
				fmt.Println(i, j, m[i], m[j])
				maxn = max(maxn, m[i]*m[j])
			}
		}
	}
	return maxn
}

func isPlalindrome(b []byte) bool {
	j := len(b) - 1
	for i := 0; i < j; i++ {
		if b[i] != b[j] {
			return false
		}
		j--
	}
	return true
}

func Test_maxProduct(t *testing.T) {
	fmt.Println(maxProduct("leetcodecom"))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3}))
}

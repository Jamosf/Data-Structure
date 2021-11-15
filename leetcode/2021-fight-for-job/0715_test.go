// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

// 第一题
// leetcode1846: 减少和重新排列数组后的最大元素
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sort.Ints(arr)
	arr[0] = 1
	for i := 0; i < len(arr); i++ {
		if minusAbs(arr[i], arr[i+1]) > 1 {
			arr[i+1] = arr[i] + 1
		}
	}
	return arr[len(arr)-1]
}

// 第二题
// leetcode1818: 绝对差值和
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	mod := int(1e9 + 7)
	// 1. 先复制一份数组并排序
	rec := append(sort.IntSlice{}, nums1...)
	rec.Sort()
	// 2. 开始找
	sum := 0
	maxn := 0
	for i := 0; i < len(nums1); i++ {
		diff := minusAbs(nums1[i], nums2[i])
		sum += diff % mod
		// 3. 在rec中找替换
		j := rec.Search(nums2[i])
		if j < len(nums1) {
			maxn = max(maxn, diff-minusAbs(rec[j], nums2[i]))
		}
		if j > 0 {
			maxn = max(maxn, diff-minusAbs(rec[j-1], nums2[i]))
		}
	}
	return (sum - maxn) % mod
}

// 第三题
// leetcode275: H指数II
func hIndex(citations []int) int {
	left, right := 0, len(citations)
	mid := (left + right) / 2
	for left < right {
		if citations[mid] > len(citations)-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return len(citations) - left
}

// 第四题
// leetcode981: 基于时间的键值存储
type TimeMap struct {
	kv map[string][]v
}

type v struct {
	value     string
	timeStamp int
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	maxn := dp[0]
	for i := range dp {
		maxn = max(dp[i], maxn)
	}
	return maxn
}

/** Initialize your data structure here. */
func ConstructorTimeMap() TimeMap {
	return TimeMap{kv: make(map[string][]v)}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.kv[key] = append(t.kv[key], v{value: value, timeStamp: timestamp})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	if vv, ok := t.kv[key]; ok && len(vv) != 0 {
		for i := len(vv) - 1; i >= 0; i-- {
			if vv[i].timeStamp <= timestamp {
				return vv[i].value
			}
		}
	}
	return ""
}

// 第五题
// leetcode 面试题17.10: 主要元素
func majorityElement(nums []int) int {
	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	if 2*count > len(nums) {
		return candidate
	}
	return -1
}

// 第六题
// leetcode
func numSubarraysWithSum(nums []int, goal int) int {
	// 3. 利用前缀和计算
	sumn := 0
	ans := 0
	cnt := make(map[int]int)
	for i := range nums {
		sumn += nums[i]
		cnt[sumn]++
		ans += cnt[sumn-goal]
	}
	return ans
}

func Test_numSubarraysWithSum(t *testing.T) {
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}

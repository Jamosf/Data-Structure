// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

// tag-[双指针]
// leetcode11: 盛最多水的容器
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

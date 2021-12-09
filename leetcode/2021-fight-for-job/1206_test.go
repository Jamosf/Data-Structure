package ojeveryday

import "math"

// tag-[数组]
// leetcode2091: 从数组中移除最大值和最小值
func minimumDeletions(nums []int) int {
	maxn, minn := nums[0], nums[0]
	maxIdx, minIdx := -1, -1
	for i := range nums {
		if nums[i] >= maxn {
			maxn = nums[i]
			maxIdx = i
		}
		if nums[i] <= minn {
			minn = nums[i]
			minIdx = i
		}
	}
	res := math.MaxInt32
	// 分情况讨论
	// 1. 都在左边
	res = min(res, max(maxIdx, minIdx)+1)
	// 2. 都在右边
	res = min(res, len(nums)-min(maxIdx, minIdx))
	// 3. 一个左边，一个右边
	res = min(res, maxIdx+1+len(nums)-minIdx)
	// 4. 一个右边，一个左边
	return min(res, len(nums)-maxIdx+1+minIdx)
}

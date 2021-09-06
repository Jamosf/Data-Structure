package _021_fight_for_job

import (
	"math"
)

// 前缀和
// 第一题
func largestAltitude(gain []int) int {
	maxn := math.MaxInt32
	sum := gain[0]
	for i := 1; i < len(gain); i++ {
		sum += gain[i]
		maxn = max(maxn, sum)
	}
	return maxn
}

// 第二题
func isCovered(ranges [][]int, left int, right int) bool {
	//sort.Slice(ranges, func(i, j int) bool {
	//	return ranges[i][0] < ranges[j][0]
	//})
	//start, end := ranges[0][0], ranges[0][1]
	//for i := 1; i < len(ranges); i++ {
	//	if ranges[i][0]
	//}
	return false
}

// 第三题
func pivotIndex(nums []int) int {
	sum := make([]int, len(nums)+2)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	sum[len(sum)-1] = sum[len(sum)-2]
	for i := 1; i < len(sum)-1; i++ {
		if sum[i-1] == sum[len(nums)-1]-sum[i] {
			return i - 1
		}
	}
	return -1
}

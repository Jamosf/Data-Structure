package ojeveryday

import (
	"math"
	"sort"
)

// tag-[前缀和]
// leetcode2090: 半径为k的子数组平均值
func getAverages(nums []int, k int) []int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		if i-k >= 0 && i+k < n {
			out[i] = (preSum[i+k+1] - preSum[i-k]) / (2*k + 1)
		} else {
			out[i] = -1
		}
	}
	return out
}

// tag-[贪心]
// leetcode2086: 从房屋收集雨水需要的最少水桶数
// 贪心解法：优先从右边添加
func minimumBuckets(street string) int {
	n := len(street)
	b := []byte(street)
	ans := 0
	for i := 0; i < n; i++ {
		if street[i] == 'H' {
			if i-1 >= 0 && b[i-1] == '.' {
				if i+1 >= n || b[i+1] == 'H' {
					b[i-1] = 'B'
					ans++
				}
				if i+1 < n && b[i+1] == '.' {
					b[i+1] = 'B'
					ans++
				}
			}
			if i-1 >= 0 && b[i-1] == 'H' {
				if i+1 >= n || b[i+1] == 'H' {
					return -1
				}
				if i+1 < n && b[i+1] == '.' {
					b[i+1] = 'B'
					ans++
				}
			}
			if i-1 < 0 && i+1 < n && b[i+1] == '.' {
				b[i+1] = 'B'
				ans++
			}
			if i-1 < 0 && (i+1 >= n || (i+1 < n && b[i+1] == 'H')) {
				return -1
			}
		}
	}
	return ans
}

// tag-[动态规划]
// leetcode2086: 从房屋收集雨水需要的最少水桶数
// dp[i][3]: 第二维的含义：0表示以H结尾且前面无桶、1表示以H结尾且前面有桶、2表示以B结尾、3表示以.结尾（前面H必有桶）
// 状态转移:
func minimumBuckets_(street string) int {
	n := len(street)
	dp := make([][4]int, n)
	if street[0] == '.' {
		dp[0][0] = math.MaxInt32
		dp[0][1] = math.MaxInt32
		dp[0][2] = 1
		dp[0][3] = 0
	} else {
		dp[0][0] = 0
		dp[0][1] = math.MaxInt32
		dp[0][2] = math.MaxInt32
		dp[0][3] = math.MaxInt32
	}
	for i := 1; i < n; i++ {
		if street[i] == '.' {
			dp[i][3] = min(min(dp[i-1][1], dp[i-1][2]), dp[i-1][3])
			dp[i][2] = dp[i-1][0] + 1
			dp[i][1] = math.MaxInt32
			dp[i][0] = math.MaxInt32
		} else {
			dp[i][0] = min(dp[i-1][1], dp[i-1][3])
			dp[i][1] = min(dp[i-1][2], dp[i-1][3]+1)
			dp[i][2] = math.MaxInt32
			dp[i][3] = math.MaxInt32
		}
	}
	ans := math.MaxInt32
	ans = min(ans, dp[n-1][1])
	ans = min(ans, dp[n-1][2])
	ans = min(ans, dp[n-1][3])
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// tag-[矩阵]
// leetcode2087: 网格图中机器人回家的最小代价
// 直接走
func minCost(startPos, homePos, rowCosts, colCosts []int) int {
	x0, y0, x1, y1 := startPos[0], startPos[1], homePos[0], homePos[1]
	ans := -rowCosts[x0] - colCosts[y0] // 初始的行列无需计算
	if x0 > x1 {
		x0, x1 = x1, x0
	} // 交换位置，保证 x0 <= x1
	if y0 > y1 {
		y0, y1 = y1, y0
	} // 交换位置，保证 y0 <= y1
	for _, cost := range rowCosts[x0 : x1+1] {
		ans += cost
	} // 统计答案
	for _, cost := range colCosts[y0 : y1+1] {
		ans += cost
	} // 统计答案
	return ans
}

// tag-[数组]
// leetcode2079: 给植物浇水
// 模拟
func wateringPlants(plants []int, capacity int) int {
	n := len(plants)
	ans := 0
	left := capacity
	for i := 0; i < n; i++ {
		if left < plants[i] {
			ans += i * 2
			left = capacity
		}
		left -= plants[i]
		ans += 1
	}
	return ans
}

// tag-[哈希表/二分查找]
// leetcode2080: 区间内查询数字的频率
// hash预处理+二分
type RangeFreqQuery struct {
	m map[int][]int
}

func ConstructorR(arr []int) RangeFreqQuery {
	r := RangeFreqQuery{m: make(map[int][]int)}
	for i := range arr {
		r.m[arr[i]] = append(r.m[arr[i]], i)
	}
	return r
}

func (r *RangeFreqQuery) Query(left int, right int, value int) int {
	s := r.m[value]
	if len(s) == 0 {
		return 0
	}
	idx1, idx2 := sort.SearchInts(s, left), sort.SearchInts(s, right+1)
	return idx2 - idx1
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */

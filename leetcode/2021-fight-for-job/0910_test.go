// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// leetcode1711: 大餐计数
func countPairs(deliciousness []int) int {
	n := len(deliciousness)
	ans := 0
	mod := int(1e9 + 7)
	m := make(map[int]int)
	for i := range deliciousness {
		m[deliciousness[i]]++
	}
	sort.Ints(deliciousness)
	for i := 0; i < n; i++ {
		for i < n-1 && deliciousness[i] == deliciousness[i+1] {
			i++
		}
		for j := i + 1; j < n; j++ {
			for j < n-1 && deliciousness[j] == deliciousness[j+1] {
				j++
			}
			v := deliciousness[i] + deliciousness[j]
			if v&(v-1) == 0 {
				ans += (m[deliciousness[i]] * m[deliciousness[j]]) % mod
			}
		}
	}
	for k, v := range m {
		t := k << 1
		if v > 1 && k != 0 && t&(t-1) == 0 {
			ans += (v * (v - 1) / 2) % mod
		}
	}
	return ans % mod
}

func countPairs_(deliciousness []int) int {
	mod := int(1e9 + 7)
	maxn := deliciousness[0]
	for i := range deliciousness {
		maxn = max(maxn, deliciousness[i])
	}
	maxSum := 2 * maxn
	ans := 0
	cnt := make(map[int]int)
	for i := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			ans += cnt[sum-deliciousness[i]]
		}
		cnt[deliciousness[i]]++
	}
	return ans % mod
}

func Test_countPairs(t *testing.T) {
	fmt.Println(countPairs([]int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}))
	fmt.Println(countPairs_([]int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}))
}

func Test_t(t *testing.T) {
	fmt.Println(int(4999950000) % int(1e9+7))
}

// leetcode lcp28: 采购方案
func purchasePlans(nums []int, target int) int {
	mod := int(1e9 + 7)
	n := len(nums)
	sort.Ints(nums)
	ans := 0
	for i := 0; i < n; i++ {
		left, right := i+1, n-1
		for left <= right {
			mid := left + (right-left)>>1
			if nums[i]+nums[mid] <= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		ans += left - i - 1
	}
	return ans % mod
}

func Test_purse(t *testing.T) {
	fmt.Println(purchasePlans([]int{2, 2, 1, 9}, 10))
}

// leetcode lcp29: 乐团站位
func orchestraLayout(num int, xPos int, yPos int) int {
	cycle := min(min(num-1-xPos, xPos), min(num-1-yPos, yPos))
	sum := (num - cycle) * cycle * 4
	cycleStart := (sum + 1) % 9
	ans := 0
	if xPos <= yPos {
		ans = (cycleStart + xPos + yPos - cycle<<1) % 9
	} else {
		ans = (cycleStart + (num-2*cycle)*4 - 4 - (xPos + yPos - cycle<<1)) % 9
	}
	if ans == 0 {
		return 9
	}
	return ans
}

func Test_orchestraLayout(t *testing.T) {
	fmt.Println(orchestraLayout(10, 5, 6))
}

// leetcode414: 第三大的数
func thirdMax(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	cnt := 0
	for i := n - 1; i > 0; i-- {
		if nums[i] != nums[i-1] {
			cnt++
		}
		if cnt == 2 {
			return nums[i-1]
		}
	}
	return nums[n-1]
}

// 参考算法思想
func thirdMax_(nums []int) int {
	var one, two, three int64 = math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		n := int64(num)
		if n > one {
			n, one = one, n
		}
		if n < one && n > two {
			n, two = two, n
		}
		if n < two && n > three {
			n, three = three, n
		}
	}
	if three != math.MinInt64 {
		return int(three)
	}
	return int(one)
}

func Test_third(t *testing.T) {
	fmt.Println(thirdMax([]int{2, 2, 1, 3}))
	fmt.Println(thirdMax_([]int{2, 2, 1, 3}))
}

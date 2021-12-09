package ojeveryday

import (
	"math"
	"sort"
	"strings"
)

// tag-[排序/二分查找]
// leetcode2070: 每一个查询的最大美丽值
func maximumBeauty(items [][]int, queries []int) []int {
	n := len(items)
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0] || (items[i][0] == items[j][0] && items[i][1] < items[j][1])
	})
	preMax := make([]int, n)
	preMax[0] = items[0][1]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], items[i][1])
	}
	out := make([]int, 0)
	for k := 0; k < len(queries); k++ {
		idx := sort.Search(n, func(i int) bool {
			return items[i][0] > queries[k]
		})
		if idx <= 0 {
			out = append(out, 0)
		} else {
			out = append(out, preMax[idx-1])
		}
	}
	return out
}

// tag-[动态规划]
// leetcode2063: 所有子字符串中的元音
// dp[i] = dp[i-1] + ？
// 如果第i位是元音，那么dp[i]需要在dp[i-1]的基础上，加上所有以i结尾的元音个数。
// 元音个数的计算采用累加的方式，如果发现第k个字符为元音，那么前0...k个数字将可以利用第K个元音，因此这个区间的所有位置能与字符尾组成的元音子串的个数需要加K+1。
func countVowels(word string) int64 {
	n := len(word)
	dp := make([]int64, n)
	vowel := []byte{'a', 'e', 'i', 'o', 'u'}
	cnt := int64(0)
	if isVowel(word[0], vowel) {
		dp[0] = 1
		cnt++
	}

	for i := 1; i < n; i++ {
		if isVowel(word[i], vowel) {
			cnt += int64(i + 1)
			dp[i] = dp[i-1] + cnt
		} else {
			dp[i] = dp[i-1] + cnt
		}
	}
	return dp[n-1]
}

func isVowel(b byte, vowel []byte) bool {
	for i := range vowel {
		if vowel[i] == b {
			return true
		}
	}
	return false
}

// tag-[数学]
// leetcode2063: 所有子字符串中的元音
// 解法二: 直接计算
func countVowels_(word string) int64 {
	ans := int64(0)
	n := len(word)
	for i := 0; i < n; i++ {
		if strings.ContainsRune("aeiou", rune(word[i])) {
			ans += int64((i + 1) * (n - i))
		}
	}
	return ans
}

// tag-[二分查找]
// leetcode2064: 分配给商店的最多商品的最小值
func minimizedMaximum(n int, quantities []int) int {
	maxn, sum := math.MinInt32, 0
	for i := range quantities {
		maxn = max(maxn, quantities[i])
		sum += quantities[i]
	}
	isOk := func(v int) bool {
		k := 0
		tmp := quantities[0]
		for i := 0; i < n; i++ {
			if tmp > v {
				tmp -= v
			} else {
				k++
				if k < len(quantities) {
					tmp = quantities[k]
				} else {
					return true
				}
			}
		}
		return false
	}
	l, r := (sum+n-1)/n, maxn
	for l < r {
		mid := (l + r) >> 1
		if isOk(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

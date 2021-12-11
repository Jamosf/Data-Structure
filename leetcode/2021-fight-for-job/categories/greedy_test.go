package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)
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

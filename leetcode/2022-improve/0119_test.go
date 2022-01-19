package _022_improve

import (
	"fmt"
	"testing"
)

// tag-[双指针]
// 每日一题
// leetcode219: 存在重复元素II
func containsNearbyDuplicate(nums []int, k int) bool {
	m, n := make(map[int]bool), len(nums)
	l, r := 0, 0
	for r < n{
		for len(m) < k{
			if _, ok := m[nums[r]]; ok{
				return true
			}
			m[nums[r]] = true
			r++
		}
		delete(m, nums[l])
		l++
	}
	return false
}

// tag-[哈希表]
// 每日一题
// leetcode219: 存在重复元素II
func containsNearbyDuplicate_(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums{
		if idx, ok := m[v]; ok && i-idx <= k{
			return true
		}else{
			m[v] = i
		}
	}
	return false
}

func Test_containsNearbyDuplicate(t *testing.T){
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1}, 3))
}
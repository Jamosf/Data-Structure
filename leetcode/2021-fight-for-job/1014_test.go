package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

func fourSum1(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				v := nums[i] + nums[j] + nums[l] + nums[r]
				if v > target {
					r--
				} else if v < target {
					l++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}
		}
	}
	return ans
}

func Test_fourSum(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

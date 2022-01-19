package _022_improve

import (
	"fmt"
	"testing"
)

// tag-[数组]
// leetcode747: 至少是其他数字两倍的最大数
func dominantIndex(nums []int) int {
	idx := 0
	for i := 1; i < len(nums); i++{
		if nums[i] > nums[idx]{
			idx = i
		}
	}
	for i := 0; i < len(nums); i++{
		if i != idx && nums[idx] < 2*nums[i]{
			return -1
		}
	}
	return idx
}

func Test_dominantIndex(t *testing.T){
	fmt.Println(dominantIndex([]int{3,6,1,0}))
}
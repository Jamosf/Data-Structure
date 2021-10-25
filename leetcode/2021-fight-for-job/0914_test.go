package ojeveryday

import (
	"container/heap"
	"fmt"
	"testing"
)

// lcp30
func magicTower_1(nums []int) int {
	sum := 1
	n := len(nums)
	h := &minHeap{}
	cnt := 0
	for i := 0; i < 2*n; i++ {
		if i >= len(nums) {
			break
		}
		v := nums[i]
		heap.Push(h, v)
		sum += v
		fmt.Println(sum, v)
		if sum <= 0 && i < len(nums)-1 {
			cnt++
			v := heap.Pop(h).(int)
			nums = append(nums, v)
			fmt.Println(nums)
			sum += -v
		}
		fmt.Println(sum)
	}
	if sum < 0 {
		return -1
	}
	return cnt
}

func Test_magicTower(t *testing.T) {
	fmt.Println(magicTower_1([]int{-1107, -19341, -36088, 27756, -73594, -4156, 12562, 50250, -93155}))
}
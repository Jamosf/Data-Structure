package leetcode

import (
	"container/heap"
	"fmt"
	"testing"
)

type minHeap []int

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	var v int
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// lcp30
func magicTower(nums []int) int {
	sum := 1
	n := len(nums)
	h := &minHeap{}
	cnt := 0
	for i := 0; i < 2*n; i++{
		if i >= len(nums){
			break
		}
		v := nums[i]
		heap.Push(h, v)
		sum += v
		fmt.Println(sum, v)
		if sum <= 0 && i < len(nums)-1{
			cnt++
			v := heap.Pop(h).(int)
			nums = append(nums, v)
			fmt.Println(nums)
			sum += -v
		}
		fmt.Println(sum)
	}
	if sum < 0{
		return -1
	}
	return cnt
}

func Test_magicTower(t *testing.T){
	fmt.Println(magicTower([]int{-1107,-19341,-36088,27756,-73594,-4156,12562,50250,-93155}))
}
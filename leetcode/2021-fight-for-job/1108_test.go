package ojeveryday

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

// tag-[前缀和]
// leetcode528: 前缀和
type Solution struct {
	preSum []int
}

func ConstructorSolution(w []int) Solution {
	s := Solution{preSum: make([]int, len(w))}
	s.preSum[0] = w[0]
	for i := 1; i < len(w); i++ {
		s.preSum[i] = s.preSum[i-1] + w[i]
	}
	return s
}

func (s *Solution) PickIndex() int {
	v := rand.Intn(s.preSum[len(s.preSum)-1]) + 1
	return sort.SearchInts(s.preSum, v)
}

func Test_Solution(t *testing.T) {
	s := ConstructorSolution([]int{1, 3})
	for i := 0; i < 100; i++ {
		fmt.Println(s.PickIndex())
	}
}

// tag-[前缀和]
// leetcode930: 前缀和
func numSubarraysWithSum930(nums []int, goal int) int {
	ans := 0
	sum1, sum2 := 0, 0
	left1, left2 := 0, 0
	for right, num := range nums {
		sum1 += num
		for left1 <= right && sum1 > goal {
			sum1 -= nums[left1]
			left1++
		}
		sum2 += num
		for left2 <= right && sum2 >= goal {
			sum2 -= nums[left2]
			left2++
		}
		ans += left2 - left1
	}
	return ans
}

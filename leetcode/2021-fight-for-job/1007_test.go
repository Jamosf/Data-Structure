package ojeveryday

import (
	"fmt"
	"testing"
)

// tag-[数组]
// 暴力
func numOfPairs(nums []string, target string) int {
	n := len(nums)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && nums[i]+nums[j] == target {
				fmt.Println(i, j)
				cnt++
			}
		}
	}
	return cnt
}

func Test_numOfPairs(t *testing.T) {
	fmt.Println(numOfPairs([]string{"777", "7", "77", "77"}, "7777"))
}

// tag-[双指针]
// 滑动窗口
func maxConsecutiveAnswers(answerKey string, k int) int {
	l, r, n := 0, 0, len(answerKey)
	sumt, sumf := 0, 0
	ans := 0
	for ; r < n; r++ {
		if answerKey[r] == 'T' {
			sumt++
		} else {
			sumf++
		}
		for sumt > k && sumf > k { // 关键：如果滑窗内的t和f都大于k，则需要收缩滑窗，并更新结果。
			if answerKey[l] == 'T' {
				sumt--
			} else {
				sumf--
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

// 与上一题相同的题目
func longestOnes(nums []int, k int) int {
	l, r, n := 0, 0, len(nums)
	sum1, sum0 := 0, 0
	ans := 0
	for ; r < n; r++ {
		if nums[r] == 1 {
			sum1++
		} else {
			sum0++
		}
		for sum0 > k { // 关键：如果滑窗内的t和f都大于k，则需要收缩滑窗，并更新结果。
			if nums[l] == 1 {
				sum1--
			} else {
				sum0--
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

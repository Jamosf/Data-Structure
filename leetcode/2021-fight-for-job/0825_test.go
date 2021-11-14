// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"strings"
	"testing"
)

// 第一题
// leetcode402: 移掉 K 位数字
// 从左到右，移除比右侧大的数，如果没有，则移除最后的数字
func removeKdigits(num string, k int) string {
	n := len(num)
	stack := make([]byte, 0, n)
	cnt := 0
	for i := 0; i < len(num); i++ {
		for len(stack) != 0 && stack[len(stack)-1] > num[i] {
			if cnt == k {
				break
			}
			stack = stack[:len(stack)-1]
			cnt++
		}
		stack = append(stack, num[i])
	}
	for cnt < k && len(stack) != 0 {
		stack = stack[:len(stack)-1]
		cnt++
	}
	if len(stack) == 0 || cnt != k {
		return "0"
	}
	ans := strings.TrimLeft(string(stack), "0")
	if len(ans) == 0 {
		return "0"
	}
	return ans
}

func Test_removeKdigits(t *testing.T) {
	fmt.Println(removeKdigits("100", 1))
}

// 第二题
// leetcode654：最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx, maxn := maxNumAndIdx(nums)
	root := &TreeNode{Val: maxn}
	root.Left = constructMaximumBinaryTree(nums[:idx])
	root.Right = constructMaximumBinaryTree(nums[idx+1:])
	return root
}

func maxNumAndIdx(nums []int) (int, int) {
	idx, maxn := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxn {
			maxn = nums[i]
			idx = i
		}
	}
	return idx, maxn
}

func Test_constructMaximumBinaryTree(t *testing.T) {
	r := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(r)
}

// 第三题
// TODO
func combinationSum1(candidates []int, target int) [][]int {
	res := make([][][]int, target+1)
	res[0] = make([][]int, 0)
	for i := 1; i < len(candidates); i++ { // 物品
		for j := candidates[i-1]; j <= target; j++ { // 背包
			for k := 0; k < len(res[j-candidates[i-1]]); k++ {
				res[j-candidates[i-1]][k] = append(res[j-candidates[i-1]][k], candidates[i])
			}
			res[j] = append(res[j], res[j-candidates[i-1]]...)
		}
	}
	return res[target]
}

func Test_combinationSum1(t *testing.T) {
	fmt.Println(combinationSum1([]int{2, 3, 6, 7}, 7))
}

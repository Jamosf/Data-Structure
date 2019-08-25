package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	var r TreeNode
	for i, num := range nums {

	}
}

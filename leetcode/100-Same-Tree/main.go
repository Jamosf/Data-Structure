package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val &&
			isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
}

/*
func mid_search(root *TreeNode, rootArray *[]int) {
	if root == nil {
		rootArray = nil
	}
	*rootArray = append(*rootArray, root.Val)
	if root.Left != nil {
		mid_search(root.Left, rootArray)
	}
	if root.Right != nil {
		mid_search(root.Right, rootArray)
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var pArray, qArray []int
	mid_search(p, &pArray)
	mid_search(q, &qArray)
    if pArray == nil || qArray == nil{
        return false
    }
    if len(pArray) != len(qArray) {
		return false
	}
	if len(pArray) == 0 {
		return true
	}
	for i := 0; i < len(pArray); i++ {
		if pArray[i] != qArray[i] {
			return false
		}
	}
	return true
}
*/

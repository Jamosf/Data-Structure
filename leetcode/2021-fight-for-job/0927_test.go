// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxn := 0
	var depth func(r *TreeNode) int
	depth = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		L := depth(r.Left)
		R := depth(r.Right)
		maxn = max(maxn, L+R-1)
		return 1 + max(L, R)
	}
	depth(root)
	return maxn
}

func groupAnagrams(strs []string) [][]string {
	mk := make(map[[26]int][]string)
	var cacl func(s string) [26]int
	cacl = func(s string) [26]int {
		ans := [26]int{}
		for i := range s {
			ans[s[i]-'a']++
		}
		return ans
	}
	for i := range strs {
		v := cacl(strs[i])
		mk[v] = append(mk[v], strs[i])
	}
	ans := make([][]string, 0)
	for _, v := range mk {
		ans = append(ans, v)
	}
	return ans
}

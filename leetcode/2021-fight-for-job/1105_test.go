// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// leetcode89:格雷编码公式i^(i>>1)
func grayCode(n int) []int {
	size := 1 << n
	ans := make([]int, 0, size)
	for i := 0; i < size; i++ {
		ans = append(ans, i^(i>>1))
	}
	return ans
}

func Test_grayCode(t *testing.T) {
	print_binary_array(grayCode(3), 3)
}

// leetcode318:位运算
func maxProduct318(words []string) int {
	n := len(words)
	count := make([]int, n)
	for i := range words {
		for j := 0; j < len(words[i]); j++ {
			count[i] |= 1 << (words[i][j] - 'a')
		}
	}
	maxn := 0
	for i := range words {
		for j := i + 1; j < n; j++ {
			if count[i]&count[j] == 0 {
				maxn = max(maxn, len(words[i])*len(words[j]))
			}
		}
	}
	return maxn
}

func Test_maxProduct318(t *testing.T) {
	fmt.Println(maxProduct318([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	fmt.Println(maxProduct318([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	fmt.Println(maxProduct318([]string{"a", "aa", "aaa", "aaaa"}))
}

// leetcode371:两数之和位运算
func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a ^= b
		b = carry
	}
	return a
}

func Test_getSum(t *testing.T) {
	fmt.Println(getSum(1, 3))
}

// leetcode421:异或交换律，a^b=c, a^c=b, b^c=a
// hash解法
func findMaximumXOR(nums []int) int {
	mask := 0
	res := 0
	for i := 31; i >= 0; i-- {
		mask |= 1 << i
		m := make(map[int]struct{})
		for i := range nums {
			m[mask&nums[i]] = struct{}{}
		}
		tmp := res | (1 << i)
		for k := range m {
			if _, ok := m[k^tmp]; ok {
				res = tmp
				break
			}
		}
	}
	return res
}

// leetcode421:字典树解法
const highBit = 30

type trie_ struct {
	left, right *trie_
}

func (t *trie_) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &trie_{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &trie_{}
			}
			cur = cur.right
		}
	}
}

func (t *trie_) check(num int) (x int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			// a_i 的第 k 个二进制位为 0，应当往表示 1 的子节点 right 走
			if cur.right != nil {
				cur = cur.right
				x = x*2 + 1
			} else {
				cur = cur.left
				x = x * 2
			}
		} else {
			// a_i 的第 k 个二进制位为 1，应当往表示 0 的子节点 left 走
			if cur.left != nil {
				cur = cur.left
				x = x*2 + 1
			} else {
				cur = cur.right
				x = x * 2
			}
		}
	}
	return
}

func findMaximumXOR_(nums []int) (x int) {
	root := &trie_{}
	for i := 1; i < len(nums); i++ {
		// 将 nums[i-1] 放入字典树，此时 nums[0 .. i-1] 都在字典树中
		root.add(nums[i-1])
		// 将 nums[i] 看作 ai，找出最大的 x 更新答案
		x = max(x, root.check(nums[i]))
	}
	return
}

func Test_findMaximumXOR(t *testing.T) {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}

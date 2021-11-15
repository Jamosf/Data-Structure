// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"leetcode/leetcode/2021-fight-for-job/basic_algo"
	"strings"
	"testing"
)

// 基础数据结构
type maxHeap struct {
	basic_algo.MaxHeap
}

type maxHeap1 struct {
	basic_algo.MaxHeapByIntSlice
}

type minHeap struct {
	basic_algo.MinHeap
}

type minHeapPair struct {
	basic_algo.MinHeapPair
}

type unionFind struct {
	*basic_algo.UnionFind
}

func newUnionFind(size int) *unionFind {
	return &unionFind{basic_algo.NewUnionFind(size)}
}

func (u *unionFind) union(p, q int) {
	u.Union(p, q)
}

func (u *unionFind) find(p int) int {
	return u.Find(p)
}

func (u *unionFind) isConnected(p, q int) bool {
	return u.IsConnected(p, q)
}

func (u *unionFind) count() int {
	return u.Count()
}

type trie struct {
	next    [maxNodeNum]*trie
	num     int
	endFlag bool
}

func (t *trie) insert(word string) {
	p := t
	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i] - 'a'
		if p.next[c] == nil {
			p.next[c] = &trie{}
			p.num++
			p = p.next[c]
		} else {
			p = p.next[c]
		}
		if i == l-1 {
			p.endFlag = true
		}
	}
}

func (t *trie) search(word string) bool {
	p := t
	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i] - 'a'
		if p.next[c] == nil {
			return false
		}
		p = p.next[c]
		if p.endFlag && i == l-1 {
			return true
		}
	}
	return false
}

func (t *trie) searchStartsWith(prefix string) bool {
	p := t
	l := len(prefix)
	for i := 0; i < l; i++ {
		c := prefix[i] - 'a'
		if p.next[c] == nil {
			return false
		}
		p = p.next[c]
	}
	return true
}

// 基础函数
var (
	max = basic_algo.Max
	min = basic_algo.Min

	abs        = basic_algo.Abs
	minusAbs   = basic_algo.MinusAbs
	quickSortK = basic_algo.QuickSortK
)

var data = "[[2,3],[4,5],[6,7],[8,9],[1,10]]"

func Test_convertInputData(t *testing.T) {
	fmt.Println(strings.ReplaceAll(strings.ReplaceAll(data, "[", "{"), "]", "}"))
}

func print_matrix(grid [][]int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v\t", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func print_matrix_b(grid [][]byte) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s\t", string(grid[i][j]))
		}
		fmt.Printf("\n")
	}
}

func newListNode(nums []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for i := range nums {
		p.Next = &ListNode{Val: nums[i]}
		p = p.Next
	}
	return dummy.Next
}

func print_byte(b []byte) {
	for i := range b {
		fmt.Printf("%c ", b[i])
	}
	fmt.Println()
}

func print_binary_array(b []int, n int) {
	for i := range b {
		print_binary(b[i], n)
	}
}

func print_binary(b int, n int) {
	fmt.Printf("%0*b ", n, b)
}

func Test_print_binary(t *testing.T) {
	print_binary(3, 4)
}

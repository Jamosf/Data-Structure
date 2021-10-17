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

// 基础函数
var (
	max = basic_algo.Max
	min = basic_algo.Min

	abs      = basic_algo.Abs
	minusAbs = basic_algo.MinusAbs
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

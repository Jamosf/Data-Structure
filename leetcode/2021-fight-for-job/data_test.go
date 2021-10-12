// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"strings"
	"testing"
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

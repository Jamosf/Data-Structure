package ojeveryday

import (
	"fmt"
	"testing"
)

// 第一题
// leetcode79: 单词搜索
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	pos := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	find := false
	var dfs func(i, j int, idx int)
	dfs = func(i, j int, idx int) {
		if find || board[i][j] != word[idx] {
			return
		}
		if idx == len(word)-1 {
			find = true
			return
		}
		visited[i][j] = true
		for k := 0; k < 4; k++ {
			x, y := i+pos[k][0], j+pos[k][1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if !visited[x][y] {
					dfs(x, y, idx+1)
				}
			}
		}
		visited[i][j] = false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				for i := 0; i < m; i++ {
					visited[i] = make([]bool, n)
				}
				dfs(i, j, 0)
			}
		}
	}
	return find
}

func Test_exist(t *testing.T) {
	fmt.Println(exist([][]byte{{'a'}, {'a'}}, "aaa"))
}

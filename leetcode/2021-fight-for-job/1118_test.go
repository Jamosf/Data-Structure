// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

// tag-[广度优先搜索/图]
// leetcode529: 扫雷游戏
func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	pos := [8][2]int{{1, 0}, {1, 1}, {1, -1}, {0, 1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{click[0], click[1]})
	visited[click[0]][click[1]] = true
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		e := board[v[0]][v[1]]
		if e >= '1' && e <= '9' {
			continue
		}
		cnt := 0
		for k := 0; k < len(pos); k++ {
			if i, j := v[0]+pos[k][0], v[1]+pos[k][1]; i < m && i >= 0 && j < n && j >= 0 && board[i][j] == 'M' {
				cnt++
			}
		}
		if board[v[0]][v[1]] == 'M' {
			continue
		}
		if cnt > 0 {
			board[v[0]][v[1]] = byte('0' + cnt)
			continue
		} else {
			board[v[0]][v[1]] = 'B'
		}
		for k := 0; k < len(pos); k++ {
			if i, j := v[0]+pos[k][0], v[1]+pos[k][1]; i < m && i >= 0 && j < n && j >= 0 {
				if !visited[i][j] {
					queue = append(queue, [2]int{i, j})
					visited[i][j] = true
				}
			}
		}
	}
	return board
}

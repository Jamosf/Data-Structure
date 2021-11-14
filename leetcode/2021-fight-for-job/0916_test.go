// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

var dx = [4]int{1, 1, 0, -1}
var dy = [4]int{0, 1, 1, 1}

func f(g [][]byte) int {
	n, m := len(g), len(g[0])
	var check func(i, j int) bool
	check = func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m
	}
	ans := 0
	for {
		done := true
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				for k := 0; k < 4; k++ {
					if g[i][j] == 'O' {
						x, y := 0, 0
						for check(i+dx[k]*x, j+dy[k]*x) && g[i+dx[k]*x][j+dy[k]*x] == 'O' {
							x++
						}
						for check(i+dx[k]*y, j+dy[k]*y) && g[i+dx[k]*y][j+dy[k]*y] == 'O' {
							y--
						}
						if check(i+dx[k]*x, j+dy[k]*x) && g[i+dx[k]*x][j+dy[k]*x] == 'X' &&
							check(i+dx[k]*y, j+dy[k]*y) && g[i+dx[k]*y][j+dy[k]*y] == 'X' {
							ans++
							done = false
							g[i][j] = 'X'
						}
					}
				}
			}
		}
		if done {
			break
		}
	}
	return ans
}

// leetcode LCP41: 黑白翻转棋
func flipChess(chessboard []string) int {
	n, m := len(chessboard), len(chessboard[0])
	g := make([][]byte, n)
	for i := range g {
		g[i] = make([]byte, m)
	}
	var init func()
	init = func() {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				g[i][j] = chessboard[i][j]
			}
		}
	}

	ans := -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			init()
			if g[i][j] == '.' {
				g[i][j] = 'X'
				ans = max(ans, f(g))
				g[i][j] = '.'
			}
		}
	}
	return ans
}

func Test_flipChesss(t *testing.T) {
	fmt.Println(flipChess([]string{".X.", ".O.", "XO."}))
}

// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

func minimumSwitchingTimes(source [][]int, target [][]int) int {
	m, n := len(source), len(source[0])
	mk := map[int]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mk[source[i][j]]++
			mk[target[i][j]]--
		}
	}
	cnt := 0
	for _, v := range mk {
		if v > 0 {
			cnt += v
		} else {
			cnt += -v
		}
	}
	return cnt >> 1
}

func Test_minimumSwitchingTimes(t *testing.T) {
	fmt.Println(minimumSwitchingTimes([][]int{{1, 3}, {5, 4}}, [][]int{{3, 1}, {6, 5}}))
	fmt.Println(maxmiumScore([]int{1, 2, 8, 9}, 3))
}

func maxmiumScore(cards []int, cnt int) int {
	n := len(cards)
	maxn := 0
	sum := 0
	visited := make([]bool, n)
	var dfs func(level int)
	dfs = func(level int) {
		if level == cnt {
			if sum&1 == 0 {
				maxn = max(maxn, sum)
			}
			return
		}
		for i := 0; i < n; i++ {
			if !visited[i] {
				sum += cards[i]
				visited[i] = true
				dfs(level + 1)
				sum -= cards[i]
				visited[i] = false
			}
		}
	}
	dfs(0)
	return maxn
}

func flipChess1(chessboard []string) int {
	m, n := len(chessboard), len(chessboard[0])
	direction := [8][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	maxn := 0
	ans := make([][]int, 0)
	var check func(i, j, dx, dy int) int
	check = func(x, y, dx, dy int) int {
		x, y = x+dx, y+dy
		step := 1
		for x >= 0 && x < m && y >= 0 && y < n {
			tmp := make([]int, 0)
			if step == 1 {
				if chessboard[x][y] == '.' || chessboard[x][y] == 'X' {
					return 0
				}
			} else {
				if chessboard[x][y] == '.' {
					return 0
				}
				if chessboard[x][y] == 'X' {
					ans = append(ans, tmp)
					return step - 1
				}
			}
			step++
			x += dx
			y += dy
			tmp = append(tmp, []int{x, y}...)
		}
		// for i := range ans{
		// 	// for i
		// }
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if chessboard[i][j] != '.' {
				continue
			}
			count := 0
			for k := 0; k < 8; k++ {
				count += check(i, j, direction[k][0], direction[k][1])
			}
			maxn = max(maxn, count)
		}
	}

	return maxn
}

func Test_flipChess(t *testing.T) {
	fmt.Println(flipChess([]string{".X.", ".O.", "XO."}))
}

// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
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

func circleGame(toys [][]int, circles [][]int, r int) int {
	var check func(toy []int, circle []int, r int) bool
	check = func(toy []int, circle []int, r int) bool {
		if toy[0]+toy[2] > circle[0]+r || toy[0]-toy[2] < circle[0]-r {
			return false
		}
		if toy[1]+toy[2] > circle[1]+r || toy[1]-toy[2] < circle[1]-r {
			return false
		}
		return true
	}
	ans := 0
	for i := 0; i < len(toys); i++ {
		for j := 0; j < len(circles); j++ {
			if check(toys[i], circles[j], r) {
				ans++
				break
			}
		}
	}
	return ans
}

// 利用二分找到离玩具圆心最近的圈的中心，如果离的最近的圈都套不上，那么离的远的肯定更加套不上。
func circleGame1(toys [][]int, circles [][]int, r0 int) (ans int) {
	sort.Slice(circles, func(i, j int) bool { a, b := circles[i], circles[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })

	// 将横坐标相同的圈分为一组
	type pair struct {
		x  int
		ys []int
	}
	a, y := []pair{}, -1
	for _, p := range circles {
		if len(a) == 0 || p[0] > a[len(a)-1].x {
			a = append(a, pair{p[0], []int{p[1]}})
			y = -1
		} else if p[1] > y { // 去重
			a[len(a)-1].ys = append(a[len(a)-1].ys, p[1])
			y = p[1]
		}
	}

	for _, t := range toys {
		x, y, r := t[0], t[1], t[2]
		if r > r0 {
			continue
		}
		i := sort.Search(len(a), func(i int) bool { return a[i].x+r0 >= x+r })
		for ; i < len(a) && a[i].x-r0 <= x-r; i++ {
			cx, ys := a[i].x, a[i].ys
			j := sort.SearchInts(ys, y)
			if j < len(ys) {
				if cy := ys[j]; (x-cx)*(x-cx)+(y-cy)*(y-cy) <= (r0-r)*(r0-r) {
					ans++
					break
				}
			}
			if j > 0 {
				if cy := ys[j-1]; (x-cx)*(x-cx)+(y-cy)*(y-cy) <= (r0-r)*(r0-r) {
					ans++
					break
				}
			}
		}
	}
	return
}

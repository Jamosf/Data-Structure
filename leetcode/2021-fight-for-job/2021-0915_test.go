package _021_fight_for_job

import (
	"fmt"
	"sort"
	"testing"
)

func storeWater(bucket []int, vat []int) int {
	maxn := vat[0]
	for i := range vat {
		maxn = max(maxn, vat[i])
	}
	if maxn == 0 {
		return 0
	}
	ans := 10001
	for i := 1; i < 10000; i++ {
		if i > ans {
			break
		}
		cur := 0
		for j := range vat {
			v := vat[j]/i - bucket[j]
			if vat[j]%i != 0 {
				v++
			}
			if v > 0 {
				cur += v
			}
			if cur >= ans {
				break
			}
		}
		ans = min(ans, cur+i)
	}
	return ans
}

func Test_storeWater(t *testing.T) {
	fmt.Println(storeWater([]int{1, 3}, []int{6, 8}))
}

func maxmiumScore(a []int, cnt int) int {
	n := len(a)
	sort.Ints(a)
	record := [2]int{-1, -1}
	sum := 0
	for i := n - cnt; i < n; i++ {
		sum += a[i]
		if a[i]&1 == 1 {
			if record[1] == -1 {
				record[1] = a[i]
			}
		} else {
			if record[0] == -1 {
				record[0] = a[i]
			}
		}
	}

	if sum&1 == 0 {
		return sum
	}

	for i := n - cnt - 1; i >= 0; i-- {
		if a[i]&1 == 1 {
			if record[0] != -1 {
				return sum - record[0] + a[i]
			}
		} else {
			if record[1] != -1 {
				return sum - record[1] + a[i]
			}
		}
	}
	return 0
}

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
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			g[i][j] = chessboard[i][j]
		}
	}

	ans := -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g[i][j] == '.' {
				g[i][j] = 'X'
				ans = max(ans, f(g))
				g[i][j] = '.'
			}
		}
	}
	return ans
}

func Test_flipChess(t *testing.T) {
	fmt.Println(flipChess([]string{"....X.", "....X.", "XOOO..", "......", "......"}))
}

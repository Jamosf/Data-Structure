// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

// 第一题
func solveNQueens(n int) [][]string {
	var ans [][]string
	tmp := make([][]byte, n)
	for i := 0; i < n; i++ {
		tmp[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			tmp[i][j] = '.'
		}
	}
	col := make([]bool, n)
	left := make([]bool, 2*n)
	right := make([]bool, 2*n)
	var backtracking func(level int)
	backtracking = func(level int) {
		if level == n {
			t := make([]string, 0, n)
			for i := range tmp {
				t = append(t, string(tmp[i]))
			}
			ans = append(ans, t)
			return
		}
		i := level
		for j := 0; j < n; j++ {
			if !col[j] && !left[i+j] && !right[n-i+j] {
				col[j] = true
				left[i+j] = true
				right[n-i+j] = true
				tmp[i][j] = 'Q'
				backtracking(level + 1)
				tmp[i][j] = '.'
				col[j] = false
				left[i+j] = false
				right[n-i+j] = false
			}
		}
	}
	backtracking(0)
	return ans
}

func Test_solveNQueens(t *testing.T) {
	fmt.Println(solveNQueens(4))
}

// 求最小翻转的个数，就是求两个岛之间的最短距离
func shortestBridge(grid [][]int) int {
	queue := list.New()
	m, n := len(grid), len(grid[0])
	direction := [4][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i == m || j == n || grid[i][j] == 2 {
			return
		}
		if grid[i][j] == 0 {
			queue.PushBack(pos{i, j})
			return
		}
		grid[i][j] = 2
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	flap := false
	for i := 0; i < m; i++ {
		if flap {
			break
		}
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				flap = true
				break
			}
		}
	}
	level := 0
	for queue.Len() != 0 {
		level++
		size := queue.Len()
		for i := 0; i < size; i++ {
			v := queue.Front()
			queue.Remove(v)
			p := v.Value.(pos)
			for k := 0; k < 4; k++ {
				x, y := p.x+direction[k][0], p.y+direction[k][1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if grid[x][y] == 1 {
						return level
					}
					if grid[x][y] == 2 {
						continue
					}
					grid[x][y] = 2
					queue.PushBack(pos{x, y})
				}
			}
		}
	}
	return 0
}

type graph struct {
	vertex []string
	edges  [][]int
	n, e   int
}

//
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	isExist, idx := isWordExist(beginWord, wordList)
	newWordList := []string{beginWord}
	if !isExist {
		newWordList = append(newWordList, wordList...)
		idx = 0
	}
	isEndExist, endIdx := isWordExist(endWord, newWordList)
	if !isEndExist {
		return nil
	}
	g := graph{n: len(newWordList), vertex: newWordList}
	edges := make([][]int, g.n)
	for i := range edges {
		edges[i] = make([]int, g.n)
	}
	g.edges = edges
	for i := 0; i < g.n; i++ {
		for j := i + 1; j < g.n; j++ {
			if isOneCharDiff(g.vertex[i], g.vertex[j]) {
				g.edges[i][j] = 1
				g.edges[j][i] = 1
			}
		}
	}
	visited := make([]bool, g.n)
	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	queue := list.New()
	queue.PushBack(idx)
	visited[idx] = true
	dist[idx] = 0
	for queue.Len() != 0 {
		v := queue.Front()
		queue.Remove(v)
		index := v.Value.(int)
		for j := 0; j < g.n; j++ {
			if !visited[j] && g.edges[index][j] == 1 {
				visited[j] = true
				queue.PushBack(j)
				if dist[j] > dist[index]+1 {
					dist[j] = dist[index] + 1
				}
			}
		}
	}
	totalLen := dist[endIdx] + 1
	ret := make([][]int, totalLen)
	for i := range dist {
		if dist[i] < totalLen {
			ret[dist[i]] = append(ret[dist[i]], i)
		}
	}
	var ans [][]string
	var tmp []string

	ans = append(ans, tmp)
	return ans
}

func isWordExist(beginWord string, wordList []string) (bool, int) {
	for i := range wordList {
		if beginWord == wordList[i] {
			return true, i
		}
	}
	return false, -1
}

func isOneCharDiff(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	cnt := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			cnt++
		}
	}
	return cnt == 1
}

func Test_findLadders(t *testing.T) {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	isExist, idx := isWordExist(beginWord, wordList)
	if !isExist {
		wordList = append(wordList, beginWord)
		idx = len(wordList) - 1
	}
	isEndExist, endIdx := isWordExist(endWord, wordList)
	if !isEndExist {
		return 0
	}
	g := graph{n: len(wordList), vertex: wordList}
	edges := make([][]int, g.n)
	for i := range edges {
		edges[i] = make([]int, g.n)
	}
	g.edges = edges
	for i := 0; i < g.n; i++ {
		for j := i + 1; j < g.n; j++ {
			if isOneCharDiff(g.vertex[i], g.vertex[j]) {
				g.edges[i][j] = 1
				g.edges[j][i] = 1
			}
		}
	}
	visited := make([]bool, g.n)
	dist := make([]int, g.n)
	queue := list.New()
	queue.PushBack(idx)
	visited[idx] = true
	dist[idx] = 1
	for queue.Len() != 0 {
		v := queue.Front()
		queue.Remove(v)
		index := v.Value.(int)
		if index == endIdx {
			return dist[index]
		}
		for j := 0; j < g.n; j++ {
			if !visited[j] && g.edges[index][j] == 1 {
				visited[j] = true
				queue.PushBack(j)
				dist[j] = dist[index] + 1
			}
		}
	}
	return 0
}

func Test_ladderLength1(t *testing.T) {
	fmt.Println(ladderLength("lost", "miss", []string{"most", "mist", "miss", "lost", "fist", "fish"}))
}

func getCandidates(word string) []string {
	var res []string
	for j := 0; j < len(word); j++ {
		for i := 0; i < 26; i++ {
			if word[j] != byte(int('a')+i) {
				res = append(res, word[:j]+string(rune(int('a')+i))+word[j+1:])
			}
		}
	}
	return res
}

func Test_getCandidates(t *testing.T) {
	fmt.Println(getCandidates("mist"))
}

package days

import (
	"fmt"
	"strconv"
	"testing"
)

// tag-[深度优先搜索]
// leetcode1036: 逃离大迷宫
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	b := make(map[string]bool)
	for _, block := range blocked {
		b[strconv.Itoa(block[0])+"->"+strconv.Itoa(block[1])] = true
	}
	inf := int(1e6)
	var dfs func(x, y int, s, t []int, vis map[string]bool) bool
	dfs = func(x, y int, s, t []int, vis map[string]bool) bool {
		if x >= inf || x < 0 || y >= inf || y < 0 {
			return false
		}
		key := strconv.Itoa(x) + "->" + strconv.Itoa(y)
		if vis[key] || b[key] {
			return false
		}
		if (x == t[0] && y == t[1]) || abs(s[0]-x)+abs(s[1]-y) > 200 { // 曼哈顿距离大于200可以绕着走
			return true
		}
		vis[key] = true
		return dfs(x+1, y, s, t, vis) || dfs(x-1, y, s, t, vis) || dfs(x, y+1, s, t, vis) || dfs(x, y-1, s, t, vis)
	}
	return dfs(source[0], source[1], source, target, map[string]bool{}) && dfs(target[0], target[1], target, source, map[string]bool{})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Test_e(t *testing.T) {
	fmt.Println(int(1e6))
}

// tag-[广度优先搜索]
// leetcode1036: 逃离大迷宫
// 思路：计算障碍物能包围的最大面积，如果从source和target能遍历的值大于最大面积，则说明一定没有被包围，则可以连通
func isEscapePossible_(blocked [][]int, source []int, target []int) bool {
	n := len(blocked)
	max := n * (n - 1) / 2
	b := make(map[string]bool)
	for _, block := range blocked {
		b[strconv.Itoa(block[0])+"->"+strconv.Itoa(block[1])] = true
	}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	bfs := func(s, t []int) bool {
		vis := make(map[string]bool)
		queue := make([][2]int, 0, max)
		queue = append(queue, [2]int{s[0], s[1]})
		vis[strconv.Itoa(s[0])+"->"+strconv.Itoa(s[1])] = true
		cnt := 0
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			cnt++
			if (v[0] == t[0] && v[1] == t[1]) || cnt > max {
				return true
			}
			for i := 0; i < 4; i++ {
				x, y := v[0]+dirs[i][0], v[1]+dirs[i][1]
				key := strconv.Itoa(x) + "->" + strconv.Itoa(y)
				if x >= 0 && x < 1e6 && y >= 0 && y < 1e6 && !vis[key] && !b[key] {
					queue = append(queue, [2]int{x, y})
					vis[key] = true
				}
			}
		}
		return false
	}
	return bfs(source, target) && bfs(target, source)
}

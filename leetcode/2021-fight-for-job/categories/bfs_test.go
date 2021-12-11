package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)
// tag-[广度优先搜索]
// 第一题
// leetcode695: 岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	visited = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	l := list.New()
	maxn := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 && !visited[i][j] {
				l.PushBack(pos{i, j})
				visited[i][j] = true
				maxn = max(maxn, bfs(grid, l))
			}
		}
	}
	return maxn
}

func bfs(grid [][]int, l *list.List) int {
	n, m := len(grid), len(grid[0])
	cnt := 1
	for l.Len() != 0 {
		v := l.Front()
		l.Remove(v)
		vv := v.Value.(pos)
		cnt++
		for i := 0; i < len(direction); i++ {
			x1, y1 := vv.x+direction[i][0], vv.y+direction[i][1]
			if x1 >= 0 && x1 < n && y1 >= 0 && y1 < m && !visited[x1][y1] && grid[x1][y1] != 0 {
				l.PushBack(pos{x1, y1})
				visited[x1][y1] = true
			}
		}
	}
	return cnt
}

// tag-[深度预先搜索]
// 第二题
// leetcode733: 图像渲染
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 {
		return nil
	}
	k := image[sr][sc]
	if k == newColor {
		return image
	}
	dfs2(image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs2(image [][]int, r, c int, k int, color int) {
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) {
		return
	}
	if image[r][c] != k {
		return
	}
	image[r][c] = color
	dfs2(image, r+1, c, k, color)
	dfs2(image, r, c+1, k, color)
	dfs2(image, r-1, c, k, color)
	dfs2(image, r, c-1, k, color)
}

func Test_floodFill(t *testing.T) {
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1))
}
// tag-[广度优先搜索]
// leetcode1992: 找到所有的农场组
// 注意：使用切片作为队列时，长度申请不合理，会超时。使用list不会超时。
func findFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])
	diret := [2][2]int{{1, 0}, {0, 1}}
	var bfs func(i, j int) []int
	bfs = func(i, j int) []int {
		l := list.New()
		l.PushBack([]int{i, j})
		land[i][j] = 0
		maxx, maxy := i, j
		for l.Len() != 0 {
			v := l.Front()
			l.Remove(v)
			vv := v.Value.([]int)
			for p := 0; p < 2; p++ {
				x, y := vv[0]+diret[p][0], vv[1]+diret[p][1]
				if x >= 0 && x < m && y >= 0 && y < n && land[x][y] == 1 {
					l.PushBack([]int{x, y})
					land[x][y] = 0
					maxx = max(maxx, x)
					maxy = max(maxy, y)
				}
			}
		}
		return []int{i, j, maxx, maxy}
	}
	var ans [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] != 0 {
				ans = append(ans, bfs(i, j))
			}
		}
	}
	return ans
}

func Test_findFarmland(t *testing.T) {
	fmt.Println(findFarmland([][]int{{1, 1}, {1, 1}}))
}// tag-[广度优先搜索]
// leetcode841:钥匙和房间 bfs
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	q := list.New()
	q.PushBack(0)
	visited[0] = true
	for q.Len() != 0 {
		v := q.Front()
		q.Remove(v)
		vv := v.Value.(int)
		for _, t := range rooms[vv] {
			if !visited[t] {
				visited[t] = true
				q.PushBack(t)
			}
		}
	}
	for i := range visited {
		if !visited[i] {
			return false
		}
	}
	return true
}
// tag-[广度优先搜索]
type pair struct {
	node *TreeNode
	left int
}

// leetcode113:广度优先搜索解法,记录父节点，重组路径。
func pathSum_(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return
	}
	parent := map[*TreeNode]*TreeNode{}
	getPath := func(node *TreeNode) (path []int) {
		for ; node != nil; node = parent[node] {
			path = append(path, node.Val)
		}
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
		return
	}
	queue := []pair{{root, targetSum}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		node := p.node
		left := p.left - node.Val
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				ans = append(ans, getPath(node))
			}
		} else {
			if node.Left != nil {
				parent[node.Left] = node
				queue = append(queue, pair{node.Left, left})
			}
			if node.Right != nil {
				parent[node.Right] = node
				queue = append(queue, pair{node.Right, left})
			}
		}
	}
	return
}

func pathSumCnt(root *TreeNode, targetSum int) int {
	sum := 0
	var dfs func(r *TreeNode, v int)
	dfs = func(r *TreeNode, v int) {
		if r == nil {
			return
		}
		v -= r.Val
		if v == 0 {
			sum++
		}
		dfs(r.Left, v)
		dfs(r.Right, v)
	}
	dfs(root, targetSum)
	return sum
}
// tag-[广度优先搜索]
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
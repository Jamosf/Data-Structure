package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)
// tag-[深度优先搜索]
// 第二题
func updateMatrix1(mat [][]int) [][]int {
	var dfs func(mat [][]int, r, c int) int
	dfs = func(mat [][]int, r, c int) int {
		if r < 0 || r >= len(mat) || c < 0 || c >= len(mat[0]) {
			return 0
		}
		var ret int
		if mat[r][c] == 0 {
			ret++
			return ret
		}
		ret = min(ret, dfs(mat, r-1, c))
		ret = min(ret, dfs(mat, r, c+1))
		ret = min(ret, dfs(mat, r+1, c))
		ret = min(ret, dfs(mat, r, c-1))
		return ret
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != 0 {
				mat[i][j] = dfs(mat, j, j)
			}
		}
	}
	return mat
}
// tag-[深度优先搜索]
// leetcode lcp39: 无人机方阵
func escapeMaze(g [][]string) bool {
	k, m, n := len(g), len(g[0]), len(g[0][0])
	dir := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	vis := make([][][][6]bool, k)
	for i := range vis {
		vis[i] = make([][][6]bool, m)
		for j := range vis[i] {
			vis[i][j] = make([][6]bool, n)
		}
	}
	var dfs func(t, i, j, s int) bool
	// t表示当前走到第几步，x,y表示当前的位置，s表示是否使用了消除术
	// s由三位组成，最低位表示是否使用临时消除术，高两位表示是否使用了永久消除术（10 为已经使用永久消除、01为当前处于永久消除位置、00为未使用永久消除术）
	dfs = func(t, x, y, s int) bool {
		if x < 0 || x >= m || y < 0 || y >= n || m-1-x+n-1-y > k-t || vis[t][x][y][s] {
			return false
		}
		if x == m-1 && y == n-1 {
			return true
		}
		vis[t][x][y][s] = true
		// 先排查清除术情况, 如果当前处于永久清除位置
		if s>>1 == 1 {
			for _, d := range dir {
				if dfs(t+1, x+d[0], y+d[1], s^6) { // 标记为已使用
					return true
				}
			}
			// 四周走不通，则留在原地
			return dfs(t+1, x, y, s)
		}
		// 尝试使用永久清除
		if s>>1 == 0 && g[t][x][y] == '#' && dfs(t, x, y, s|2) {
			return true
		}
		// 尝试使用临时清除
		if g[t][x][y] == '#' {
			if s&1 == 1 {
				return false
			}
			s |= 1
		}
		for _, d := range dir {
			if dfs(t+1, x+d[0], y+d[1], s) { // 标记为已使用
				return true
			}
		}
		return dfs(t+1, x, y, s)
	}
	return dfs(0, 0, 0, 0)
}
// tag-[深度优先搜索/图]
// leetcode797:图中所有可能的路径
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	edge := make([][]int, n)
	for i := range edge {
		edge[i] = make([]int, n)
	}
	for i := range graph {
		for j := range graph[i] {
			edge[i][graph[i][j]] = 1
		}
	}
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n-1 {
			ans = append(ans, append([]int{}, tmp...))
			return
		}
		for j := 0; j < n; j++ {
			if edge[i][j] == 1 {
				tmp = append(tmp, j)
				dfs(j)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	tmp = append(tmp, 0)
	dfs(0)
	return ans
}

// tag-[排序]
// leetcode802:最终的安全位置
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	rg := make([][]int, n)
	indegree := make([]int, n)
	for i := range graph {
		for _, v := range graph[i] {
			rg[v] = append(rg[v], i)
			indegree[i]++
		}
	}
	q := list.New()
	for i := range indegree {
		if indegree[i] == 0 {
			q.PushBack(i)
		}
	}
	ans := make([]int, 0)
	for q.Len() != 0 {
		v := q.Front()
		q.Remove(v)
		vv := v.Value.(int)
		ans = append(ans, vv)
		for _, t := range rg[vv] {
			indegree[t]--
			if indegree[t] == 0 {
				q.PushBack(t)
			}
		}
	}
	sort.Ints(ans)
	return ans
}

func Test_eventualSafeNodes(t *testing.T) {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
	fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
}

// tag-[排序]
// leetcode802:三色标记解法，参考
func eventualSafeNodes__(graph [][]int) (ans []int) {
	n := len(graph)
	color := make([]int, n)
	var safe func(int) bool
	safe = func(x int) bool {
		if color[x] > 0 {
			return color[x] == 2
		}
		color[x] = 1
		for _, y := range graph[x] {
			if !safe(y) {
				return false
			}
		}
		color[x] = 2
		return true
	}
	for i := 0; i < n; i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return
}

const (
	NotExplored = 0
	Explored    = 1
	Safe        = 2
)

// leetcode802:三色标记解法，参考
func eventualSafeNodes_(graph [][]int) []int {
	state := make([]int, len(graph))
	var res []int
	for v := range graph {
		if checkSafe(graph, state, v) {
			res = append(res, v)
		}
	}

	return res
}

func checkSafe(graph [][]int, state []int, v int) bool {
	switch state[v] {
	case NotExplored:
		state[v] = Explored
		for _, n := range graph[v] {
			if !checkSafe(graph, state, n) {
				return false
			}
		}

		state[v] = Safe
		return true

	case Explored:
		return false

	case Safe:
		return true
	}

	panic("should not reach here")
}
// tag-[深度优先搜索]
// leetcode841:钥匙和房间 dfs
func canVisitAllRooms_(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if visited[i] {
			return
		}
		visited[i] = true
		for _, t := range rooms[i] {
			dfs(t)
		}
	}
	dfs(0)
	for i := range visited {
		if !visited[i] {
			return false
		}
	}
	return true
}

func Test_canVisitAllRooms(t *testing.T) {
	fmt.Println(canVisitAllRooms_([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
	fmt.Println(canVisitAllRooms_([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms_([][]int{{1}, {1}}))
	fmt.Println(canVisitAllRooms_([][]int{{1}, {2}, {}, {3}}))
	fmt.Println(canVisitAllRooms_([][]int{{1, 2}, {2, 1}, {1}}))
}// tag-[深度优先搜索]
// leetcode638:每日一题,dfs记忆化搜索（非回溯）
func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	filterSpecial := make([][]int, 0)
	totalCnt, totalPrice := 0, 0
	// 先过滤出优先的大礼包
	for i := range special {
		totalCnt, totalPrice = 0, 0
		for j, v := range special[i][:n] {
			totalCnt++
			totalPrice += v * price[j]
		}
		if totalCnt > 0 && totalPrice > special[i][n] {
			filterSpecial = append(filterSpecial, special[i])
		}
	}
	dp := make(map[string]int)
	var dfs func(needs []byte) int
	dfs = func(needs []byte) int {
		if v, ok := dp[string(needs)]; ok {
			return v
		}
		ans := 0
		for i := range needs {
			ans += int(needs[i]) * price[i]
		}
		nextNeeds := make([]byte, len(needs))
	outer:
		for i := range filterSpecial {
			for j, v := range filterSpecial[i][:n] {
				if v > int(needs[j]) {
					continue outer
				}
				nextNeeds[j] = byte(int(needs[j]) - v)
			}
			ans = min(ans, dfs(nextNeeds)+filterSpecial[i][n])
		}
		dp[string(needs)] = ans
		return ans
	}
	needs_ := make([]byte, len(needs))
	for i := range needs {
		needs_[i] = byte(needs[i])
	}
	return dfs(needs_)
}
// tag-[深度优先搜索]
// leetcode周赛第三题
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	// 建树
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		v := parents[i]
		g[v] = append(g[v], i)
	}
	ans := 0
	maxScore := 0
	var dfs func(int) int
	dfs = func(v int) int {
		size, score := 1, 1
		for _, w := range g[v] {
			sz := dfs(w)
			size += sz
			score *= sz // 左右相乘
		}
		if v > 0 {
			score *= n - size
		}
		if score > maxScore {
			maxScore, ans = score, 1
		} else if score == maxScore {
			ans++
		}
		return size
	}
	dfs(0)
	return ans
}

// 周赛第三题参考
func countHighestScoreNodes_(parents []int) (ans int) {
	n := len(parents)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		v := parents[w]
		g[v] = append(g[v], w) // 建树
	}

	maxScore := 0
	var dfs func(int) int
	dfs = func(v int) int {
		size, score := 1, 1
		for _, w := range g[v] {
			sz := dfs(w)
			size += sz
			score *= sz // 由于是二叉树所以 score 最大约为 (1e5/3)^3，在 64 位整数范围内
		}
		if v > 0 {
			score *= n - size
		}
		if score > maxScore {
			maxScore, ans = score, 1
		} else if score == maxScore {
			ans++
		}
		return size
	}
	dfs(0)
	return
}// tag-[深度优先搜索]
// leetcode924: 尽量减少恶意软件的传播
// dfs解法
func minMalwareSpread_(graph [][]int, initial []int) int {
	size := len(graph)
	color := make([]int, size)
	var dfs func(node, c int)
	dfs = func(node, c int) {
		color[node] = c
		for i := range graph {
			if graph[i][node] == 1 && color[i] == 0 {
				dfs(i, c)
			}
		}
	}
	// 上色
	c := 0
	for i := 0; i < size; i++ {
		if color[i] == 0 {
			c++
			dfs(i, c)
		}
	}
	num := make([]int, c+1)
	// 统计每种颜色的个数
	for i := 0; i < size; i++ {
		num[color[i]]++
	}
	// 统计initial中颜色种类
	colorCnt := make(map[int]int)
	for i := range initial {
		colorCnt[color[initial[i]]]++
	}
	sort.Ints(initial)
	maxn, ans := -1, -1
	for _, v := range initial {
		if colorCnt[color[v]] == 1 && (maxn < num[color[v]] || ans == -1) {
			maxn = num[color[v]]
			ans = v
		}
	}
	if ans == -1 {
		return initial[0]
	}
	return ans
}

func Test_minMalwareSpread(t *testing.T) {
	fmt.Println(minMalwareSpread([][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 1}, {0, 0, 1, 1}}, []int{3, 1}))
	fmt.Println(minMalwareSpread([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, []int{0, 1, 2}))
	fmt.Println(minMalwareSpread_([][]int{{1, 0, 0, 0, 0, 0}, {0, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 0, 0}, {0, 0, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1}}, []int{2, 3}))
}
// tag-[深度优先搜索]
// leetcode464: 我能赢吗
// 记忆化搜索，博弈
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	dp := make([]int, 1<<maxChoosableInteger)
	var dfs func(state, remain int) bool
	dfs = func(state, remain int) bool {
		if dp[state] != 0 {
			return dp[state] == 1
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			if (1<<(i-1))&state != 0 {
				continue
			}
			if i >= remain || !dfs((1<<(i-1))|state, remain-i) {
				dp[state] = 1
				return true
			}
		}
		dp[state] = -1
		return false
	}
	if maxChoosableInteger > desiredTotal {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	return dfs(0, desiredTotal)
}

func Test_canIWin(t *testing.T) {
	fmt.Println(canIWin(3, 5))
	// fmt.Println(canIWin(10, 11))
}// tag-[深度优先搜索]
// leetcode375: 猜数字大小II
// 记忆化搜索
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l >= r {
			return 0
		}
		if dp[l][r] != 0 {
			return dp[l][r]
		}
		ans := 0x3f3f3f3f
		for k := l; k <= r; k++ {
			ans = min(ans, k+max(dfs(l, k-1), dfs(k+1, r)))
		}
		dp[l][r] = ans
		return ans
	}
	return dfs(1, n)
}

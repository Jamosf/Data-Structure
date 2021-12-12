package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[图]
// leetcode127: 单词接龙
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

func Test_ladderLength(t *testing.T) {
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
// tag-[图]
// leetcode1976: 达到目的地的方案数
func countPaths(n int, roads [][]int) int {
	edge := make([][]int, n)
	for i := range edge {
		edge[i] = make([]int, n)
	}
	for i := range roads {
		x, y, v := roads[i][0], roads[i][1], roads[i][2]
		edge[x][y] = v
		edge[y][x] = v
	}
	return dijkstraWithPathCount(n, edge)
}

func dijkstra(n int, edge [][]int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	used := make([]bool, n)
	dist[0] = 0
	for {
		// 1. 遍历取最小未使用的顶点
		u := -1
		for v, ok := range used {
			if !ok && (v < 0 || dist[v] < dist[u]) {
				u = v
			}
		}
		used[u] = true
		// 2. 更新最小顶点邻接的节点路径
		// for x := 0; x < n; x++ {
		// 	if edge[u][x] > 0 && dist[x] > dist[u]+edge[u][x] {
		// 		dist[x] = dist[u] + edge[u][x]
		// 	}
		// }
		// 2. 简写
		for w, wt := range edge[u] {
			if nd := dist[u] + wt; nd < dist[w] {
				dist[w] = nd
			}
		}
	}
}

func dijkstraWithPathCount(n int, edge [][]int) int {
	mod := int64(1e9 + 7)
	dist := make([]int64, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	cnt := make([]int64, n)
	used := make([]bool, n)
	dist[0] = 0
	cnt[0] = 1
	for {
		// 1. 遍历取最小未使用的顶点
		u := -1
		for v, ok := range used {
			if !ok && (u < 0 || dist[v] < dist[u]) {
				u = v
			}
		}
		if u < 0 {
			break
		}
		used[u] = true
		// 2. 简写
		for w, wt := range edge[u] {
			if wt == 0 {
				continue
			}
			if nd := dist[u] + int64(wt); nd < dist[w] {
				dist[w] = nd
				cnt[w] = cnt[u]
			} else if nd == dist[w] {
				cnt[w] = (cnt[w] + cnt[u]) % mod
			}
		}
	}
	return int(cnt[n-1] % mod)
}

func numberOfCombinations(num string) int {
	n := len(num)
	dp0, dp1 := 0, 0
	mod := int(1e9 + 7)
	for i := n - 1; i >= 0; i-- {
		if num[i] == '0' {
			dp0 = dp0 + dp1 + 1
		} else {
			dp1 = dp0 + dp1 + 1
		}
	}
	return dp1 % mod
}

func Test_number(t *testing.T) {
	fmt.Println(numberOfCombinations("327"))
}
// tag-[图]
type graphNode struct {
	Val       int
	Neighbors []*graphNode
}

// leetcode133:克隆图
func cloneGraph(node *graphNode) *graphNode {
	visited := make(map[*graphNode]*graphNode)
	var dfs func(v *graphNode) *graphNode
	dfs = func(v *graphNode) *graphNode {
		if v == nil {
			return nil
		}
		if _, ok := visited[v]; ok {
			return visited[v]
		}
		root := &graphNode{Val: v.Val}
		visited[v] = root
		for _, neighbor := range node.Neighbors {
			root.Neighbors = append(root.Neighbors, dfs(neighbor))
		}
		return root
	}
	return dfs(node)
}

// tag-[图]
// leetcode1263: 推箱子
// 箱子和目标有一条路径，最短路
// A*启发式搜索算法
type state struct {
	cost, heu int    // cost表示起点到该点、heu表示该点到终点
	bits      uint64 // 低32位表示箱子、高32位表示人
}

type minTop []state

func (h minTop) Len() int      { return len(h) }
func (h minTop) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h minTop) Less(i, j int) bool {
	return h[i].cost+h[i].heu < h[j].cost+h[j].heu
}
func (h *minTop) Push(x interface{}) { *h = append(*h, x.(state)) }
func (h *minTop) Pop() interface{} {
	old := *h
	top := len(old) - 1
	*h = old[:top]
	return old[top]
}

const (
	mask = 0xffff
)

var (
	G        [][]byte
	row, col int
	explored map[uint64]int
	vis      [20][20]bool
	h        = new(minTop)
	final    uint32
	offset   = [5]int{0, 1, 0, -1, 0}
	next     = make([]uint64, 0, 4)
)

func isFinal(cur uint64) bool { return uint32(cur) == final }

// 启发函数：如果只能沿着上下左右搜索，则可以使用曼哈顿距离
func heuristic(cur uint64) int {
	return abs(int(cur&mask)-int(final&mask)) +
		abs(int((cur>>16)&mask)-int((final>>16)&mask))
}

func in(r, c int) bool {
	return 0 <= r && r < row && 0 <= c && c < col
}

func dfs_(r, c int) {
	vis[r][c] = true
	for i := 0; i < 4; i++ {
		nr, nc := r+offset[i], c+offset[i+1]
		if in(nr, nc) && G[nr][nc] != '#' && !vis[nr][nc] {
			dfs_(nr, nc)
		}
	}
}

func iter(r, c int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			vis[i][j] = false
		}
	}
	dfs_(r, c)
}

func getNext(cur uint64) []uint64 {
	next = next[:0]
	rs, cs := int((cur>>32)&mask), int((cur>>48)&mask) // 人的位置
	rb, cb := int(cur&mask), int((cur>>16)&mask)       // 箱子的位置
	old := G[rb][cb]
	G[rb][cb] = '#'
	iter(rs, cs) // 遍历人可以走到的地方
	G[rb][cb] = old
	for i := 0; i < 4; i++ {
		// 箱子后方和前进
		pr, pc := rb+offset[i], cb+offset[i+1]
		tr, tc := rb-offset[i], cb-offset[i+1]
		// 箱子前方不是墙壁，后方人可到达即认为是合理。箱子四个方向之一即当前人站的位置，必然vis为true，主要是考虑箱子其他三个方位，人是否可达。
		if in(pr, pc) && in(tr, tc) && vis[pr][pc] && G[tr][tc] != '#' {
			next = append(next, (cur<<32)^uint64(tr^(tc<<16))) // 人的位置更新为箱子，箱子更新为新的位置
		}
	}
	return next
}

// a*寻路算法：dijkstra和最佳优先搜索算法的结合体，选择下一个点的标准为：起点到该点的距离+该点到终点的直线距离
func aStar(init uint64) int {
	explored = make(map[uint64]int)
	*h = (*h)[:0]
	explored[init] = 0
	heap.Push(h, state{cost: 0, heu: heuristic(init), bits: init})
	for len(*h) > 0 {
		cur := heap.Pop(h).(state)
		if isFinal(cur.bits) {
			return cur.cost
		}
		newCost := cur.cost + 1
		for _, ch := range getNext(cur.bits) {
			if oldCost, exist := explored[ch]; !exist || oldCost > newCost {
				explored[ch] = newCost
				heap.Push(h, state{cost: newCost, heu: heuristic(ch), bits: ch})
			}
		}
	}
	return -1
}

func minPushBox(grid [][]byte) int {
	G = grid
	row, col = len(G), len(G[0])
	init := uint64(0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch G[i][j] {
			case 'S':
				init ^= uint64(i<<32) ^ uint64(j<<48)
			case 'B':
				init ^= uint64(i) ^ uint64(j<<16)
			case 'T':
				final = uint32(i) ^ uint32(j<<16)
			}
		}
	}
	iter(int((init>>32)&mask), int(init>>48))
	if !vis[init&mask][(init>>16)&mask] { // 玩家无法走到箱子
		return -1
	}
	iter(int(init&mask), int((init>>16)&mask))
	if !vis[final&mask][(final>>16)&mask] { // 箱子无法走到目标
		return -1
	}
	return aStar(init)
}
// tag-[图]
// leetcode815: 公交路线
// 主要是建图，将公交线路看成一个点
func numBusesToDestination(routes [][]int, source int, target int) int {
	n := len(routes)
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	m := make(map[int][]int)
	for i := range routes {
		for j := range routes[i] {
			v := m[routes[i][j]]
			for _, k := range v {
				g[i][k] = 1
				g[k][i] = 1
			}
			v = append(v, i)
			m[routes[i][j]] = v
		}
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	queue := make([]int, 0)
	queue = append(queue, m[source]...)
	for i := range queue {
		dist[queue[i]] = 1
	}
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		for i := range g {
			if g[i][t] == 1 {
				if dist[i] == math.MaxInt32 {
					queue = append(queue, i)
				}
				if dist[i] > dist[t]+1 {
					dist[i] = dist[t] + 1
				}
			}
		}
	}
	minn := math.MaxInt32
	for _, v := range m[target] {
		if v == 0 {
			continue
		}
		minn = min(minn, dist[v])
	}
	if minn == 1 && source == target {
		return 0
	}
	if minn == math.MaxInt32 {
		return -1
	}
	return minn
}

func Test_numBusesToDestination(t *testing.T) {
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))
	fmt.Println(numBusesToDestination([][]int{{1, 7}, {3, 5}}, 5, 5))
	fmt.Println(numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12))
}

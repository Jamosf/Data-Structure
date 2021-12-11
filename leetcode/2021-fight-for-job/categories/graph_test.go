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
}// tag-[图]
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
}// tag-[图]
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

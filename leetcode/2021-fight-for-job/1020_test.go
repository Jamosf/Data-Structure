// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"testing"
)

// leetcode453
func minMoves(nums []int) int {
	n := len(nums)
	sum, minn := 0, math.MaxInt32
	for i := range nums {
		sum += nums[i]
		minn = min(minn, nums[i])
	}
	return sum - minn*n
}

func Test_minMoves(t *testing.T) {
	fmt.Println(minMoves([]int{1, 2, 3}))
	fmt.Println(minMoves([]int{1, 1, 1}))
	fmt.Println(minMoves([]int{1, 2, 3, 4}))
}

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
}

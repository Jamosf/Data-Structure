package days

import (
	"fmt"
	"math"
	"testing"
)

// tag-[数组]
// leetcode334: 递增的三元子序列（每日一题）
func increasingTriplet(nums []int) bool {
	small, mid := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v <= small {
			small = v
		} else if v <= mid {
			mid = v
		} else {
			return true
		}
	}
	return false
}

func Test_increasingTriplet(t *testing.T) {
	fmt.Println(increasingTriplet([]int{1, 1, 1, 1}))
}

// tag-[数学]
// leetcode204: 计数质数
// 思路：排除法
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	cnt := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			fmt.Println(i)
			cnt++
		}
	}
	return cnt
}

func Test_countPrimes(t *testing.T) {
	fmt.Println(countPrimes(10))
}

// tag-[最短路]
// leetcode787: K站中转内最便宜的航班
// 思路：回溯思想，超时
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, f := range flights {
		i, j, v := f[0], f[1], f[2]
		grid[i][j] = v
	}
	vis := make([]bool, n)
	cnt, minn := 0, math.MaxInt32
	cost := 0
	var dfs func(curr int)
	dfs = func(curr int) {
		if cnt >= k+2 || cost >= minn {
			return
		}
		if curr == dst {
			minn = cost
			return
		}
		cnt++
		for j := range grid[curr] {
			if grid[curr][j] > 0 && !vis[j] {
				cost += grid[curr][j]
				vis[j] = true
				dfs(j)
				vis[j] = false
				cost -= grid[curr][j]
			}
		}
		cnt--
	}
	vis[src] = true
	dfs(src)
	if minn == math.MaxInt32 {
		return -1
	}
	return minn
}

// tag-[最短路]
// leetcode787: K站中转内最便宜的航班
// 思路：bfs思想
func findCheapestPrice_(n int, flights [][]int, src int, dst int, k int) int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, f := range flights {
		i, j, v := f[0], f[1], f[2]
		grid[i][j] = v
	}
	path := make([]int, n)
	for i := range path {
		path[i] = math.MaxInt32
	}
	path[src] = 0
	queue := make([][2]int, 0, n)
	queue = append(queue, [2]int{src, 0})
	for len(queue) > 0 && k+1 > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			v := queue[i][0]
			for j := range grid[v] {
				dist := queue[i][1] + grid[v][j]
				if grid[v][j] > 0 && path[j] > dist && path[dst] > dist {
					path[j] = dist
					if j != dst {
						queue = append(queue, [2]int{j, dist})
					}
				}
			}
		}
		queue = queue[size:]
		k--
	}
	if path[dst] == math.MaxInt32 {
		return -1
	}
	return path[dst]
}

func Test_findCheapestPrice_(t *testing.T) {
	fmt.Println(findCheapestPrice_(4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1))
	fmt.Println(findCheapestPrice_(3, [][]int{{0, 1, 2}, {1, 2, 1}, {2, 0, 10}}, 1, 2, 1))
}

// tag-[最短路]
// leetcode787: K站中转内最便宜的航班
// 思路：动态规划dp[k][i] = dp[k-1][j] + cost(i, j)
func findCheapestPrice__(n int, flights [][]int, src int, dst int, k int) int {
	inf := 10000*101 + 1
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][src] = 0
	for t := 1; t <= k+1; t++ {
		for _, f := range flights {
			i, j, v := f[0], f[1], f[2]
			dp[t][j] = min(dp[t][j], dp[t-1][i]+v)
		}
	}
	ans := inf
	for t := 1; t <= k+1; t++ {
		ans = min(ans, dp[t][dst])
	}
	if ans == inf {
		return -1
	}
	return ans
}

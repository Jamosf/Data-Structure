package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)
// tag-[并查集]
// 第一题
// leetcode130：被围绕的区域
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	u := newUnionFind(m*n + 1)
	dummyNode := m * n
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if isEdge(i, j, m, n) {
					u.union(node(i, j, n), dummyNode)
				} else {
					if i-1 >= 0 && board[i-1][j] == 'O' {
						u.union(node(i, j, n), (i-1)*n+j)
					}
					if i+1 < m && board[i+1][j] == 'O' {
						u.union(node(i, j, n), (i+1)*n+j)
					}
					if j-1 >= 0 && board[i][j-1] == 'O' {
						u.union(node(i, j, n), i*n+j-1)
					}
					if j+1 < n && board[i][j+1] == 'O' {
						u.union(node(i, j, n), i*n+j+1)
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if u.isConnected(node(i, j, n), dummyNode) {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func node(i, j, n int) int {
	return i*n + j
}

func isEdge(i, j, m, n int) bool {
	return i == 0 || i == m-1 || j == 0 || j == n-1
}

func Test_solve(t *testing.T) {
	b := [][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}
	fmt.Println(b)
	solve(b)
	fmt.Println(b)
}
// tag-[并查集]
// 第二题
// leetcode684: 冗余连接
// 删除图中多余的边
func findRedundantConnection(edges [][]int) []int {
	nodeNum := len(edges)
	u := newUnionFind(nodeNum + 1)
	ans := make([]int, 2)
	for i := 0; i < nodeNum; i++ {
		if u.isConnected(edges[i][0], edges[i][1]) {
			ans[0], ans[1] = edges[i][0], edges[i][1]
		} else {
			u.union(edges[i][0], edges[i][1])
		}
	}
	return ans
}
// tag-[并查集]
// 第三题
// leetcode547: 省份数量
func findCircleNum(isConnected [][]int) int {
	cityNum := len(isConnected)
	u := newUnionFind(cityNum)
	for i := 0; i < cityNum; i++ {
		for j := i + 1; j < cityNum; j++ {
			if isConnected[i][j] == 1 {
				u.union(i, j)
			}
		}
	}
	return u.count()
}
// tag-[并查集]
// 第四题
// leetcode1905：统计子岛屿
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	u := newUnionFind(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				if i-1 >= 0 && grid2[i-1][j] == 1 {
					u.union(node(i, j, n), (i-1)*n+j)
				}
				if i+1 < m && grid2[i+1][j] == 1 {
					u.union(node(i, j, n), (i+1)*n+j)
				}
				if j-1 >= 0 && grid2[i][j-1] == 1 {
					u.union(node(i, j, n), i*n+j-1)
				}
				if j+1 < n && grid2[i][j+1] == 1 {
					u.union(node(i, j, n), i*n+j+1)
				}
			}
		}
	}
	root := make(map[int][][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				tmp := root[u.find(i*n+j)]
				tmp = append(tmp, []int{i, j})
				root[u.find(i*n+j)] = tmp
			}
		}
	}
	ans := 0
	for _, vv := range root {
		l := len(vv)
		cnt := 0
		for _, v := range vv {
			if grid1[v[0]][v[1]] == 1 {
				cnt++
			}
		}
		if l == cnt {
			ans++
		}
	}
	return ans
}

func Test_countSubIslands(t *testing.T) {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
	fmt.Println(countSubIslands(grid1, grid2))
}// tag-[并查集]
// leetcode1998: 数组的最大公因数排序
func gcdSort(nums []int) bool {
	n := len(nums)
	u := newUnionFind(1e5 + 1)
	tmp := make([]int, n)
	for i := range nums {
		for a := 2; a*a <= nums[i]; a++ {
			if nums[i]%a == 0 {
				b := nums[i] / a
				u.union(nums[i], a)
				u.union(nums[i], b)
			}
		}
		tmp[i] = nums[i]
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] != tmp[i] && !u.isConnected(nums[i], tmp[i]) {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Test_gcdSort(t *testing.T) {
	fmt.Println(gcdSort([]int{8, 9, 4, 2, 3}))
}
// tag-[并查集]
// leetcode924: 尽量减少恶意软件的传播
// 并查集
func minMalwareSpread(graph [][]int, initial []int) int {
	size := len(graph)
	uf := newUnionFind(size)
	for i := range graph {
		for j := range graph[i] {
			if i != j && graph[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	// 统计每个root下挂的结点个数
	m := make(map[int]int)
	for i := range initial {
		m[uf.find(initial[i])]++
	}
	sort.Ints(initial)
	maxn, ans := -1, -1
	for i := range initial {
		r := uf.find(initial[i])
		s := uf.size(initial[i])
		if m[r] == 1 && (maxn < s || ans == -1) {
			maxn = s
			ans = initial[i]
		}
	}
	if ans == -1 {
		return initial[0]
	}
	return ans
}
// tag-[并查集]
// leetcode721: 账户合并
func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)
	u := newUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if hasSameEmail(accounts[i][1:], accounts[j][1:]) {
				u.union(i, j)
			}
		}
	}
	// 统计每个root下面的节点
	m := make(map[int][]int)
	for i := 0; i < n; i++ {
		r := u.find(i)
		m[r] = append(m[r], i)
	}
	var ans [][]string
	for k, v := range m {
		ans = append(ans, []string{accounts[k][0]})
		for i := range v {
			ans[len(ans)-1] = append(ans[len(ans)-1], accounts[v[i]][1:]...)
		}
	}
	for i := range ans {
		v := []string{ans[i][0]}
		sort.Strings(ans[i][1:])
		for j := 1; j < len(ans[i]); j++ {
			if ans[i][j-1] != ans[i][j] {
				v = append(v, ans[i][j])
			}
		}
		ans[i] = v
	}
	return ans
}

func hasSameEmail(a, b []string) bool {
	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				return true
			}
		}
	}
	return false
}

func Test_accountsMerge(t *testing.T) {
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}))
}

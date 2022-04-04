package days

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// tag-[����]
// ÿ��һ��
// leetcode1001: ��������
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	light := make(map[[2]int]bool)
	row, col, left, right := make(map[int]int), make(map[int]int), make(map[int]int), make(map[int]int)
	for i := 0; i < len(lamps); i++ {
		x, y := lamps[i][0], lamps[i][1]
		if light[[2]int{x, y}] {
			continue
		}
		light[[2]int{x, y}] = true
		row[x]++
		col[y]++
		left[x+y]++
		right[n+x-y]++
	}
	ans := make([]int, len(queries))
	direct := [9][2]int{{0, 0}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, -1}}
	for p := 0; p < len(queries); p++ {
		x, y := queries[p][0], queries[p][1]
		if row[x] > 0 || col[y] > 0 || left[x+y] > 0 || right[n+x-y] > 0 {
			ans[p] = 1
		}
		for k := 0; k < len(direct); k++ {
			i, j := x+direct[k][0], y+direct[k][1]
			if i >= 0 && i < n && j >= 0 && j < n {
				if light[[2]int{i, j}] {
					light[[2]int{i, j}] = false
					if row[i] > 0 {
						row[i]--
					}
					if col[j] > 0 {
						col[j]--
					}
					if left[i+j] > 0 {
						left[i+j]--
					}
					if right[n+i-j] > 0 {
						right[n+i-j]--
					}
				}
			}
		}
	}
	return ans
}

func Test_gridIllumination(t *testing.T) {
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 0}}))
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 1}}))
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {0, 4}}, [][]int{{0, 4}, {0, 1}, {1, 4}}))
}

// tag-[��ϣ��]
// ÿ��һ��
// leetcode2006: ��ľ���ֵΪ K ��������Ŀ
func countKDifference(nums []int, k int) int {
	m := make(map[int]int)
	cnt := 0
	for i := range nums {
		cnt += m[nums[i]+k] + m[nums[i]-k] + m[k-nums[i]]
		m[nums[i]]++
	}
	return cnt
}

func Test_count(t *testing.T) {
	fmt.Println(countKDifference([]int{3, 2, 1, 5, 4}, 2))
}

// tag-[��ѧ]
// ÿ��һ��
// leetcode1447: ������
func simplifiedFractions(n int) []string {
	var gcd func(m, n int) int
	gcd = func(m, n int) int {
		if n == 0 {
			return m
		}
		return gcd(n, m%n)
	}
	var ans []string
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) == 1 {
				ans = append(ans, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return ans
}

// tag-[�����������]
// ÿ��һ��
// leetcode1020: �ɵص�����
func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 && i >= m && j < 0 && j >= n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	// ��û���±߽�
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	// ��û���ұ߽�
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

// tag-[��������]
// ÿ��һ��
// leetcode1984: ѧ����������С��ֵ
func minimumDifference(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	minn := math.MaxInt32
	l, r := 0, k-1
	for r < n {
		minn = min(minn, nums[r]-nums[l])
		r++
		l++
	}
	return minn
}

func minimumDifference_(nums []int, k int) int {
	sort.Ints(nums)
	rep := math.MaxInt32
	for i, num := range nums[:len(nums)-k+1] {
		rep = func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}(rep, nums[i+k-1]-num)
	}
	return rep
}

func Test_minimumDifference_(t *testing.T) {
	fmt.Println(minimumDifference_([]int{1, 3, 7}, 2))
}

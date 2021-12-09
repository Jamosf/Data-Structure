// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// tag-[字符串]
// leetcode1945: 字符串转化后的各位数字之和
func getLucky(s string, k int) int {
	ans := make([]byte, 0)
	for i := range s {
		v := s[i] - 'a' + 1
		if v > 9 {
			ans = append(ans, v/10)
		}
		ans = append(ans, v%10)
	}
	tmp := countSum(string(ans), k)
	ret, _ := strconv.Atoi(tmp)
	return ret
}

func countSum(s string, k int) string {
	if k == 0 {
		return s
	}
	tmp := 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v >= '0' {
			v -= '0'
		}
		tmp += int(v)
	}
	return countSum(strconv.Itoa(tmp), k-1)
}

func Test_getLucky(t *testing.T) {
	fmt.Println(getLucky("dbvmfhnttvr", 5))
}

// tag-[字符串]
// leetcode1946: 子字符串突变后可能得到的最大整数
func maximumNumber(num string, change []int) string {
	b := []byte(num)
	cnt := 0
	for i := 0; i < len(b); i++ {
		v := change[b[i]-'0']
		if int(b[i]-'0') < v {
			b[i] = byte(v + '0')
			cnt++
		} else if int(b[i]-'0') > v {
			if cnt != 0 {
				break
			}
		}
	}
	return string(b)
}

func Test_maximumNumber(t *testing.T) {
	fmt.Println(maximumNumber("334111", []int{0, 9, 2, 3, 3, 2, 5, 5, 5, 5}))
}

// tag-[回溯]
// leetcode1947: 最大兼容性评分和
func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
	maxn := 0
	sum := 0
	visited := make([]bool, m)
	var backtracking func(level int)
	backtracking = func(level int) {
		if level == m {
			maxn = max(maxn, sum)
			return
		}
		for i := 0; i < m; i++ {
			if !visited[i] {
				v := caclSum(students[level], mentors[i])
				visited[i] = true
				sum += v
				backtracking(level + 1)
				sum -= v
				visited[i] = false
			}
		}
	}
	backtracking(0)
	return maxn
}

func caclSum(a, b []int) int {
	cnt := 0
	for i := range a {
		if a[i] == b[i] {
			cnt++
		}
	}
	return cnt
}

func Test_maxCompatibilitySum(t *testing.T) {
	fmt.Println(maxCompatibilitySum([][]int{{0, 1, 0, 1, 1, 1}, {1, 0, 0, 1, 0, 1}, {1, 0, 1, 1, 0, 0}}, [][]int{{1, 0, 0, 0, 0, 1}, {0, 1, 0, 0, 1, 1}, {0, 1, 0, 0, 1, 1}}))
}

// tag-[字典树]
type folder struct {
	son map[string]*folder
	val string // 文件夹名称
	del bool   // 删除标记
}

// leetcode1948：删除系统中的重复文件夹
func deleteDuplicateFolder(paths [][]string) (ans [][]string) {
	root := &folder{}
	for _, path := range paths {
		// 将 path 加入字典树
		f := root
		for _, s := range path {
			if f.son == nil {
				f.son = map[string]*folder{}
			}
			if f.son[s] == nil {
				f.son[s] = &folder{}
			}
			f = f.son[s]
			f.val = s
		}
	}

	folders := map[string][]*folder{} // 存储括号表达式及其对应的文件夹节点列表
	var dfs func(*folder) string
	dfs = func(f *folder) string {
		if f.son == nil {
			return "(" + f.val + ")"
		}
		expr := make([]string, 0, len(f.son))
		for _, son := range f.son {
			expr = append(expr, dfs(son))
		}
		sort.Strings(expr)
		subTreeExpr := strings.Join(expr, "") // 按字典序拼接所有子树
		folders[subTreeExpr] = append(folders[subTreeExpr], f)
		return "(" + f.val + subTreeExpr + ")"
	}
	dfs(root)

	for _, fs := range folders {
		if len(fs) > 1 { // 将括号表达式对应的节点个数大于 1 的节点全部删除
			for _, f := range fs {
				f.del = true
			}
		}
	}

	// 再次 DFS 这颗字典树，仅访问未被删除的节点，并将路径记录到答案中
	path := []string{}
	var dfs2 func(*folder)
	dfs2 = func(f *folder) {
		if f.del {
			return
		}
		path = append(path, f.val)
		ans = append(ans, append([]string(nil), path...))
		for _, son := range f.son {
			dfs2(son)
		}
		path = path[:len(path)-1]
	}
	for _, son := range root.son {
		dfs2(son)
	}
	return
}

func Test_deleteDuplicateFolder(t *testing.T) {
	fmt.Println(deleteDuplicateFolder([][]string{{"a"}, {"a", "x"}, {"a", "x", "y"}, {"a", "z"}, {"b"}, {"b", "x"}, {"b", "x", "y"}, {"b", "z"}, {"b", "w"}}))
}

// tag-[矩阵]
// leetcode1895: 最大的幻方
func largestMagicSquare(grid [][]int) int {
	print_matrix(grid)
	maxn := 0
	m, n := len(grid), len(grid[0])
	sumi := make([][]int, m)
	for i := range sumi {
		sumi[i] = make([]int, n)
	}
	sumj := make([][]int, n)
	for i := range sumj {
		sumj[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		sumi[i][0] = grid[i][0]
		for j := 1; j < n; j++ {
			sumi[i][j] = sumi[i][j-1] + grid[i][j]
		}
	}
	for j := 0; j < n; j++ {
		sumj[j][0] = grid[0][j]
		for i := 1; i < m; i++ {
			sumj[j][i] = sumj[j][i-1] + grid[i][j]
		}
	}

	var check func(i, j, endi, endj int) bool
	check = func(i, j, endi, endj int) bool {
		if endi < 0 || endi >= m || endj < 0 || endj >= n {
			return false
		}
		v := sumi[i][endj] - sumi[i][j] + grid[i][j]
		for posx := i + 1; posx <= endi; posx++ {
			if sumi[posx][endj]-sumi[posx][j]+grid[posx][j] != v {
				return false
			}
		}
		for posy := j; posy <= endj; posy++ {
			if sumj[posy][endi]-sumj[posy][i]+grid[i][posy] != v {
				return false
			}
		}
		sumk := 0
		sumkk := 0
		for k := 0; k <= endi-i; k++ {
			sumk += grid[i+k][j+k]
			sumkk += grid[i+k][endj-k]
		}
		if sumk != v || sumkk != v {
			return false
		}
		return true
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < min(m, n); k++ {
				if check(i, j, i+k, j+k) {
					maxn = max(maxn, k+1)
				}
			}
		}
	}
	return maxn
}

func Test_largestMagicSquare(t *testing.T) {
	fmt.Println(largestMagicSquare([][]int{{1, 17, 15, 17, 5, 16, 8, 9}, {1, 19, 11, 18, 8, 18, 3, 18}, {6, 6, 5, 8, 3, 15, 6, 11}, {19, 5, 6, 11, 9, 2, 14, 13}, {12, 16, 16, 15, 14, 18, 10, 7}, {3, 11, 15, 15, 7, 1, 9, 8}, {15, 5, 11, 17, 18, 20, 14, 17}, {13, 17, 7, 20, 12, 2, 13, 19}}))
}

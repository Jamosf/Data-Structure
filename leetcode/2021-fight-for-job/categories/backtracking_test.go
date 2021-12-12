package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[回溯]
// 第二题
// leetcode17: 电话号码的字母组合
func letterCombinations(digits string) []string {
	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	size := len(digits)
	if size == 0 {
		return nil
	}
	var ret []string
	var tmp []byte
	var dfs func(int)
	dfs = func(level int) {
		if level == size {
			ret = append(ret, string(tmp))
			return
		}
		// 没有标记是否访问过，所有的解都是可行的
		for j := 0; j < len(m[digits[level]-'0']); j++ {
			v := m[digits[level]-'0'][j] - 'a'
			tmp = append(tmp, v+'a')
			dfs(level + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0)
	return ret
}

func Test_letterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("22"))
}
// tag-[回溯]
// 第七题
// leetcode556: 下一个更大元素III
func nextGreaterElementIII(n int) int {
	var nums []int
	num := n
	for n != 0 {
		nums = append(nums, n%10)
		n = n / 10
	}
	sort.Ints(nums)
	var dfs func(level int)
	var all []int
	var tmp []int
	visited := make([]bool, len(nums))
	dfs = func(level int) {
		if level == len(nums) {
			sum := 0
			t := 1
			for j := len(tmp) - 1; j >= 0; j-- {
				sum += tmp[j] * t
				t *= 10
			}
			all = append(all, sum)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				tmp = append(tmp, nums[i])
				visited[i] = true
				dfs(level + 1)
				tmp = tmp[:len(tmp)-1]
				visited[i] = false
			}
		}
	}
	dfs(0)
	idx := sort.SearchInts(all, num+1)
	for i := idx; i < len(all); i++ {
		if all[i] > num && all[i] <= math.MaxInt32 {
			return all[i]
		}
	}
	return -1
}

func Test_nextGreaterElementIII(t *testing.T) {
	fmt.Println(nextGreaterElementIII(1234))
}

// tag-[回溯]
// leetcode1981: 最小化目标值与所选元素的差
func minimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])
	for i := range mat {
		sort.Ints(mat[i])
	}
	minn := math.MaxInt32
	var backtrack func(level int)
	var sum int
	var dp [71][4901]bool
	backtrack = func(level int) {
		if sum-target > minn || dp[level][sum] {
			return
		}
		dp[level][sum] = true
		if level == m {
			if minusAbs(sum, target) < minn {
				minn = minusAbs(sum, target)
			}
			return
		}
		for i := 0; i < n; i++ {
			sum += mat[level][i]
			backtrack(level + 1)
			sum -= mat[level][i]
		}
	}
	backtrack(0)
	return minn
}

func Test_min(t *testing.T) {
	fmt.Println(minimizeTheDifference([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 13))
}
// tag-[回溯]
// leetcode46: 全排列
func permute(nums []int) [][]int {
	n := len(nums)
	l := factorial(n)
	ans := make([][]int, 0, l)
	var backtracking func(level int)
	backtracking = func(level int) {
		if level == n {
			t := make([]int, n)
			copy(t, nums)
			ans = append(ans, t)
		}
		for i := level; i < n; i++ {
			nums[i], nums[level] = nums[level], nums[i]
			backtracking(level + 1)
			nums[i], nums[level] = nums[level], nums[i]
		}
	}
	backtracking(0)
	return ans
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

// tag-[回溯]
// leetcode77: 组合
func combine(n int, k int) [][]int {
	var ans [][]int
	tmp := make([]int, 0, k)
	var backtracking func(level int)
	backtracking = func(idx int) {
		if len(tmp) == k {
			t := make([]int, k)
			copy(t, tmp)
			ans = append(ans, t)
		}
		for i := idx; i <= n; i++ {
			if len(tmp)+(n-i+1) >= k {
				tmp = append(tmp, i)
				backtracking(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtracking(1)
	return ans
}

func Test_combine(t *testing.T) {
	fmt.Println(combine(4, 3))
}
// tag-[回溯]
// 第一题
// leetcode79: 单词搜索
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	pos := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	find := false
	var dfs func(i, j int, idx int)
	dfs = func(i, j int, idx int) {
		if find || board[i][j] != word[idx] {
			return
		}
		if idx == len(word)-1 {
			find = true
			return
		}
		visited[i][j] = true
		for k := 0; k < 4; k++ {
			x, y := i+pos[k][0], j+pos[k][1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if !visited[x][y] {
					dfs(x, y, idx+1)
				}
			}
		}
		visited[i][j] = false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				for i := 0; i < m; i++ {
					visited[i] = make([]bool, n)
				}
				dfs(i, j, 0)
			}
		}
	}
	return find
}

func Test_exist(t *testing.T) {
	fmt.Println(exist([][]byte{{'a'}, {'a'}}, "aaa"))
}
// tag-[回溯]
// 第一题
// leetcode51: N皇后
func solveNQueens(n int) [][]string {
	var ans [][]string
	tmp := make([][]byte, n)
	for i := 0; i < n; i++ {
		tmp[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			tmp[i][j] = '.'
		}
	}
	col := make([]bool, n)
	left := make([]bool, 2*n)
	right := make([]bool, 2*n)
	var backtracking func(level int)
	backtracking = func(level int) {
		if level == n {
			t := make([]string, 0, n)
			for i := range tmp {
				t = append(t, string(tmp[i]))
			}
			ans = append(ans, t)
			return
		}
		i := level
		for j := 0; j < n; j++ {
			if !col[j] && !left[i+j] && !right[n-i+j] {
				col[j] = true
				left[i+j] = true
				right[n-i+j] = true
				tmp[i][j] = 'Q'
				backtracking(level + 1)
				tmp[i][j] = '.'
				col[j] = false
				left[i+j] = false
				right[n-i+j] = false
			}
		}
	}
	backtracking(0)
	return ans
}

func Test_solveNQueens(t *testing.T) {
	fmt.Println(solveNQueens(4))
}

// tag-[回溯]
// leetcode934: 最短的桥
// 求最小翻转的个数，就是求两个岛之间的最短距离
func shortestBridge(grid [][]int) int {
	queue := list.New()
	m, n := len(grid), len(grid[0])
	direction := [4][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i == m || j == n || grid[i][j] == 2 {
			return
		}
		if grid[i][j] == 0 {
			queue.PushBack(pos{i, j})
			return
		}
		grid[i][j] = 2
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	flap := false
	for i := 0; i < m; i++ {
		if flap {
			break
		}
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				flap = true
				break
			}
		}
	}
	level := 0
	for queue.Len() != 0 {
		level++
		size := queue.Len()
		for i := 0; i < size; i++ {
			v := queue.Front()
			queue.Remove(v)
			p := v.Value.(pos)
			for k := 0; k < 4; k++ {
				x, y := p.x+direction[k][0], p.y+direction[k][1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if grid[x][y] == 1 {
						return level
					}
					if grid[x][y] == 2 {
						continue
					}
					grid[x][y] = 2
					queue.PushBack(pos{x, y})
				}
			}
		}
	}
	return 0
}

type graph struct {
	vertex []string
	edges  [][]int
	n, e   int
}

// tag-[回溯]
// leetcode1986: 完成任务的最少工作时间段
// dfs解这个题
func minSessions_(tasks []int, sessionTime int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})
	ans := 20
	n := len(tasks)
	times := make([]int, n)
	var dfs func(u, k int)
	dfs = func(u, k int) {
		if k >= ans { // 再往后搜索已经不可能出现更小的解
			return
		}
		if u == n { // 搜索到了最后一个数
			ans = k
			return
		}
		// 1. 先尝试往老的时间段里面塞塞看
		for i := 0; i < k; i++ {
			if times[i]+tasks[u] <= sessionTime {
				times[i] += tasks[u]
				dfs(u+1, k)
				times[i] -= tasks[u]
			}
		}
		// 2. 使用新的时间段
		times[k] = tasks[u]
		dfs(u+1, k+1)
		times[k] = 0
	}
	dfs(0, 0)
	return ans
}

func Test_minSessions2(t *testing.T) {
	fmt.Println(minSessions([]int{1, 1, 1, 3, 3, 1}, 8))
	fmt.Println(minSessions_([]int{1, 1, 1, 3, 3, 1}, 8))
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

// tag-[回溯]
func maxmiumScore1(cards []int, cnt int) int {
	n := len(cards)
	maxn := 0
	sum := 0
	visited := make([]bool, n)
	var dfs func(level int)
	dfs = func(level int) {
		if level == cnt {
			if sum&1 == 0 {
				maxn = max(maxn, sum)
			}
			return
		}
		for i := 0; i < n; i++ {
			if !visited[i] {
				sum += cards[i]
				visited[i] = true
				dfs(level + 1)
				sum -= cards[i]
				visited[i] = false
			}
		}
	}
	dfs(0)
	return maxn
}

// tag-[矩阵]
func flipChess1(chessboard []string) int {
	m, n := len(chessboard), len(chessboard[0])
	direction := [8][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	maxn := 0
	ans := make([][]int, 0)
	var check func(i, j, dx, dy int) int
	check = func(x, y, dx, dy int) int {
		x, y = x+dx, y+dy
		step := 1
		for x >= 0 && x < m && y >= 0 && y < n {
			tmp := make([]int, 0)
			if step == 1 {
				if chessboard[x][y] == '.' || chessboard[x][y] == 'X' {
					return 0
				}
			} else {
				if chessboard[x][y] == '.' {
					return 0
				}
				if chessboard[x][y] == 'X' {
					ans = append(ans, tmp)
					return step - 1
				}
			}
			step++
			x += dx
			y += dy
			tmp = append(tmp, []int{x, y}...)
		}
		// for x := range ans{
		// 	// for x
		// }
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if chessboard[i][j] != '.' {
				continue
			}
			count := 0
			for k := 0; k < 8; k++ {
				count += check(i, j, direction[k][0], direction[k][1])
			}
			maxn = max(maxn, count)
		}
	}

	return maxn
}

func Test_flipChess1(t *testing.T) {
	fmt.Println(flipChess([]string{".X.", ".O.", "XO."}))
}
// tag-[回溯]
// leetcode78: 子集
func subsets(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	var backtracking func(lvl int)
	backtracking = func(lvl int) {
		vv := make([]int, len(tmp))
		copy(vv, tmp)
		ans = append(ans, vv)

		for i := lvl; i < n; i++ {
			if len(tmp) == 0 || nums[i] > tmp[len(tmp)-1] {
				tmp = append(tmp, nums[i])
				backtracking(lvl + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtracking(0)
	return ans
}

func Test_subSet(t *testing.T) {
	fmt.Println(subsets([]int{4, 1, 0}))
}
// tag-[回溯]
// 第二题
// leetcode22: 括号生成
func generateParenthesis(n int) []string {
	s := make([]byte, 2*n)
	for i := 0; i < 2*n; i++ {
		if i < n {
			s[i] = '('
		} else {
			s[i] = ')'
		}
	}
	ans := make([]string, 0)
	t := make([]byte, 2*n)
	var backtracking func(lvl int)
	backtracking = func(first int) {
		if first == n {
			if isValid(s) && notContain(string(s), ans) {
				copy(t, s)
				ans = append(ans, string(t))
			}
			return
		}
		for i := first; i < 2*n; i++ {
			if i == first || s[i] != s[first] {
				s[i], s[first] = s[first], s[i]
				backtracking(first + 1)
				s[i], s[first] = s[first], s[i]
			}
		}
	}
	backtracking(0)
	return ans
}

func isValid(s []byte) bool {
	cnt := 0
	for i := range s {
		if s[i] == '(' {
			cnt++
		} else {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}

func notContain(s string, t []string) bool {
	for i := range t {
		if t[i] == s {
			return false
		}
	}
	return true
}

func Test_generateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(8))
}

// tag-[回溯]
// leetcode39: 组合总和
// 回溯
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)
	sum := 0
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	minn := math.MaxInt32
	for i := range candidates {
		minn = min(minn, candidates[i])
	}
	maxIdx := target / minn
	var backtracking func(idx int)
	backtracking = func(idx int) {
		if sum == target {
			t := make([]int, len(tmp))
			copy(t, tmp)
			ans = append(ans, t)
			return
		}
		if idx == maxIdx {
			return
		}
		for i := 0; i < n; i++ {
			if len(tmp) == 0 || candidates[i] >= tmp[len(tmp)-1] {
				tmp = append(tmp, candidates[i])
				sum += candidates[i]
				backtracking(idx + 1)
				tmp = tmp[:len(tmp)-1]
				sum -= candidates[i]
			}
		}
	}
	backtracking(0)
	return ans
}

func Test_combinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

// tag-[回溯]
func countMaxOrSubsets(nums []int) int {
	var backtrace func(index int)
	m := make(map[int]int)
	n := len(nums)
	sum := 0
	v := make([]int, 0)
	maxn := 0
	backtrace = func(index int) {
		sum = or(v)
		if sum > 0 {
			m[sum]++
			if sum >= maxn {
				maxn = sum
			}
		}
		if index == n {
			return
		}
		for i := index; i < n; i++ {
			v = append(v, nums[i])
			backtrace(i + 1)
			v = v[:len(v)-1]
		}
	}
	backtrace(0)
	return m[maxn]
}

func Test_countMaxOrSubsets(t *testing.T) {
	fmt.Println(countMaxOrSubsets([]int{3, 1}))
	fmt.Println(countMaxOrSubsets([]int{2, 2, 2}))
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5}))
}

func Test_bin(t *testing.T) {
	sum := 10
	sum |= 51
	fmt.Println(sum)
	sum ^= 51
	fmt.Println(sum)
}
// tag-[回溯]
// leetcode40:回溯
func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	sort.Ints(candidates)
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	sum := 0
	var backtrace func(index int)
	backtrace = func(index int) {
		if sum == target {
			t := make([]int, len(tmp))
			copy(t, tmp)
			ans = append(ans, t)
		}
		if sum > target || index == n {
			return
		}
		for i := index; i < n; i++ {
			if i > index && candidates[i] == candidates[i-1] { // 去重思想，可以参考leetcode高赞解释
				continue
			}
			sum += candidates[i]
			tmp = append(tmp, candidates[i])
			backtrace(i + 1)
			sum -= candidates[i]
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrace(0)
	return ans
}

func Test_combinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

// tag-[回溯]
// leetcode47: 不重复全排列
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	used := make([]bool, n)
	var backtrace func(index int)
	backtrace = func(index int) {
		if index == n {
			ans = append(ans, append([]int{}, tmp...))
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			tmp = append(tmp, nums[i])
			backtrace(index + 1)
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}
	backtrace(0)
	return ans
}

func Test_permuteUnique(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 2, 3}))
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{0, 0, 0}))
	fmt.Println(permuteUnique([]int{0, 1, 0, 0, 9}))
}

// tag-[回溯]
// leetcode78
func subsets78(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	var backtracking func(lvl int)
	backtracking = func(lvl int) {
		vv := make([]int, len(tmp))
		copy(vv, tmp)
		ans = append(ans, vv)

		for i := lvl; i < n; i++ {
			if len(tmp) == 0 || nums[i] > tmp[len(tmp)-1] {
				tmp = append(tmp, nums[i])
				backtracking(lvl + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtracking(0)
	return ans
}

// tag-[回溯]
// leetcode90
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	used := make([]bool, n)
	var backtracking func(lvl int)
	backtracking = func(lvl int) {
		vv := make([]int, len(tmp))
		copy(vv, tmp)
		ans = append(ans, vv)

		for i := lvl; i < n; i++ {
			if used[i] {
				continue
			}
			if i > lvl && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if len(tmp) == 0 || nums[i] >= tmp[len(tmp)-1] {
				used[i] = true
				tmp = append(tmp, nums[i])
				backtracking(lvl + 1)
				tmp = tmp[:len(tmp)-1]
				used[i] = false
			}
		}
	}
	backtracking(0)
	return ans
}

func Test_subsetsWithDup(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0, 1, 0, 0, 9}))
}

// tag-[回溯]
// leetcode216
func combinationSum3(k int, n int) [][]int {
	tmp := make([]int, 0)
	ans := make([][]int, 0)
	used := make([]bool, 10)
	var backtrace func(index int, target int)
	backtrace = func(index int, target int) {
		if len(tmp) == k {
			fmt.Println(tmp)
			if target == 0 {
				ans = append(ans, append([]int{}, tmp...))
			}
			return
		}
		for i := index; i <= 9; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			tmp = append(tmp, i)
			backtrace(i, target-i)
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}
	backtrace(1, n)
	return ans
}

func Test_combinationSum3(t *testing.T) {
	fmt.Println(combinationSum3(3, 7))
}

// 官方题解
func combinationSum3_(k int, n int) (ans [][]int) {
	var temp []int
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		// 找到一个答案
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		// 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
		if len(temp)+10-cur < k || rest < 0 {
			return
		}
		// 跳过当前数字
		dfs(cur+1, rest)
		// 选当前数字
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return
}

// tag-[回溯]
// leetcode93:复原ip地址
func restoreIpAddresses(s string) []string {
	n := len(s)
	if n > 12 {
		return nil
	}
	ans := make([]string, 0)
	tmp := make([]string, 0)
	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == n {
			if len(tmp) == 4 {
				fmt.Println(tmp)
				ans = append(ans, strings.Join(tmp, "."))
			}
			return
		}
		for i := idx + 1; i <= n; i++ {
			if isValidIp(s[idx:i]) {
				tmp = append(tmp, s[idx:i])
				backtrace(i)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrace(0)
	return ans
}

func isValidIp(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return v >= 0 && v <= 255
}

func Test_isValidIp(t *testing.T) {
	fmt.Println(isValidIp("265"))
}

func Test_restoreIpAddresses(t *testing.T) {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("101023"))
}

// tag-[字符串/回溯]
// leetcode97:交错字符串(记忆化搜索)
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n3 != n1+n2 {
		return false
	}
	dp := make([][]bool, n1) // 记忆化，存储中间过程数据
	for i := range dp {
		dp[i] = make([]bool, n2)
	}
	var backtrace func(idx1, idx2, idx3 int) bool
	backtrace = func(idx1, idx2, idx3 int) bool {
		if idx3 == n3 {
			return true
		}
		if idx1 < n1 && idx2 < n2 && dp[idx1][idx2] {
			return false
		}
		if idx1 < n1 && s3[idx3] == s1[idx1] && backtrace(idx1+1, idx2, idx3+1) {
			return true
		}
		if idx2 < n2 && s3[idx3] == s2[idx2] && backtrace(idx1, idx2+1, idx3+1) {
			return true
		}
		if idx1 < n1 && idx2 < n2 {
			dp[idx1][idx2] = true
		}
		return false
	}
	return backtrace(0, 0, 0)
}

// tag-[回溯]
// leetcode95:回溯法
func generateTrees(n int) []*TreeNode {
	var backtrace func(start, end int) []*TreeNode
	backtrace = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		allTrees := make([]*TreeNode, 0)
		for i := start; i <= end; i++ {
			leftTrees := backtrace(start, i-1)
			rightTrees := backtrace(i+1, end)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					currTree := &TreeNode{i, nil, nil}
					currTree.Left = left
					currTree.Right = right
					allTrees = append(allTrees, currTree)
				}
			}
		}
		return allTrees
	}
	return backtrace(1, n)
}

// tag-[回溯]
// leetcode131: 分割回文串(典型的回溯)
func partition(s string) [][]string {
	n := len(s)
	ans := make([][]string, 0)
	tmp := make([]string, 0)
	isPlalindrome := func(b string) bool {
		j := len(b) - 1
		for i := 0; i < j; i++ {
			if b[i] != b[j] {
				return false
			}
			j--
		}
		return true
	}

	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == n {
			ans = append(ans, append([]string{}, tmp...))
			return
		}
		for i := idx + 1; i <= n; i++ {
			if isPlalindrome(s[idx:i]) {
				tmp = append(tmp, s[idx:i])
				backtrace(i)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrace(0)
	return ans
}

// tag-[回溯]
// leetcode131: 分割回文串(典型的回溯) 使用记忆化优化或者使用dp来预处理字符串把任意i->j是否为回文算出来
func partition_(s string) [][]string {
	n := len(s)
	ans := make([][]string, 0)
	tmp := make([]string, 0)
	dp := make([][]int8, n)
	for i := range dp {
		dp[i] = make([]int8, n)
	}
	var isPlalindrome func(i, j int) int8
	isPlalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		dp[i][j] = -1
		if s[i] == s[j] {
			dp[i][j] = isPlalindrome(i+1, j-1)
		}
		return dp[i][j]
	}

	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == n {
			ans = append(ans, append([]string{}, tmp...))
			return
		}
		for i := idx; i < n; i++ {
			if isPlalindrome(idx, i) > 0 {
				tmp = append(tmp, s[idx:i+1])
				backtrace(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrace(0)
	return ans
}

func Test_partition(t *testing.T) {
	fmt.Println(partition_("aabbaababab"))
}

// tag-[回溯]
// leetcode638:回溯法思想
func shoppingOffers_(price []int, special [][]int, needs []int) int {
	n := len(price)
	m := len(special)
	sum := 0
	minn := math.MaxInt32
	smaller := func(a, b []int) bool {
		for i := range a {
			if a[i] > b[i] {
				return false
			}
		}
		return true
	}
	var backtracking func(idx int, left []int)
	backtracking = func(idx int, left []int) {
		if idx == m {
			extra := 0
			for i := range left {
				extra += left[i] * price[i]
			}
			minn = min(minn, sum+extra)
			return
		}
		for i := idx; i < m; i++ { // 先尝试在大礼包中选择，同步更新left和最小值
			if smaller(special[i][:n], left) {
				sum += special[i][n]
				for j := range left {
					left[j] -= special[i][j]
				}
				backtracking(i, left) // 有条件的递归，i有可能无法达到m
				sum -= special[i][n]
				for j := range left {
					left[j] += special[i][j]
				}
			}
		}
		backtracking(m, left) // i不一定能达到m，因此需要最后用m来收尾处理。大礼包选择完后，还有多余的，则按照价格购买。
	}
	backtracking(0, needs)
	return minn
}

func Test_shoppingOffers(t *testing.T) {
	// fmt.Println(shoppingOffers_([]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}))
	fmt.Println(shoppingOffers_([]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}))
}

// tag-[回溯]
// leetcode周赛第二题
func nextBeautifulNumber(n int) int {
	v := n
	cnt := 0
	for v != 0 {
		cnt++
		v = v / 10
	}
	isBeautiful := func(t int) bool {
		if t == 0 {
			return false
		}
		numCnt := [9]int{}
		for t != 0 {
			numCnt[t%10]++
			t = t / 10
		}
		for i := range numCnt {
			if numCnt[i] != 0 && numCnt[i] != i {
				return false
			}
		}
		return true
	}
	convert := func(b []int) int {
		res := 0
		multi := 1
		for i := len(b) - 1; i >= 0; i-- {
			res += b[i] * multi
			multi *= 10
		}
		return res
	}
	ans := 0
	flag := false
	maxn := 6
	tmp := make([]int, 0)
	var backtrace func(idx int, delta int)
	backtrace = func(idx int, delta int) {
		if flag {
			return
		}
		if idx >= cnt {
			val := convert(tmp)
			fmt.Println(tmp)
			if val > n && isBeautiful(val) {
				flag = true
				ans = val
				return
			}
		}
		if idx > cnt+delta {
			return
		}
		for i := 1; i <= maxn; i++ {
			tmp = append(tmp, i)
			backtrace(idx+1, delta)
			tmp = tmp[:len(tmp)-1]
		}
	}
	for i := 1; i <= 6; i++ {
		backtrace(1, i-1)
		if flag {
			break
		}
	}

	return ans
}

// 暴力即可通过版本
func nextBeautifulNumber_(n int) int {
	isBeautiful := func(t int) bool {
		if t == 0 {
			return false
		}
		numCnt := [10]int{}
		for t != 0 {
			numCnt[t%10]++
			t = t / 10
		}
		for i := range numCnt {
			if numCnt[i] != 0 && numCnt[i] != i {
				return false
			}
		}
		return true
	}
	for {
		n++
		if isBeautiful(n) {
			return n
		}
	}
}

func Test_nextBeautifulNumber(t *testing.T) {
	// fmt.Println(nextBeautifulNumber(1))
	fmt.Println(nextBeautifulNumber_(3000))
}

// tag-[二叉树/回溯]
// leetcode113:路径总和,前序遍历
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	var backtracking func(r *TreeNode, v int)
	backtracking = func(r *TreeNode, v int) {
		if r == nil {
			return
		}
		tmp = append(tmp, r.Val)
		v -= r.Val
		if v == 0 && r.Left == nil && r.Right == nil {
			ans = append(ans, append([]int{}, tmp...))
		}
		backtracking(r.Left, v)
		backtracking(r.Right, v)
		tmp = tmp[:len(tmp)-1]
	}
	backtracking(root, targetSum)
	return ans
}

// tag-[回溯]
// leetcode869：每日一题
func reorderedPowerOf2(n int) bool {
	isPow2 := func(num int) bool {
		return num&(num-1) == 0
	}
	ss := []byte(strconv.Itoa(n))
	size := len(ss)
	flag := false
	var permute func(idx int)
	permute = func(idx int) {
		if flag {
			return
		}
		if idx == size {
			if len(ss) > 0 && ss[0] != '0' {
				v, _ := strconv.Atoi(string(ss))
				if isPow2(v) {
					flag = true
					return
				}
			}
			return
		}
		for i := idx; i < size; i++ {
			ss[i], ss[idx] = ss[idx], ss[i]
			permute(idx + 1)
			ss[i], ss[idx] = ss[idx], ss[i]
		}
	}
	permute(0)
	return flag
}

func Test_reorderedPowerOf2(t *testing.T) {
	fmt.Println(reorderedPowerOf2(2))
	fmt.Println(reorderedPowerOf2(4))
	fmt.Println(reorderedPowerOf2(6))
	fmt.Println(reorderedPowerOf2(16))
	fmt.Println(reorderedPowerOf2(64))
	fmt.Println(reorderedPowerOf2(46))
	fmt.Println(reorderedPowerOf2(1234))
}

// tag-[回溯]
// 周赛
// 第一题 leetcode5942: 找出三位偶数
func findEvenNumbers(digits []int) []int {
	sort.Ints(digits)
	// n := len(digits)
	toInt := func(t []int) int {
		out := 0
		factor := 1
		for i := len(t) - 1; i >= 0; i-- {
			out += t[i] * factor
			factor *= 10
		}
		return out
	}

	// 遍历
	var tmp []int
	var out []int
	dp := make([]bool, 1000)
	visited := make([]bool, len(digits))
	var dfs func(idx int)
	dfs = func(idx int) {
		if len(tmp) == 3 {
			v := toInt(tmp)
			if v&1 == 0 && !dp[v] {
				dp[v] = true
				out = append(out, v)
			}
			return
		}
		for i := 0; i < len(digits); i++ {
			if !visited[i] {
				if len(tmp) == 0 && digits[i] == 0 {
					continue
				}
				tmp = append(tmp, digits[i])
				visited[i] = true
				dfs(i + 1)
				visited[i] = false
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0)
	return out
}

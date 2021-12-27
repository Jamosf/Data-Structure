package ojeveryday

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

// tag-[二叉树]
// leetcode124: 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := dfs(node.Left), dfs(node.Right)
		l = max(l, 0)
		r = max(r, 0)
		maxSum = max(maxSum, l+r+node.Val)
		return max(l, r) + node.Val
	}
	dfs(root)
	return maxSum
}

// tag-[二叉树]
// leetcode: 129 求根节点到叶节点数字之和
func sumNumbers(root *TreeNode) int {
	tmp := make([]int, 0, 10)
	ans := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		tmp = append(tmp, node.Val)
		if node.Left == nil && node.Right == nil {
			ans += convertToInt(tmp)
		}
		dfs(node.Left)
		dfs(node.Right)
		tmp = tmp[:len(tmp)-1]
	}
	dfs(root)
	fmt.Println(ans)
	return ans
}

func convertToInt(arr []int) int {
	ans := 0
	for i, c := len(arr)-1, 1; i >= 0; i, c = i-1, c*10 {
		ans += arr[i] * c
	}
	return ans
}

func Test_sumNumber(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 0}
	fmt.Println(sumNumbers(root))
}

// tag-[回溯]
// leetcode301: 删除无效的括号
func removeInvalidParentheses(s string) []string {
	lc, rc := 0, 0
	for _, v := range s {
		if v == '(' {
			lc++
		} else if v == ')' {
			if lc == 0 {
				rc++
			} else {
				lc--
			}
		}
	}
	var ans []string
	m := make(map[string]bool)
	var dfs func(ss string, idx int, l, r int)
	dfs = func(ss string, idx int, l, r int) {
		if l == 0 && r == 0 {
			if !isValid([]byte(ss)) {
				return
			}
			if _, ok := m[ss]; !ok {
				m[ss] = true
				ans = append(ans, ss)
			}
			return
		}
		for i := idx; i < len(ss); i++ {
			if i > idx && ss[i] == ss[i-1] {
				continue
			}
			if l+r > len(ss)-i {
				return
			}
			if l > 0 && ss[i] == '(' {
				dfs(ss[:i]+ss[i+1:], i, l-1, r)
			}
			if r > 0 && ss[i] == ')' {
				dfs(ss[:i]+ss[i+1:], i, l, r-1)
			}
		}
	}
	dfs(s, 0, lc, rc)
	return ans
}

func Test_removeInvalidParentheses(t *testing.T) {
	fmt.Println(removeInvalidParentheses(")("))
}

// tag-[回溯]
// leetcode22: 括号生成
func generateParenthesis_(n int) []string {
	var ans []string
	var dfs func(s string, l, r int)
	dfs = func(s string, l, r int) {
		if l == n && r == n {
			if isValid([]byte(s)) {
				ans = append(ans, s)
			}
			return
		}
		if l < n {
			dfs(s+"(", l+1, r)
		}
		if r < l {
			dfs(s+")", l, r+1)
		}
	}
	dfs("", 0, 0)
	return ans
}

func Test_generateParenthesis_(t *testing.T) {
	fmt.Println(generateParenthesis_(3))
}

// tag-[动态规划]
// leetcode312: 戳气球
// dp[i][j]表示i到j之间的能获得的硬币的最大数量
// dp[i][j] = dp[i][k] + dp[k][j] + pos[i]*pos[k]*pos[j]
func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	pointers := make([]int, n+2)
	pointers[0] = 1
	pointers[n+1] = 1
	for i := 1; i < n+1; i++ {
		pointers[i] = nums[i-1]
	}
	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+pointers[i]*pointers[k]*pointers[j])
			}
		}
	}
	return dp[0][n+1]
}

// tag-[动态规划]
// leetcode10: 字符串匹配
// dp[i][j]表示s的前i个字符能否和p的前j个字符匹配
// dp[i][j] = dp[i-1][j-1](s[i]==s[j] || s[j] == '.')
func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	dp := make([][]bool, ls+1)
	for i := range dp {
		dp[i] = make([]bool, lp+1)
	}
	dp[0][0] = true
	// 初始化首列
	for i := 1; i <= ls; i++ {
		dp[i][0] = false
	}
	// 初始化首行
	for j := 1; j <= lp; j++ {
		if j == 1 || p[j-1] != '*' {
			dp[0][j] = false
		} else {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' { // i, j 相等或j为.
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' { // i和j不相等，并且j为*
				if p[j-2] != s[i-1] && p[j-2] != '.' { // 如果j-1不能匹配
					dp[i][j] = dp[i][j-2]
				} else { // 如果j-1匹配，则分为匹配0次，匹配一次和匹配多次
					dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j]
				}
			}
		}
	}
	return dp[ls][lp]
}

// tag-[栈]
// leetcode394: 字符串解码
func decodeString(s string) string {
	stacks := make([]byte, 0, len(s))
	stackn := make([]int, 0, len(s))
	for i := 0; i < len(s); {
		if s[i] >= '0' && s[i] <= '9' {
			multi := 0
			for s[i] >= '0' && s[i] <= '9' {
				multi = multi*10 + int(s[i]-'0')
				i++
			}
			stackn = append(stackn, multi)
		} else if s[i] != ']' {
			stacks = append(stacks, s[i])
			i++
		} else {
			multi := 0
			if len(stackn) > 0 {
				multi = stackn[len(stackn)-1]
				stackn = stackn[:len(stackn)-1]
			}
			tmp := make([]byte, 0)
			for len(stacks) > 0 && stacks[len(stacks)-1] != '[' {
				tmp = append(tmp, stacks[len(stacks)-1])
				stacks = stacks[:len(stacks)-1]
			}
			if len(stacks) > 0 {
				stacks = stacks[:len(stacks)-1]
			}
			v := make([]byte, 0, multi*len(tmp))
			for i := 0; i < multi; i++ {
				for j := len(tmp) - 1; j >= 0; j-- {
					v = append(v, tmp[j])
				}
			}
			stacks = append(stacks, v...)
			i++
		}
	}
	return string(stacks)
}

// tag-[递归]
// leetcode394: 字符串解码
func decodeString_(s string) string {
	n := len(s)
	var dfs func(idx *int) string
	dfs = func(idx *int) string {
		var res string
		num := 0
		for *idx < n {
			if s[*idx] >= '0' && s[*idx] <= '9' {
				num = num*10 + int(s[*idx]-'0')
			} else if s[*idx] == '[' {
				*idx = *idx + 1
				tmp := dfs(idx)
				for i := 0; i < num; i++ {
					res += tmp
				}
				num = 0
			} else if s[*idx] == ']' {
				break
			} else {
				res += string(s[*idx])
			}
			*idx++
		}
		return res
	}
	idx := 0
	return dfs(&idx)
}

// tag-[字符串]
// leetcode: 字符串解码
type T struct {
	s string // 存储前面的结果
	n int    // 存储后续要重复的数字
}

func decodeString__(s string) string {
	res := ""
	num := 0
	st := make([]T, 0)
	for _, c := range s {
		switch true {
		case c >= '0' && c <= '9':
			num = num*10 + int(c-'0')
		case c == '[':
			st = append(st, T{s: res, n: num})
			res, num = "", 0
		case c == ']':
			wth := st[len(st)-1]
			st = st[:len(st)-1]
			res = wth.s + strings.Repeat(res, wth.n)
		default:
			res = res + string(c)
		}
	}
	return res
}

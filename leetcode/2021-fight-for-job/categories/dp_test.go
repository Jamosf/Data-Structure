package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[动态规划]
// 第四题
// leetcode981: 基于时间的键值存储
type TimeMap struct {
	kv map[string][]v
}

type v struct {
	value     string
	timeStamp int
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	maxn := dp[0]
	for i := range dp {
		maxn = max(dp[i], maxn)
	}
	return maxn
}

/** Initialize your data structure here. */
func ConstructorTimeMap() TimeMap {
	return TimeMap{kv: make(map[string][]v)}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.kv[key] = append(t.kv[key], v{value: value, timeStamp: timestamp})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	if vv, ok := t.kv[key]; ok && len(vv) != 0 {
		for i := len(vv) - 1; i >= 0; i-- {
			if vv[i].timeStamp <= timestamp {
				return vv[i].value
			}
		}
	}
	return ""
}

// tag-[动态规划]
// 第十一题
// leetcode70: 爬楼梯
func climbStairs(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

// tag-[动态规划]
// 第三题
// leetcode198: 打家劫舍
func rob(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else if i == 1 {
			dp[i] = max(nums[0], nums[1])
		} else {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
	}
	return dp[len(nums)-1]
}

// tag-[动态规划]
// 第四题
// leetcode120: 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	minn := math.MaxInt64
	for j := 0; j < len(dp[m-1]); j++ {
		minn = min(minn, dp[m-1][j])
	}
	return minn
}

func Test_minimumTotal(t *testing.T) {
	fmt.Println(minimumTotal([][]int{{-1}, {-2, -3}}))
}

// tag-[动态规划]
// 第三题
// leetcode 剑指offer63: 股票的最大利润
// dp[i]表示以i结尾，最大利润
// dp[i+1] = max(dp[i], dp[i] + nums[i+1]- nums[i])
func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	maxn := dp[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = max(0, dp[i-1]+prices[i]-prices[i-1])
		maxn = max(dp[i], maxn)
	}
	return maxn
}

// tag-[动态规划]
// leetcode413: 等差数列划分
// dp[i]表示以i结尾的等差数列的个数
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	dp := make([]int, len(nums))
	if nums[2]-nums[1] == nums[1]-nums[0] {
		dp[2] = 1
	}
	for i := 3; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	ans := 0
	for i := range dp {
		ans += dp[i]
	}
	return ans
}

func Test_numberOfArithmeticSlices(t *testing.T) {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 0, 5, 6, 7}))
}

// tag-[动态规划]
// 第二题
// leetcode542: 01矩阵
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if j < n-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
				if i < m-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
			}
		}
	}
	return dp
}

// tag-[动态规划]
// 第三题
// leetcode221: 最大正方形
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	maxn := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != '0' {
				if i == 0 || j == 0 {
					dp[i][j] = int(matrix[i][j] - '0')
				} else {
					dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				}
				maxn = max(maxn, dp[i][j])
			}
		}
	}
	return maxn * maxn
}

func Test_maximalSquare(t *testing.T) {
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}

// tag-[动态规划]
// 第四题
// leetcode91: 解码方法
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	for i := 1; i < n; i++ {
		if s[i] > '0' && s[i] <= '9' {
			dp[i] += dp[i-1]
		}
		if s[i-1:i+1] >= "10" && s[i-1:i+1] <= "26" {
			if i >= 2 {
				dp[i] += dp[i-2]
			} else {
				dp[i] += 1
			}
		}
	}
	return dp[n-1]
}

// tag-[动态规划]
// leetcode91: 解码方法
// 空间压缩版，内存消耗略优于上面的版本
func numDecodings_(s string) int {
	n := len(s)
	dp := make([]int, 3)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i%3] = 0
		if s[i] > '0' && s[i] <= '9' {
			dp[i%3] += dp[(i-1)%3]
		}
		if s[i-1:i+1] >= "10" && s[i-1:i+1] <= "26" {
			if i >= 2 {
				dp[i%3] += dp[(i-2)%3]
			} else {
				dp[i%3] += 1
			}
		}
	}
	return dp[(n-1)%3]
}

func Test_numDecodings(t *testing.T) {
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings_("226"))
}

// tag-[动态规划]
// 第五题
// leetcode139：单词拆分
func wordBreak139(s string, wordDict []string) bool {
	s = " " + s
	n := len(s)
	dp := make([]bool, n)
	dp[0] = true
	for i := 1; i < n; i++ {
		for j := 0; j < len(wordDict); j++ {
			if i >= len(wordDict[j]) && wordDict[j] == s[i+1-len(wordDict[j]):i+1] {
				dp[i] = dp[i] || dp[i-len(wordDict[j])]
			}
		}
	}
	return dp[n-1]
}

// tag-[动态规划]
// 第六题
// leetcode300
// dp[i]表示以i结尾的最长子序列
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	maxn := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxn = max(maxn, dp[i])
	}
	return maxn
}

// 优化版解法：贪心+二分查找
func lengthOfLIS_(nums []int) int {
	n := len(nums)
	tail := make([]int, 0)
	tail = append(tail, nums[0])
	for i := 1; i < n; i++ {
		if nums[i] > tail[len(tail)-1] {
			tail = append(tail, nums[i])
		} else {
			idx := sort.SearchInts(tail, nums[i])
			tail[idx] = nums[i]
		}
	}
	return len(tail)
}

func Test_lengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	fmt.Println(lengthOfLIS_([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}

// tag-[动态规划]
// leetcode1986: 完成任务的最少工作时间段
// dp[i]的含义是：i表示的二进制对应的索引全部选中的情况下，最小的组合数
func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	m := 1 << n
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 20 // 以普遍理论而论，状态压缩的数据范围会小于20
	}
	// 1. 计算哪些子集可以在sessionTime内完成，标记为1
	for i := 0; i < m; i++ {
		state, idx := i, 0
		spend := 0
		for state > 0 {
			if state&1 == 1 {
				spend += tasks[idx]
			}
			state >>= 1
			idx++
		}
		if spend <= sessionTime {
			dp[i] = 1
		}
	}
	// 2. 从第一步计算出子集推到全集的最小值，枚举二进制子集并完成状态转移。
	// 一个111的子集可以由110 + 1构成。即i = 111， j = 110， i^j = 1。i^j，有等于i-j的语义。
	// 另外还有一个更新j的技巧：j = (j-1) &i， j: 110->101->100->011->010->001。这样会出现一个重复计算的问题，因为j减少的时候，会等于i^j。
	// 原则是，j与i^j一定是i的子集。
	for i := 1; i < m; i++ {
		if dp[i] == 1 {
			continue
		}
		for j := i; j > 0; j = (j - 1) & i {
			dp[i] = min(dp[i], dp[j]+dp[i^j])
		}
	}
	return dp[m-1]
}

// tag-[动态规划]
// leetcode1987: 不同的好子序列数目
func numberOfUniqueGoodSubsequences(binary string) int {
	n := len(binary)
	dp0, dp1 := 0, 0
	has0, mod := 0, int(1e9+7)
	for i := n - 1; i >= 0; i-- {
		if binary[i] == '0' {
			has0 = 1
			dp0 = (dp0 + dp1 + 1) % mod
		} else {
			dp1 = (dp0 + dp1 + 1) % mod
		}
	}
	return (dp1 + has0) % mod
}

// tag-[动态规划]
// leetcode1997: 访问完所有房间的第一天
// 超时了
func firstDayBeenInAllRooms(nextVisit []int) int {
	m := make(map[int]int)
	n := len(nextVisit)
	visited := 0
	current := 0
	ans := 0
	for visited < n {
		if _, ok := m[current]; !ok {
			visited++
		}
		ans++
		m[current]++
		if m[current]&1 == 1 {
			current = nextVisit[current]
		} else {
			current = (current + 1) % n
		}
	}
	return ans
}

// leetcode1997: 访问完所有房间的第一天
// dp求解, +mod 可以防止出现负数
func firstDayBeenInAllRooms_(nextVisit []int) int {
	mod := int64(1e9 + 7)
	n := len(nextVisit)
	dp := make([]int64, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = (dp[i-1]*2 - dp[nextVisit[i-1]] + 2 + mod) % mod
	}
	return int((dp[n-1]) % mod)
}

func Test_firstDayBeenInAllRooms1(t *testing.T) {
	fmt.Println(firstDayBeenInAllRooms([]int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 23, 24, 24, 25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47, 47, 48}))
	fmt.Println(firstDayBeenInAllRooms_([]int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 23, 24, 24, 25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47, 47, 48}))
}

// tag-[动态规划]
// leetcode1959: k次调整数组大小浪费的最小总空间
// 区间dp，dp[x][y]表示前i个元素被分成j段的最小值
// 状态转移方程：dp[x][y] = min(dp[x][y], dp[l-1][y-1] + weight[l][x])，其中l范围0...x-1
func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	weight := make([][]int, n)
	for i := range weight {
		weight[i] = make([]int, n)
	}
	// 1. 预处理权值
	for i := 0; i < n; i++ {
		sum := 0
		maxn := nums[i]
		for j := i; j < n; j++ {
			sum += nums[j]
			maxn = max(maxn, nums[j])
			weight[i][j] = maxn*(j-i+1) - sum
		}
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+2)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	// 2. dp计算
	for i := 0; i < n; i++ {
		for j := 1; j < k+2; j++ {
			for l := 0; l <= i; l++ {
				if l == 0 {
					dp[i][j] = min(dp[i][j], 0+weight[0][i])
				} else {
					dp[i][j] = min(dp[i][j], dp[l-1][j-1]+weight[l][i])
				}
			}
		}
	}
	return dp[n-1][k+1]
}
// tag-[动态规划]
// leetcode2008: 出租车的最大盈利
func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool {
		a, b := rides[i], rides[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	m := len(rides)
	dp := make([]int64, m)
	dp[0] = int64(rides[0][1] - rides[0][0] + rides[0][2])
	maxn := dp[0]
	for i := 1; i < m; i++ {
		v := int64(rides[i][1] - rides[i][0] + rides[i][2])
		flag := false
		for j := i - 1; j >= 0; j-- {
			if rides[i][0] >= rides[j][1] {
				flag = true
				dp[i] = maxInt64(dp[i], dp[j]+v)
			}
		}
		if !flag {
			dp[i] = v
		}
		maxn = maxInt64(maxn, dp[i])
	}
	return maxn
}

// leetcode2008: 出租车的最大盈利
func maxTaxiEarnings_(n int, rides [][]int) int64 {
	g := make([][][2]int, n+1)
	for _, v := range rides {
		start, end, trips := v[0], v[1], v[2]
		g[end] = append(g[end], [2]int{start, trips})
	}
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[i-1]
		for _, e := range g[i] {
			f[i] = max(f[i], f[e[0]]+i-e[0]+e[1])
		}
	}
	return int64(f[n])
}

func Test_maxTaxiEarnings(t *testing.T) {
	fmt.Println(maxTaxiEarnings(10, [][]int{{9, 10, 2}, {4, 5, 6}, {6, 8, 1}, {1, 5, 5}, {4, 9, 5}, {1, 6, 5}, {4, 8, 3}, {4, 7, 10}, {1, 9, 8}, {2, 3, 5}}))
	fmt.Println(maxTaxiEarnings_(10, [][]int{{9, 10, 2}, {4, 5, 6}, {6, 8, 1}, {1, 5, 5}, {4, 9, 5}, {1, 6, 5}, {4, 8, 3}, {4, 7, 10}, {1, 9, 8}, {2, 3, 5}}))
}
// tag-[动态规划]
// 第六题
// leetcode55: 跳跃游戏
func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				dp[i] = dp[i] || dp[j]
			}
			if dp[i] {
				break
			}
		}
		if !dp[i] {
			return false
		}
	}
	return dp[n-1]
}
// tag-[动态规划]
// leetcode32: 最长有效括号
// 动态规划解法
func longestValidParentheses_(s string) int {
	n := len(s)
	dp := make([]int, n) // 表示以）结尾的最长的子串长度
	dp[0] = 0
	maxn := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-2*dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 1
			}
		}
		// if s[i] == '(' {
		// 	dp[i] = dp[i-1]
		// }
		maxn = max(maxn, dp[i])
	}
	return 2 * maxn
}

func Test_longestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses("()()(((()())))"))
	fmt.Println(longestValidParentheses_("()()(((()())))"))
}

// tag-[动态规划]
// leetcode152: 乘积最大的子数组
// 动态规划
func maxProduct152(nums []int) int {
	n := len(nums)
	dp := make([][2]int, n) // 以i结尾的最大值或最小值，
	dp[0][0] = nums[0]      // 0存最小值
	dp[0][1] = nums[0]      // 1存最大值
	maxn := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] >= 0 {
			dp[i][0] = min(dp[i-1][0]*nums[i], nums[i])
			dp[i][1] = max(dp[i-1][1]*nums[i], nums[i])
		} else {
			dp[i][0] = min(dp[i-1][1]*nums[i], nums[i])
			dp[i][1] = max(dp[i-1][0]*nums[i], nums[i])
		}
		maxn = max(maxn, dp[i][1])
	}
	return maxn
}

func Test_maxProduct1(t *testing.T) {
	fmt.Println(maxProduct152([]int{-3, 2, -4}))
}

// tag-[动态规划]
// leetcode 面试题17.16: 按摩师
// 动态规划
func massage(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	if n < 2 {
		return dp[0]
	}
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		for j := i - 2; j >= 0; j-- {
			dp[i] = max(dp[i], dp[j]+nums[i])
		}
	}
	return max(dp[n-1], dp[n-2])
}

// tag-[动态规划]
// leetcode337: 打家劫舍
// 树形dp
func rob1(root *TreeNode) int {
	var dfs func(r *TreeNode) [2]int
	dfs = func(r *TreeNode) [2]int {
		if r == nil {
			return [2]int{}
		}
		L := dfs(r.Left)
		R := dfs(r.Right)

		dp := [2]int{} // 0表示不选父节点、1表示选择父节点
		dp[0] = max(L[0], L[1]) + max(R[0], R[1])
		dp[1] = r.Val + L[0] + R[0]
		return dp
	}
	ans := dfs(root)
	return max(ans[0], ans[1])
}

// tag-[动态规划]
// leetcode62: 不同路径
// 动态规划
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1]
			}
			if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j]
			}
			if i > 0 && j > 0 {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

// tag-[动态规划]
// leetcode96: 不同的二叉搜索树
// 每个数字作为根节点，左右子树的各种组合乘积
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
// tag-[动态规划]
// 第五题
// leetcode309: 最佳买卖股票时机含冷冻期
// 动态规划：dp[i]表示第i天获取的最大利润, 0：持有一只股票；1：不持有股票，处于冷冻期；2：不持有股票，不处于冷冻期
func maxProfit1(prices []int) int {
	n := len(prices)
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

// tag-[动态规划]
// 第一题
// O(n^2)解法
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func Test_jump(t *testing.T) {
	fmt.Println(jump1([]int{1, 1, 1, 4, 1, 1, 1}))
}

// 第二题
// O(n)解法
func jump1(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

// tag-[动态规划]
// leetcode673: 可以用双dp来理解
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxn := 1
	count := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		if dp[i] > maxn {
			maxn = dp[i]
			ans = count[i]
		} else if dp[i] == maxn {
			ans += count[i]
		}
	}
	return ans
}

func Test_findNumberOfLIS(t *testing.T) {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))
}
// tag-[动态规划]
// leetcode97:交错字符串(记忆化搜索)
func isInterleave_(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n3 != n1+n2 {
		return false
	}
	dp := make([][]bool, n1+1) // 记忆化，存储中间过程数据
	for i := range dp {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s1[i-1] == s3[p] // i-1对应数组的索引是第i个，和dp刚好差一
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[n1][n2]
}

func Test_isInterleave(t *testing.T) {
	fmt.Println(isInterleave_("aabcc", "dbbca", "aadbbcbcac"))
}
// tag-[动态规划]
// leetcode63:动态规划(可以降维)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}

// tag-[动态规划]
// leetcode375: 猜数字大小II
// 区间dp状态转移方程: f(i, j) = min{k + max(f(i, k-1), f(k+1, j))}  1<=k<=j
func getMoneyAmount_(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = j + dp[i][j-1]
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], k+max(dp[i][k-1], dp[k+1][j]))
			}
		}
	}
	return dp[1][n]
}

// tag-[动态规划]
// leetcode2086: 从房屋收集雨水需要的最少水桶数
// dp[i][3]: 第二维的含义：0表示以H结尾且前面无桶、1表示以H结尾且前面有桶、2表示以B结尾、3表示以.结尾（前面H必有桶）
// 状态转移:
func minimumBuckets_(street string) int {
	n := len(street)
	dp := make([][4]int, n)
	if street[0] == '.' {
		dp[0][0] = math.MaxInt32
		dp[0][1] = math.MaxInt32
		dp[0][2] = 1
		dp[0][3] = 0
	} else {
		dp[0][0] = 0
		dp[0][1] = math.MaxInt32
		dp[0][2] = math.MaxInt32
		dp[0][3] = math.MaxInt32
	}
	for i := 1; i < n; i++ {
		if street[i] == '.' {
			dp[i][3] = min(min(dp[i-1][1], dp[i-1][2]), dp[i-1][3])
			dp[i][2] = dp[i-1][0] + 1
			dp[i][1] = math.MaxInt32
			dp[i][0] = math.MaxInt32
		} else {
			dp[i][0] = min(dp[i-1][1], dp[i-1][3])
			dp[i][1] = min(dp[i-1][2], dp[i-1][3]+1)
			dp[i][2] = math.MaxInt32
			dp[i][3] = math.MaxInt32
		}
	}
	ans := math.MaxInt32
	ans = min(ans, dp[n-1][1])
	ans = min(ans, dp[n-1][2])
	ans = min(ans, dp[n-1][3])
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// tag-[矩阵]
// leetcode2087: 网格图中机器人回家的最小代价
// 直接走
func minCost(startPos, homePos, rowCosts, colCosts []int) int {
	x0, y0, x1, y1 := startPos[0], startPos[1], homePos[0], homePos[1]
	ans := -rowCosts[x0] - colCosts[y0] // 初始的行列无需计算
	if x0 > x1 {
		x0, x1 = x1, x0
	} // 交换位置，保证 x0 <= x1
	if y0 > y1 {
		y0, y1 = y1, y0
	} // 交换位置，保证 y0 <= y1
	for _, cost := range rowCosts[x0 : x1+1] {
		ans += cost
	} // 统计答案
	for _, cost := range colCosts[y0 : y1+1] {
		ans += cost
	} // 统计答案
	return ans
}

// tag-[动态规划]
// leetcode2063: 所有子字符串中的元音
// dp[i] = dp[i-1] + ？
// 如果第i位是元音，那么dp[i]需要在dp[i-1]的基础上，加上所有以i结尾的元音个数。
// 元音个数的计算采用累加的方式，如果发现第k个字符为元音，那么前0...k个数字将可以利用第K个元音，因此这个区间的所有位置能与字符尾组成的元音子串的个数需要加K+1。
func countVowels(word string) int64 {
	n := len(word)
	dp := make([]int64, n)
	vowel := []byte{'a', 'e', 'i', 'o', 'u'}
	cnt := int64(0)
	if isVowel(word[0], vowel) {
		dp[0] = 1
		cnt++
	}

	for i := 1; i < n; i++ {
		if isVowel(word[i], vowel) {
			cnt += int64(i + 1)
			dp[i] = dp[i-1] + cnt
		} else {
			dp[i] = dp[i-1] + cnt
		}
	}
	return dp[n-1]
}

func isVowel(b byte, vowel []byte) bool {
	for i := range vowel {
		if vowel[i] == b {
			return true
		}
	}
	return false
}

// tag-[动态规划]
// leetcode1235: 规划兼职工作
// 上面的解法会占用内存过大而出现内存分配失败的情况
// fn[i]表示0~i内最多选择k个的最大值。
// fn[i] = max(fn[i-1], fn[high]+profit[i])
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})
	fn := make([]int, n)
	for i := 0; i < n; i++ {
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) >> 1
			if jobs[mid][1] <= jobs[i][0] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fn[i] = jobs[i][2]
		if high >= 0 {
			fn[i] += fn[high]
		}
		if i > 0 {
			fn[i] = max(fn[i], fn[i-1])
		}
	}
	return fn[n-1]
}

// tag-[动态规划/堆]
// leetcode2054: 两个最好的不重叠活动
// 按照开始时间进行排序，同时用小根堆维护结束时间的队列，结束必须是小根堆。
func maxTwoEvents(events [][]int) (ans int) {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	maxn := 0
	h := hp_{}
	for i := 0; i < len(events); i++ {
		start, end, val := events[i][0], events[i][1], events[i][2]
		for len(h) > 0 && h[0].end < start {
			maxn = max(maxn, heap.Pop(&h).(pair_).val)
		}
		ans = max(ans, maxn+val)
		heap.Push(&h, pair_{end, val})
	}
	return ans
}

type pair_ struct{ end, val int }
type hp_ []pair_

func (h hp_) Len() int            { return len(h) }
func (h hp_) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp_) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp_) Push(v interface{}) { *h = append(*h, v.(pair_)) }
func (h *hp_) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// tag-[动态规划/二分]
// leetcode2054: 两个最好的不重叠活动
// 排序+动态规划+二分
// fn表示0~i选一个的最大值：fn[i] = max(fn[i-1], event[i][2])，可以使用前缀最大值来代替
// gn表示0~i选两个的最大值：gn[i] = max(gn[i-1], f[j]+event[i][2])
func maxTwoEvents_(events [][]int) (ans int) {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	fn := make([]int, n)
	gn := make([]int, n)
	for i := 0; i < n; i++ {
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) >> 1
			if events[mid][1] < events[i][0] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fn[i] = events[i][2]
		if high >= 0 {
			gn[i] = fn[high] + events[i][2]
		}
		if i > 0 {
			fn[i] = max(fn[i], fn[i-1])
			gn[i] = max(gn[i], gn[i-1])
		}
	}
	return max(fn[n-1], gn[n-1])
}

// tag-[动态规划/二分]
// leetcode2054: 两个最好的不重叠活动
// 排序+动态规划+二分
// fn表示0~i选两个的最大值：fn[i] = max(fn[i-1], preMax[high]+event[i][2])
func maxTwoEvents__(events [][]int) (ans int) {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	fn := make([]int, n)
	preMax := make([]int, n)
	for i := 0; i < n; i++ {
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) >> 1
			if events[mid][1] < events[i][0] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fn[i] = events[i][2]
		preMax[i] = events[i][2]
		if high >= 0 {
			fn[i] += preMax[high]
		}
		if i > 0 {
			fn[i] = max(fn[i], fn[i-1])
			preMax[i] = max(preMax[i], preMax[i-1])
		}
	}
	return fn[n-1]
}

package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

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

// leetcode464: 我能赢吗
// 记忆化搜索，博弈
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	dp := make([]int, 1<<maxChoosableInteger)
	var dfs func(state, remain int) bool
	dfs = func(state, remain int) bool {
		if dp[state] != 0 {
			return dp[state] == 1
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			if (1<<(i-1))&state != 0 {
				continue
			}
			if i >= remain || !dfs((1<<(i-1))|state, remain-i) {
				dp[state] = 1
				return true
			}
		}
		dp[state] = -1
		return false
	}
	if maxChoosableInteger > desiredTotal {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	return dfs(0, desiredTotal)
}

func Test_canIWin(t *testing.T) {
	fmt.Println(canIWin(3, 5))
	// fmt.Println(canIWin(10, 11))
}

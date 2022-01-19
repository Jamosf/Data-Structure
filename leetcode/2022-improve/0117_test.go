package _022_improve

import (
	"fmt"
	"testing"
)

// tag-[深度优先搜索]
// leetcode1220: 统计元音字母序列的数目
func countVowelPermutation(n int) int {
	m := map[byte][]byte{}
	m['a'] = []byte{'e'}
	m['e'] = []byte{'a', 'i'}
	m['i'] = []byte{'a', 'e', 'o', 'u'}
	m['o'] = []byte{'i', 'u'}
	m['u'] = []byte{'a'}
	mod := int(1e9+7)
	memo := make(map[byte]map[int]int)
	for k, _ := range m{
		memo[k] = make(map[int]int)
	}
	var dfs func(v byte, idx int) int
	dfs = func(v byte, idx int) int{
		if idx == n {
			return 1
		}
		if memo[v][idx] != 0{
			return memo[v][idx]
		}
		next := m[v]
		sum := 0
		for i := range next{
			sum += (dfs(next[i], idx+1))%mod
		}
		memo[v][idx] = sum%mod
		return sum%mod
	}
	ans := 0
	for k, _ := range m{
		ans += (dfs(k, 1))%mod
	}
	return ans%mod
}

func Test_countVowelPermutation(t *testing.T){
	fmt.Println(countVowelPermutation(10000))
}

// tag-[动态规划]
// leetcode1220: 统计元音字母序列的数目
// dp[i][j]表示长度为i结尾，以j结尾的字符串数目，以0、1、2、3、4表示a,e,i,o,u
// dp[i][0] = dp[i-1][1] + dp[i-1][2] + dp[i-1][3]
// dp[i][1] = dp[i-1][0] + dp[i-1][2]
// ...
// 由于状态至于上一个状态有关，可以进行状态压缩
func countVowelPermutation_(n int) int {
	dp := [5]int{1,1,1,1,1}
	mod := int(1e9+7)
	for i := 2; i < n; i++{
		dp = [5]int{
			(dp[1]+dp[2]+dp[4])%mod,
			(dp[0]+dp[2])%mod,
			(dp[1]+dp[3])%mod,
			dp[2],
			(dp[2]+dp[3])%mod,
		}
	}
	ans := 0
	for _, v := range dp{
		ans = (ans+v)%mod
	}
	return ans
}

// tag-[矩阵快速幂]
// leetcode1220: 统计元音字母序列的数目
const mod int = 1e9+7

type matrix [5][5]int

func (a matrix) mul(b matrix)matrix{
	c := matrix{}
	for i, row := range a{
		for j := 0; j < len(b[0]); j++{
			for k, v := range row{
				c[i][j] = (c[i][j] + v*b[k][j])%mod
			}
		}
	}
	return c
}

func (a matrix) pow(n int) matrix{
	res := matrix{}
	for i := range res{
		res[i][i] = 1
	}
	for ; n > 0; n >>= 1{
		if n & 1 > 0{
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

func countVowelPermutation__(n int) int {
	m := matrix{
		{0, 1, 0, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0},
	}

	res := m.pow(n-1)
	ans := 0
	for _, row := range res{
		for _, v := range row{
			ans = (ans+v)%mod
		}
	}
	return ans
}

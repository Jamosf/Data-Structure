package acwing

import "fmt"

func input0119()(int, int, [][]int){
	m, n := 0, 0
	fmt.Scanf("%d %d", &m, &n)
	res := make([][]int, m)
	for i := 0; i < m; i++{
		v, w := 0, 0
		fmt.Scanf("%d %d", &v, &w)
		res[i] = []int{v, w}
	}
	return m, n, res
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}

// 01背包，dp[i] = max(dp[i], dp[i-v[i]] + w[i])
func main() {
	m, n, res := input0119()
	dp := make([]int, n+1)
	for i := 0; i <= m; i++{
		for j := n; j >= res[i][0]; j--{
			dp[j] = max(dp[j], dp[j-res[i][0]]+res[i][1])
		}
	}
	fmt.Println(dp[n])
}

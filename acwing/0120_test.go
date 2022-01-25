package acwing

import "fmt"

func main(){
	m, n, res := input0119()
	dp := make([]int, n+1)
	for i := 0; i < m; i++{
		for j := res[i][0]; j <= n ; j++{
			dp[j] = max(dp[j], dp[j-res[i][0]]+res[i][1])
		}
	}
	fmt.Println(dp[n])
}

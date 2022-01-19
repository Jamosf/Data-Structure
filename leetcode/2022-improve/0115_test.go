package _022_improve

// tag-数学
// leetcode1716: 计算力扣银行的钱
func totalMoney(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		v := i / 7
		sum += v + i%7 + 1
	}
	return sum
}

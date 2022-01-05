package _022_improve

// tag-[字符串]
// leetcode1576. 替换所有的问号
func modifyString(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == '?' {
			for j := byte('a'); j < byte('z'); j++ {
				if (i == 0 || i > 0 && b[i-1] != j) && (i == len(b)-1 || i < len(b)-1 && b[i+1] != j) {
					b[i] = j
					break
				}
			}
		}
	}
	return string(b)
}

// tag-[动态规划]
// leetcode446: 等差数列划分II - 子序列
// dp[i][d]表示以i结尾，公差为d的等差数列
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	ans := 0
	f := make([]map[int]int, n)
	for i, x := range nums {
		f[i] = map[int]int{}
		for j, y := range nums[:i] {
			d := x - y
			cnt := f[j][d]
			ans += cnt
			f[i][d] += cnt + 1
		}
	}
	return ans
}

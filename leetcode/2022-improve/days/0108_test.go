package days

// tag-[二进制]
// leetcode89: 格雷编码
func grayCode(n int) []int {
	head := 1
	ans := []int{0}
	for i := 0; i < n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, head+ans[j])
		}
		head <<= 1
	}
	return ans
}

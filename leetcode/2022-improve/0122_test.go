package _022_improve

// tag-[字符串]
// leetcode1332: 删除回文子序列
func removePalindromeSub(s string) int {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return 2
		}
	}
	return 1
}

package days

// tag-[数学]
// 每日一题
// leetcode1688: 比赛中的配对次数
func numberOfMatches(n int) int {
	cnt := 0
	for n != 1 {
		if n&1 == 1 {
			n = (n-1)/2 + 1
			cnt += n - 1
		} else {
			n = n / 2
			cnt += n
		}
	}
	return cnt
}

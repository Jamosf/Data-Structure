package days

// tag-[数组]
// leetcode1629: 按键持续时间最长的键
func slowestKey(releaseTimes []int, keysPressed string) byte {
	n := len(releaseTimes)
	tmp := append([]int{0}, releaseTimes...)
	var count [26]int
	for i := 1; i <= n; i++ {
		diff := tmp[i] - tmp[i-1]
		v := keysPressed[i-1] - 'a'
		count[v] = max(count[v], diff)
	}
	maxn := count[0]
	ans := byte('a')
	for i := range count {
		if count[i] >= maxn {
			ans = byte(i + 'a')
			maxn = count[i]
		}
	}
	return ans
}

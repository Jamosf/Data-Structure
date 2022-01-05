package ojeveryday

// tag-[广度优先搜索]
// leetcode752: 打开转盘锁
func openLock(deadends []string, target string) int {
	m := make(map[string]bool)
	for i := range deadends {
		m[deadends[i]] = true
	}
	if target == "0000" {
		return 0
	}
	var q []string
	if m["0000"] {
		return -1
	}
	q = append(q, "0000")
	var step int
	vis := make(map[string]bool)
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			v := q[0]
			q = q[1:]
			if v == target {
				return step
			}
			for _, ss := range plus(v) {
				if m[ss] || vis[ss] {
					continue
				}
				vis[ss] = true
				q = append(q, ss)
			}
			for _, ss := range minus(v) {
				if m[ss] || vis[ss] {
					continue
				}
				vis[ss] = true
				q = append(q, ss)
			}
		}
		step++
	}
	return -1
}

func plus(s string) []string {
	var ans []string
	for i := range s {
		b := []byte(s)
		if s[i] == '9' {
			b[i] = '0'
		} else {
			b[i] = b[i] + 1
		}
		ans = append(ans, string(b))
	}
	return ans
}

func minus(s string) []string {
	var ans []string
	for i := range s {
		b := []byte(s)
		if s[i] == '0' {
			b[i] = '9'
		} else {
			b[i] = b[i] - 1
		}
		ans = append(ans, string(b))
	}
	return ans
}

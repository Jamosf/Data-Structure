package _022_improve

// tag-[广度优先搜索]
// leetcode1345: 跳跃游戏
func minJumps(arr []int) int {
	n := len(arr)
	m := make(map[int][]int)
	for i, v := range arr {
		m[v] = append(m[v], i)
	}
	q := make([][2]int, 0)
	q = append(q, [2]int{0, 0})
	vis := make([]bool, n)
	vis[0] = true
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		i, step := v[0], v[1]
		if i == n-1 {
			return step
		}
		if i+1 < n && !vis[i+1] {
			q = append(q, [2]int{i + 1, step + 1})
			vis[i+1] = true
		}
		if i-1 > 0 && !vis[i-1] {
			q = append(q, [2]int{i - 1, step + 1})
			vis[i-1] = true
		}
		for _, j := range m[arr[i]] {
			if !vis[j] {
				q = append(q, [2]int{j, step + 1})
				vis[j] = true
			}
		}
		delete(m, arr[i])
	}
	return -1
}

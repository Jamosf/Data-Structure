package _022_improve

// tag-[博弈]
// 每日一题
// leetcode2029: 石子游戏IX
// 博弈
func stoneGameIX(stones []int) bool {
	check := func(c [3]int) bool{
		if c[1] == 0{
			return false
		}
		c[1]--
		turn := 1 + min(c[1], c[2])*2 + c[0]
		if c[1] > c[2]{
			turn++
			c[1]--
		}
		return turn%2 == 1 && c[1] != c[2]
	}
	cnt := [3]int{}
	for _, v:= range stones{
		cnt[v%3]++
	}
	return check(cnt) || check([3]int{cnt[0], cnt[2], cnt[1]})
}

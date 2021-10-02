package leetcode

func trap(height []int) int {
	n := len(height)
	maxn := height[0]
	for i := 1; i < n; i++ {
		maxn = max(maxn, height[i])
	}
	sum := 0
	for i := 0; i < maxn; i++ {
		hasHighInLeft := false
		tmp := 0
		for j := 0; j < n; j++ {
			if hasHighInLeft && height[j] < i {
				tmp++
			}
			if height[j] >= i {
				sum += tmp
				tmp = 0
				hasHighInLeft = true
			}
		}
	}
	return sum
}

func rotate(matrix [][]int) {

}

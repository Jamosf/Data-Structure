package ojeveryday

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j, pivot := l-1, r+1, nums[(l+r)>>1]
	for i < j {
		j--
		for nums[j] > pivot {
			j--
		}
		i++
		for nums[i] < pivot {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}

// leetcode912: 排序数组
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// leetcode1314: 矩阵区域和
// 二维前缀和
// sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + mat[i-1][j-1]
// 区域和: sum[x2][y2] - sum[x1-1][y2] - sum[x2][y1-1] + sum[x1-1][y1-1]
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + mat[i-1][j-1]
		}
	}
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x1, y1, x2, y2 := i-k, j-k, i+k, j+k
			if x1 < 0 {
				x1 = 0
			}
			if y1 < 0 {
				y1 = 0
			}
			if x2 > m-1 {
				x2 = m - 1
			}
			if y2 > n-1 {
				y2 = n - 1
			}
			ans[i][j] = sum[x2+1][y2+1] - sum[x1][y2+1] - sum[x2+1][y1] + sum[x1][y1]
		}
	}
	return ans
}
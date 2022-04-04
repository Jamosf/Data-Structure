package days

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// tag
// ����
// ģ��
func countOperations(num1 int, num2 int) int {
	cnt := 0
	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			num1 = num1 - num2
		} else {
			num2 = num2 - num1
		}
		cnt++
	}
	return cnt
}

// ����
func minimumOperations(nums []int) int {
	n := len(nums)
	m, k := make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		if i&1 == 1 {
			m[nums[i]]++
		} else {
			k[nums[i]]++
		}
	}
	maxj, maxo := [][2]int{}, [][2]int{}
	for i, v := range m {
		maxj = append(maxj, [2]int{i, v})
	}
	for i, v := range k {
		maxo = append(maxo, [2]int{i, v})
	}
	sort.Slice(maxj, func(i, j int) bool {
		return maxj[i][1] > maxj[j][1]
	})
	sort.Slice(maxo, func(i, j int) bool {
		return maxo[i][1] > maxo[j][1]
	})
	if maxo[0][0] == maxj[0][0] {
		if len(maxo) == 1 && len(maxj) > 1 {
			return n - maxj[1][1] - maxo[0][0]
		}
		if len(maxj) == 1 && len(maxo) > 1 {
			return n - maxo[1][1] - maxo[0][0]
		}
		return n - (max(maxo[1][1], maxj[1][1]) + maxo[0][0])
	}
	return n - maxj[0][1] - maxo[0][1]
}

// ����
func minimumRemoval(beans []int) int64 {
	n := len(beans)
	sort.Ints(beans)
	preSum := make([]int, n+1)
	for i := range beans {
		preSum[i+1] = preSum[i] + beans[i]
	}
	count := int64(math.MaxInt64)
	for j := range beans {
		idx := search(beans, beans[j])
		r := int64(preSum[idx+1] + preSum[n] - preSum[j] - beans[j]*(n-j))
		count = minInt64(count, r)
	}
	return count
}

func search(a []int, x int) int {
	l, r := 0, len(a)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if a[mid] >= x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	if a[l] >= x {
		return -1
	}
	return l
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Test_minimumRemoval(t *testing.T) {
	fmt.Println(minimumRemoval([]int{1, 4, 5, 6}))
}

// ÿ��һ��
func maxNumberOfBalloons(text string) int {
	char := [5]uint8{'b', 'a', 'l', 'o', 'n'}
	m := make(map[uint8]int)
	for i := range char {
		m[char[i]] = 0
	}
	for i := range text {
		for j := range char {
			if char[j] == text[i] {
				m[char[j]]++
			}
		}
	}
	minn := math.MaxInt32
	for k, v := range m {
		if k == 'l' || k == 'o' {
			minn = min(minn, v/2)
		} else {
			minn = min(minn, v)
		}
	}
	return minn
}

func Test_maxNumber(t *testing.T) {
	fmt.Println(maxNumberOfBalloons("balon"))
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
}

// ÿ��һ��
func singleNonDuplicate(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

func Test_single(t *testing.T) {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3, 3, 4, 5, 5}))
	fmt.Println(singleNonDuplicate([]int{1}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 3, 2, 2}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}

// ÿ��һ��
func luckyNumbers(matrix [][]int) []int {
	getMinIdx := func(a []int) int {
		minn, idx := a[0], 0
		for i := 1; i < len(a); i++ {
			if minn > a[i] {
				minn = a[i]
				idx = i
			}
		}
		return idx
	}
	var ans []int
	for i := 0; i < len(matrix); i++ {
		j := getMinIdx(matrix[i])
		for up := i - 1; up >= 0; up-- {
			if matrix[i][j] <= matrix[up][j] {
				goto end
			}
		}
		for down := i + 1; down < len(matrix); down++ {
			if matrix[i][j] <= matrix[down][j] {
				goto end
			}
		}
		ans = append(ans, matrix[i][j])
	end:
	}
	return ans
}

// ÿ��һ��
// ��ʱ
func knightProbability_(n int, k int, row int, column int) float64 {
	direct := [8][2]int{{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}}
	q := [][2]int{}
	q = append(q, [2]int{row, column})
	out, kk := 0, k
	for len(q) > 0 && k > 0 {
		size := len(q)
		for s := 0; s < size; s++ {
			v := q[0]
			q = q[1:]
			for i := 0; i < len(direct); i++ {
				x, y := v[0]+direct[i][0], v[1]+direct[i][1]
				if x < 0 || x >= n || y < 0 || y >= n {
					out++
					continue
				}
				q = append(q, [2]int{x, y})
			}
		}
		k--
	}
	return float64(len(q)) / math.Pow(8, float64(kk))
}

func Test_knightProbability(t *testing.T) {
	fmt.Println(knightProbability(3, 2, 0, 0))
	fmt.Println(knightProbability(25, 100, 0, 0))
}

// dp
var dirs = []struct{ i, j int }{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

func knightProbability(n, k, row, column int) float64 {
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, d := range dirs {
						if x, y := i+d.i, j+d.j; 0 <= x && x < n && 0 <= y && y < n {
							dp[step][i][j] += dp[step-1][x][y] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}

// ÿ��һ��
func findCenter(edges [][]int) int {
	m := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		if m[edges[i][0]] > 0 {
			return edges[i][0]
		}
		if m[edges[i][1]] > 0 {
			return edges[i][1]
		}
		m[edges[i][0]]++
		m[edges[i][1]]++
	}
	return 0
}

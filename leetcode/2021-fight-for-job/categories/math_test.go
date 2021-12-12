package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[数学]
// 第五题
// leetcode 面试题17.10: 主要元素
func majorityElement(nums []int) int {
	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	if 2*count > len(nums) {
		return candidate
	}
	return -1
}

// tag-[数学]
// 第七题
// leetcode566: 重塑矩阵
func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	if r*c > n*m {
		return mat
	}
	out := make([][]int, r)
	for i := 0; i < len(out); i++ {
		out[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			x := (i*c + j) / m
			y := (i*c + j) % m
			out[i][j] = mat[x][y]
		}
	}
	return out
}

func Test_matrixReshape(t *testing.T) {
	mat := [][]int{{1, 2}}
	fmt.Println(matrixReshape(mat, 1, 1))
}

// tag-[数学]
// 第三题
// leetcode 剑指offer 17: 打印从1到最大的n位数
func printNumbers(n int) []int {
	var pow func(n int) int
	pow = func(n int) int {
		if n == 0 {
			return 1
		}
		return 10 * pow(n-1)
	}
	var ret []int
	for i := 0; i < pow(n); i++ {
		ret = append(ret, i)
	}
	return ret
}

// tag-[数学]
// 第一题
// leetcode191: 2的幂
func isPowerOfTwo(n int) bool {
	cnt := 0
	for n != 0 {
		n &= n - 1
		cnt++
	}
	return cnt == 1
}

// tag-[数学]
// 第二题
// leetcode461: 汉明距离
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= num - 1
		cnt++
	}
	return cnt
}

// tag-[数学]
// 第二题
// leetcode 171: Excel表列序号
func titleToNumber(columnTitle string) int {
	l := len(columnTitle)
	sum := 0
	for i := l - 1; i >= 0; i-- {
		sum += int(columnTitle[l-i-1]-'A'+1) * pow(26, i)
	}
	return sum
}

func pow(a, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	return a * pow(a, n-1)
}

func Test_pow(t *testing.T) {
	fmt.Println(pow(10, 0))
}

func Test_titleToNumber(t *testing.T) {
	fmt.Println(titleToNumber("FXSHRXW"))
}

// tag-[数学]
// 第八题
// leetcode 剑指offer 16: 数值的整数次方
func myPow(x float64, n int) float64 {
	if n < 0 {
		return float64(1) / powF(x, -n)
	}
	return powF(x, n)
}

func powF(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		tmp := powF(x, n/2)
		return tmp * tmp
	}
	return powF(x, n-1) * x
}

// tag-[数学]
// 第五题
// leetcode172：阶乘后的零
func trailingZeroes(n int) int {
	cnt := 0
	for n != 0 {
		n /= 5
		cnt += n
	}
	return cnt
}
// tag-[数学]
// leetcode1952: 三除数
func isThree(n int) bool {
	if n == 1 || n == 2 || n == 3 {
		return false
	}
	cnt := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i < n {
				cnt += 2
			} else {
				cnt += 1
			}
		}
	}
	return cnt == 3
}

func Test_isThree(t *testing.T) {
	fmt.Println(isThree(8))
}
// tag-[数学]
// leetcode lcp29: 乐团站位
func orchestraLayout(num int, xPos int, yPos int) int {
	cycle := min(min(num-1-xPos, xPos), min(num-1-yPos, yPos))
	sum := (num - cycle) * cycle * 4
	cycleStart := (sum + 1) % 9
	ans := 0
	if xPos <= yPos {
		ans = (cycleStart + xPos + yPos - cycle<<1) % 9
	} else {
		ans = (cycleStart + (num-2*cycle)*4 - 4 - (xPos + yPos - cycle<<1)) % 9
	}
	if ans == 0 {
		return 9
	}
	return ans
}

func Test_orchestraLayout(t *testing.T) {
	fmt.Println(orchestraLayout(10, 5, 6))
}

// tag-[数学]
// 第三题
// leetcode2029: 石子游戏IX
// 数学
func stoneGameIX(stones []int) bool {
	c := [3]int{}
	for _, v := range stones {
		c[v%3]++
	}
	return checkW(c) || checkW([3]int{c[0], c[2], c[1]})
}

func checkW(c [3]int) bool {
	if c[1] == 0 { // 如果余数为1的个数等于0，则直接看余数为2的个数
		return false
	}
	c[1]--
	turn := 1 + min(c[1], c[2])*2 + c[0]
	if c[1] > c[2] { // 如果以1开头，序列末尾可以再加个1，和也不能被3整除
		turn++
		c[1]--
	}
	return turn%2 == 1 && c[1] != c[2] // 回合为奇数，且还有石子剩余，轮到bob出，则alice胜出
}

// tag-[数学]
// leetcode29
// 快速乘
// x 和 y 是负数，z 是正数
// 判断 z * y >= x 是否成立
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 { // 不能使用除法
		if z&1 > 0 {
			// 需要保证 result + add >= x
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			// 需要保证 add + add >= x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

// tag-[数学]
func divide(dividend, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 // 注意溢出，并且不能使用除法
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 { // 注意溢出
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans
}

// leetcode319 脑筋急转弯
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

// tag-[数学]
// leetcode66:每日一题
func plusOne66(digits []int) []int {
	n := len(digits)
	carry := 1
	for i := n - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
		if i == 0 && carry > 0 {
			return append([]int{carry}, digits...)
		}
	}
	return digits
}

// tag-[数学]
// leetcode229: 求众数
func majorityElement229(nums []int) []int {
	n := len(nums)
	e1, e2 := 0, 0
	vote1, vote2 := 0, 0
	for i := range nums {
		if vote1 > 0 && nums[i] == e1 {
			vote1++
		} else if vote2 > 0 && nums[i] == e2 {
			vote2++
		} else if vote1 == 0 {
			e1 = nums[i]
			vote1++
		} else if vote2 == 0 {
			e2 = nums[i]
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	// 验证
	cnt1, cnt2 := 0, 0
	for i := range nums {
		if vote1 > 0 && nums[i] == e1 {
			cnt1++
		} else if vote2 > 0 && nums[i] == e2 {
			cnt2++
		}
	}
	var ans []int
	if vote1 > 0 && cnt1*3 > n {
		ans = append(ans, e1)
	}
	if vote2 > 0 && cnt2*3 > n {
		ans = append(ans, e2)
	}
	return ans
}
// tag-[数学]
// leetcode166
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	num, den := int64(numerator), int64(denominator)
	isNegtive := false
	if num < 0 {
		isNegtive = !isNegtive
		num = -num
	}
	if den < 0 {
		isNegtive = !isNegtive
		den = -den
	}
	ans := make([]byte, 0)
	m := make(map[int64][2]int64)
	first := true
	cycle := [2]int{}
	for num != 0 {
		quotient := num / den
		reminder := num % den
		if v, ok := m[reminder]; ok && reminder != 0 && v[1] == quotient {
			cycle[0] = int(v[0])
			cycle[1] = len(ans)
			break
		} else if !first {
			m[reminder] = [2]int64{int64(len(ans)), quotient}
		}
		if quotient < 1 {
			if first {
				first = !first
				if len(ans) == 0 {
					ans = append(ans, '0')
				}
				ans = append(ans, '.')
			} else {
				ans = append(ans, '0')
			}
			num *= 10
		} else {
			ans = append(ans, intToBytes(quotient)...)
			if first && reminder != 0 {
				first = !first
				ans = append(ans, '.')
			}
			num = reminder * 10
		}
	}
	if cycle[0] != 0 || cycle[1] != 0 {
		ans = append(ans[:cycle[0]], append([]byte{'('}, ans[cycle[0]:]...)...)
		ans = append(ans, ')')
	}
	if isNegtive {
		ans = append([]byte{'-'}, ans...)
	}
	return string(ans)
}

func Test_fractionToDecimal(t *testing.T) {
	fmt.Println(fractionToDecimal(20, 3))
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(1, 5))
	fmt.Println(fractionToDecimal(45, 698))
	fmt.Println(fractionToDecimal(100, 9))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(0, 3))
	fmt.Println(fractionToDecimal(500, 10))
	fmt.Println(fractionToDecimal(140898435, 17))
	fmt.Println(fractionToDecimal(-50, 8))
}

// tag-[数学]
// leetcode263
func isUgly(n int) bool {
	factors := []int{2, 3, 5}
	if n <= 0 {
		return false
	}
	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}

// tag-[数学]
// leetcode89:格雷编码公式i^(i>>1)
func grayCode(n int) []int {
	size := 1 << n
	ans := make([]int, 0, size)
	for i := 0; i < size; i++ {
		ans = append(ans, i^(i>>1))
	}
	return ans
}

func Test_grayCode(t *testing.T) {
	print_binary_array(grayCode(3), 3)
}

// tag-[数学]
// 蓄水池采样算法，保证留下来的概率都是K/N
func ReservoirSampling(nums []int, k int) []int {
	out := make([]int, k)
	for i := 0; i < k; i++ {
		out[i] = nums[i]
	}
	for i := k; i < len(nums); i++ {
		x := rand.Intn(i)
		if x < k {
			out[x] = nums[i]
		}
	}
	return out
}
// tag-[数学]
// leetcode2063: 所有子字符串中的元音
// 解法二: 直接计算
func countVowels_(word string) int64 {
	ans := int64(0)
	n := len(word)
	for i := 0; i < n; i++ {
		if strings.ContainsRune("aeiou", rune(word[i])) {
			ans += int64((i + 1) * (n - i))
		}
	}
	return ans
}

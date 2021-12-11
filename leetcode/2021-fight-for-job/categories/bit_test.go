package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)
// tag-[位运算]
// 第三题
// leetcode 剑指offer 65: 不用加减乘除做加法
func add(a int, b int) int {
	for b != 0 {
		c := a & b << 1
		a ^= b
		b = c
	}
	return a
}
// tag-[位运算]
// 第五题
// leetcode190: 颠倒二进制位
func reverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		bit <<= 31 - i
		ret += bit
	}
	return ret
}

// leetcode190: 颠倒二进制位
// 方法2
func reverseBits1(n uint32) uint32 {
	n = (n >> 16) | (n << 16)
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)
	return n
}

func Test_reverseBits(t *testing.T) {
	fmt.Println(reverseBits(0b00000010100101000001111010011100))
}
// tag-[位运算]
// 第六题
// leetcode136: 只出现一次的数字
func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}// tag-[位运算]
// 第二题
// leetcode 剑指offer56-I：数组中数字出现的次数
// 数组中两个出现一次的数，分开异或
func singleNumbers(nums []int) []int {
	m, n := 0, 1
	x, y := 0, 0
	for _, v := range nums {
		m ^= v
	}
	for m&n == 0 {
		n <<= 1
	}
	for _, v := range nums {
		if v&n == 0 {
			x ^= v
		} else {
			y ^= v
		}
	}
	return []int{x, y}
}

func Test_singleNumbers(t *testing.T) {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}
// tag-[位运算]
// 第三题
// leetcode剑指offer56-II: 数组中数字出现的次数II
func singleNumberII(nums []int) int {
	m := make([]int, 32)
	for _, v := range nums {
		for i := 0; i < 32; i++ {
			if v&(1<<i) != 0 {
				m[i]++
			}
		}
	}
	ans := 0
	for i, v := range m {
		if v%3 != 0 {
			ans += 1 << i
		}
	}
	return ans
}
// tag-[位运算]
// 第一题
// leetcode1806: 还原排列的最少操作步数
// 位运算
func reinitializePermutation(n int) int {
	ans := 0
	i := 1
	for {
		if i&1 == 0 {
			i = i >> 1
		} else {
			i = n>>1 + (i-1)>>1
		}
		ans++
		if i == 1 {
			break
		}
	}
	return ans
}

func Test_reinitializePermutation(t *testing.T) {
	fmt.Println(reinitializePermutation(8))
}
// tag-[位运算]
// leetcode287: 寻找重复数
// 二进制解法
// 解题思路：如果重复的数字在第i位为1，那么第i位上1的个数大于1~n所有数字第i位上1的个数。
func findDuplicate__(nums []int) int {
	n := len(nums)
	bit_max := 31
	for (n-1)>>bit_max == 0 {
		bit_max--
	}
	ans := 0
	for bit := 0; bit <= bit_max; bit++ {
		x, y := 0, 0
		for i := 0; i < n; i++ {
			if nums[i]&(1<<bit) > 0 {
				x++
			}
			if i >= 1 && (i&(1<<bit)) > 0 {
				y++
			}
		}
		if x > y {
			ans |= 1 << bit
		}
	}
	return ans
}

func Test_findDuplicate1(t *testing.T) {
	fmt.Println(findDuplicate([]int{1, 3, 2, 2, 4}))
	fmt.Println(findDuplicate_([]int{1, 3, 2, 2, 4}))
	fmt.Println(findDuplicate__([]int{1, 3, 2, 2, 4}))
}// tag-[位运算]
// leetcode137:只出现一次的数字
func singleNumber137(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, v := range nums {
			total += (int32(v) >> i) & 1
		}
		if total%3 != 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

func Test_singleNumber137(t *testing.T) {
	// fmt.Println(singleNumber137([]int{2, 2, 3, 2}))
	// fmt.Println(singleNumber137([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber137([]int{2, 2, 2, -1}))
	// fmt.Println(singleNumber137([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}
// tag-[位运算]
// leetcode260
func singleNumber260(nums []int) []int {
	m, n := 0, 1
	for i := range nums {
		m ^= nums[i]
	}
	n = m & (-m)
	x, y := 0, 0
	for i := range nums {
		if nums[i]&n == 0 { // 这样分组的原因是：x和y必然不在一个组，相同的两个数必然在同一个组。就可以进行分组异或。
			x ^= nums[i]
		} else {
			y ^= nums[i]
		}
	}
	return []int{x, y}
}

func Test_singleNumber260(t *testing.T) {
	fmt.Println(singleNumber260([]int{1, 2, 1, 3, 2, 5}))
}

// tag-[二分查找]
// leetcode162:
func findPeakElement(nums []int) int {
	// 满足二段性所以可以用二分查找
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func Test_findPeakElement(t *testing.T) {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}

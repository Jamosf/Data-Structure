package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[哈希表]
// 哈希表
// 第一题
// leetcode12: 整数转罗马数字
func intToRoman(num int) string {
	i := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	r := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var ans string
	for idx := len(i) - 1; idx >= 0; idx-- {
		for num >= i[idx] {
			ans += r[idx]
			num -= i[idx]
		}
	}
	return ans
}

func Test_intToRoman(t *testing.T) {
	fmt.Println(intToRoman(3999))
}

// tag-[哈希表]
// leetcode1995: 统计特殊四元组
func countQuadruplets(nums []int) int {
	n := len(nums)
	m := make(map[int][]int)
	for i := range nums {
		v := m[nums[i]]
		v = append(v, i)
		m[nums[i]] = v
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if vv, ok := m[sum]; ok {
					for _, v := range vv {
						if v > k {
							ans++
						}
					}
				}
			}
		}
	}
	return ans
}

func Test_count(t *testing.T) {
	fmt.Println(countQuadruplets([]int{9, 6, 23, 8, 39, 23}))
}

// tag-[哈希表]
func minimumSwitchingTimes(source [][]int, target [][]int) int {
	m, n := len(source), len(source[0])
	mk := map[int]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mk[source[i][j]]++
			mk[target[i][j]]--
		}
	}
	cnt := 0
	for _, v := range mk {
		if v > 0 {
			cnt += v
		} else {
			cnt += -v
		}
	}
	return cnt >> 1
}

func Test_minimumSwitchingTimes(t *testing.T) {
	fmt.Println(minimumSwitchingTimes([][]int{{1, 3}, {5, 4}}, [][]int{{3, 1}, {6, 5}}))
	fmt.Println(maxmiumScore([]int{1, 2, 8, 9}, 3))
}

// tag-[哈希表]
// leetcode2006：差的绝对值为k的数对数目
func countKDifference(nums []int, k int) int {
	m := make(map[int]int)
	ans := 0
	for i := range nums {
		ans += m[nums[i]-k]
		ans += m[nums[i]+k]
		m[nums[i]]++
	}
	return ans
}

// tag-[哈希表]
// leetcode380
type RandomizedSet struct {
	m map[int]int
	l []int
}

func ConstructorRandomizedSet() RandomizedSet {
	return RandomizedSet{m: make(map[int]int), l: make([]int, 0)}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.m[val]; !ok {
		r.l = append(r.l, val)
		r.m[val] = len(r.l) - 1
		return true
	}
	return false
}

func (r *RandomizedSet) Remove(val int) bool {
	if idx, ok := r.m[val]; ok {
		delete(r.m, val)
		n := len(r.l)
		r.l[idx], r.l[n-1] = r.l[n-1], r.l[idx]
		r.l = r.l[:n-1]
		if len(r.l) > idx {
			r.m[r.l[idx]] = idx
		}
		return true
	}
	return false
}

func (r *RandomizedSet) GetRandom() int {
	return r.l[rand.Intn(len(r.l))]
}

func Test_RandomizedSet(t *testing.T) {
	v := ConstructorRandomizedSet()
	vv := &v
	vv.Remove(0)
	vv.Remove(0)
	vv.Insert(0)
	vv.GetRandom()
	vv.Remove(0)
	vv.Insert(0)
}
// tag-[哈希表]
// leetcode869：预处理加hash表，词频统计;思路，因为可以任意顺序排列，则词频相同的最终可以排列等到的数据是一致的。
// 时间复杂度：O(logn)
// 空间负责度：O(1)
func reorderedPowerOf2_(n int) bool {
	m := make(map[[10]int]bool)
	countDigital := func(v int) [10]int {
		cnt := [10]int{}
		for v != 0 {
			cnt[v%10]++
			v /= 10
		}
		return cnt
	}
	for i := 1; i < 1e9; i <<= 1 {
		m[countDigital(i)] = true
	}
	return m[countDigital(n)]
}

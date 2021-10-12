// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/list"
	"fmt"
	"sort"
	"testing"
)

func smallestK(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	return quickSort1(arr, 0, len(arr)-1, k)
}

func quickSort1(nums []int, l, r int, k int) []int {
	if l+1 >= r {
		return nums
	}
	first, last := l, r-1
	key := nums[first]
	for first < last {
		for first < last && nums[last] >= key {
			last--
		}
		nums[first] = nums[last]
		for first < last && nums[first] <= key {
			first++
		}
		nums[last] = nums[first]
	}
	nums[first] = key
	if first > k {
		return quickSort1(nums, l, first, k)
	}
	if first < k {
		return quickSort1(nums, first+1, r, k)
	}
	return nums[:k]
}

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

func numberOfWeakCharacters(properties [][]int) int {
	n := len(properties)
	sort.Slice(properties, func(i, j int) bool {
		a, b := properties[i], properties[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})
	ans := 0
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) != 0 && properties[stack[len(stack)-1]][1] < properties[i][1] {
			ans++
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func Test_numberOfWeakCharacters(t *testing.T) {
	fmt.Println(numberOfWeakCharacters([][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}}))
}

// 超时了
func firstDayBeenInAllRooms(nextVisit []int) int {
	m := make(map[int]int)
	n := len(nextVisit)
	visited := 0
	current := 0
	ans := 0
	for visited < n {
		if _, ok := m[current]; !ok {
			visited++
		}
		ans++
		m[current]++
		if m[current]&1 == 1 {
			current = nextVisit[current]
		} else {
			current = (current + 1) % n
		}
	}
	return ans
}

// dp求解, +mod 可以防止出现负数
func firstDayBeenInAllRooms1(nextVisit []int) int {
	mod := int64(1e9 + 7)
	n := len(nextVisit)
	dp := make([]int64, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = (dp[i-1]*2 - dp[nextVisit[i-1]] + 2 + mod) % mod
	}
	return int((dp[n-1]) % mod)
}

func Test_firstDayBeenInAllRooms1(t *testing.T) {
	fmt.Println(firstDayBeenInAllRooms1([]int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 23, 24, 24, 25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 45, 45, 46, 46, 47, 47, 48}))
}

func gcdSort(nums []int) bool {
	n := len(nums)
	u := newUnionFind(1e5 + 1)
	tmp := make([]int, n)
	for i := range nums {
		for a := 2; a*a <= nums[i]; a++ {
			if nums[i]%a == 0 {
				b := nums[i] / a
				u.union(nums[i], a)
				u.union(nums[i], b)
			}
		}
		tmp[i] = nums[i]
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] != tmp[i] && !u.isConnected(nums[i], tmp[i]) {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Test_gcdSort(t *testing.T) {
	fmt.Println(gcdSort([]int{8, 9, 4, 2, 3}))
}

// 注意：使用切片作为队列时，长度申请不合理，会超时。使用list不会超时。
func findFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])
	diret := [2][2]int{{1, 0}, {0, 1}}
	var bfs func(i, j int) []int
	bfs = func(i, j int) []int {
		l := list.New()
		l.PushBack([]int{i, j})
		land[i][j] = 0
		maxx, maxy := i, j
		for l.Len() != 0 {
			v := l.Front()
			l.Remove(v)
			vv := v.Value.([]int)
			for p := 0; p < 2; p++ {
				x, y := vv[0]+diret[p][0], vv[1]+diret[p][1]
				if x >= 0 && x < m && y >= 0 && y < n && land[x][y] == 1 {
					l.PushBack([]int{x, y})
					land[x][y] = 0
					maxx = max(maxx, x)
					maxy = max(maxy, y)
				}
			}
		}
		return []int{i, j, maxx, maxy}
	}
	var ans [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] != 0 {
				ans = append(ans, bfs(i, j))
			}
		}
	}
	return ans
}

func Test_findFarmland(t *testing.T) {
	fmt.Println(findFarmland([][]int{{1, 1}, {1, 1}}))
}

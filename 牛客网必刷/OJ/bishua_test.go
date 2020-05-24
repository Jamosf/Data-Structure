package main

import (
	"fmt"
	"math"
	"testing"
)

// 01 暴力法
func minimumValueAfterDispel(nums []int) int64 {
	// write code here
	var sum int64
	var minSum int64 = math.MaxInt64
	for i := 0; i < len(nums); i++ {
		sum = 0
		for _, num := range nums {
			if nums[i] < num {
				sum += int64(num - nums[i])
				continue
			}
			sum += int64(num)
		}
		if sum < minSum {
			minSum = sum
		}
	}
	return minSum
}

func Test_minimumValueAfterDispel(t *testing.T) {
	minimumValueAfterDispel([]int{})
}

// 02 暴力法
func countLR(a []int, b []int) int {
	// write code here
	n := len(a)
	var count = 0
	for i := 0; i < n; i++ {
		suma := a[i]
		sumb := b[i]
		for j := i; j < n; j++ {
			if i != j {
				suma += a[j]
			}
			if suma == sumb+b[j] {
				count++
			}
		}
	}
	return count
}

func Test_countLR(t *testing.T) {
	fmt.Println(countLR([]int{99, 25, 93, 28, 29, 78, 91, 85, 67, 67, 54, 41, 66, 39, 4, 33, 57, 68, 45, 56, 77, 0, 94, 13, 67, 72, 62, 15, 17, 36, 12, 30, 97, 52, 76, 17, 21, 99, 62, 46, 97, 55, 26, 8, 4, 6, 46, 86, 47, 47, 51, 21, 60, 18, 79, 85, 97, 28, 76, 60, 87, 48, 36, 74, 51, 69, 28, 8, 48, 78, 62, 18, 83, 4, 6, 73, 10, 99},
		[]int{1509, 118, 1240, 571, 1190, 1631, 1539, 241, 152, 656, 813, 246, 940, 389, 211, 1937, 330, 1822, 1830, 1051, 1070, 1584, 540, 1412, 1798, 1649, 1374, 492, 1590, 1072, 8, 234, 1765, 1419, 348, 1881, 317, 1927, 431, 1446, 991, 1595, 42, 1203, 1132, 1301, 1130, 1655, 1421, 84, 1760, 444, 1148, 220, 1186, 671, 1977, 1861, 920, 1380, 806, 1676, 1072, 527, 491, 1073, 75, 1052, 1667, 263, 750, 648, 1661, 1591, 386, 949, 1276, 1038}))
}

/**
 * 计算六边形01矩阵中最大的全1子三角形的边长
 * @param a int整型 六边形01矩阵的边长
 * @param maps int整型一维数组 矩阵的具体数据，从上到下，从左到右顺次排列
 * @return int整型
 */
// 03 动态规划
func largestSubTriangle(a int, maps []int) int {
	// write code here
}

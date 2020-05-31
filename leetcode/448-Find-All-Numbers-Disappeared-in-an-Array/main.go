// 448-Find-All-Numbers-Disappeared-in-an-Array project main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(findDisappearedNumbers([]int{1, 1, 2, 2}))
}

//func findDisappearedNumbers(nums []int) []int {
//	l := len(nums)
//	nums = append(nums, 1)
//	nums = append(nums, l)
//	sort.Ints(nums)
//	var r []int
//	var d int = 1
//	if nums[0] < nums[1] {
//		r = append(r, nums[0])
//	}
//	for i := 1; i < l+1; i++ {
//		for nums[i]+d < nums[i+1] {
//			r = append(r, nums[i]+d)
//			d++
//		}
//		d = 1
//	}
//	if nums[l] < nums[l+1] {
//		r = append(r, nums[l+1])
//	}
//	return r
//}

func findDisappearedNumbers(nums []int) []int {
	l := len(nums)
	m := make(map[int]bool)
	var r []int
	for i := 0; i < l; i++ {
		m[nums[i]] = true
	}
	for j := 1; j <= l; j++ {
		if _, ok := m[j]; !ok {
			r = append(r, j)
		}
	}
	return r
}

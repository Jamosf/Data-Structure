// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package basic_algo

// 快速排序算法
func QuickSort(nums []int, l, r int) {
	if l+1 >= r {
		return
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
	QuickSort(nums, l, first)
	QuickSort(nums, first+1, r)
}

func QuickSortK(nums []int, l, r int, k int) []int {
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
		return QuickSortK(nums, l, first, k)
	}
	if first < k {
		return QuickSortK(nums, first+1, r, k)
	}
	return nums[:k]
}

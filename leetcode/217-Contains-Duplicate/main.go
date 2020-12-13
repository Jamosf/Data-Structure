package _17_Contains_Duplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums{
		if _, ok := m[v]; ok{
			return true
		}
		m[v] = true
	}
	return false
}

func containsDuplicate_(nums []int) bool {
	m := make(map[int]bool, len(nums)/2)
	for i := range nums{
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
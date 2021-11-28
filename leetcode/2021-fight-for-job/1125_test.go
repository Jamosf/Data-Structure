package ojeveryday

import "sort"

// leetcode692: 前K个高频单词
func topKFrequent_(words []string, k int) []string {
	m := make(map[string]int)
	maxn := 0
	for i := range words {
		m[words[i]]++
		maxn = max(maxn, m[words[i]])
	}
	bucket := make([][]string, maxn+1)
	for k, v := range m {
		bucket[v] = append(bucket[v], k)
	}
	ans := make([]string, 0, k)
	for i := maxn; i >= 0; i-- {
		if k == 0 {
			break
		}
		if bucket[i] != nil {
			sort.Strings(bucket[i])
			for j := 0; j < len(bucket[i]) && k > 0; j++ {
				ans = append(ans, bucket[i][j])
				k--
			}
		}
	}
	return ans
}

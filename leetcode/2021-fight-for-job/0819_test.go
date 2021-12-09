// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
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

// tag-[回溯]
// 第二题
// leetcode17: 电话号码的字母组合
func letterCombinations(digits string) []string {
	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	size := len(digits)
	if size == 0 {
		return nil
	}
	var ret []string
	var tmp []byte
	var dfs func(int)
	dfs = func(level int) {
		if level == size {
			ret = append(ret, string(tmp))
			return
		}
		// 没有标记是否访问过，所有的解都是可行的
		for j := 0; j < len(m[digits[level]-'0']); j++ {
			v := m[digits[level]-'0'][j] - 'a'
			tmp = append(tmp, v+'a')
			dfs(level + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0)
	return ret
}

func Test_letterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("22"))
}

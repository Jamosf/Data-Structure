// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

func magicalString(n int) int {
	str := "122"
	fast := 2
	ans := 1
	for i := 2; i < n; i++ {
		if str[i] == '2' && fast < n-2 {
			if str[fast] == '2' {
				str += "11"
				ans += 2
			} else {
				str += "22"
			}
			fast += 2
		}
		if str[i] == '1' && fast < n-1 {
			if str[fast] == '2' {
				str += "1"
				ans++
			} else {
				str += "2"
			}
			fast++
		}
	}
	return ans
}

func Test_magicalString(t *testing.T) {
	fmt.Println(magicalString(4))
}

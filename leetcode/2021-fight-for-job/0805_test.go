// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// 第一题
func convertToTitle(columnNumber int) string {
	var ans []uint8
	for columnNumber > 0 {
		tmp := columnNumber % 26
		if tmp == 0 {
			ans = append(ans, 'Z')
			columnNumber = columnNumber/26 - 1
		} else {
			ans = append(ans, uint8(tmp+'A'-1))
			columnNumber = columnNumber / 26
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return string(ans)
}

func Test_convert(t *testing.T) {
	fmt.Println(convertToTitle(701))
}

// 第二题
func maxProfitII(prices []int) int {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}

func Test_maxProfitII(t *testing.T) {
	fmt.Println(maxProfitII([]int{1, 2, 3, 4, 5}))
}

// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

func Test_abc(t *testing.T) {
	a := 'a'
	fmt.Println(a)
	fmt.Println(a ^ 32)
	fmt.Println(a | 32)
	fmt.Println(a & -33)
	b := 'B'
	fmt.Println(b)
	fmt.Println(b ^ 32)
	fmt.Println(b | 32)
	fmt.Println(b & -33)

	c := 10
	fmt.Println(^c + 1)
	fmt.Println(-c)
	fmt.Println(c & -c)
}

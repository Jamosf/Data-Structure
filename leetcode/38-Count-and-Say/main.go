// 38-Count-and-Say project main.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	r := []rune(countAndSay(n - 1))
	var result []string
	var l, val int
	result = append(result, "1", string(r[0]))
	for i := 1; i < len(r); i++ {
		l = len(result)
		if string(r[i]) == result[l-1] {
			val, _ = strconv.Atoi(result[l-2])
			result[l-2] = strconv.Itoa(val + 1)
			continue
		}
		result = append(result, "1", string(r[i]))
	}
	return strings.Join(result, "")
}

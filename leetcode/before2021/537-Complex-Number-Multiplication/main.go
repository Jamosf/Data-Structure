// 537-Complex-Number-Multiplication project main.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}

func complexNumberMultiply(a string, b string) string {
	listA := strings.Split(a, "+")
	listB := strings.Split(b, "+")
	ra, _ := strconv.Atoi(listA[0])
	rb, _ := strconv.Atoi(listB[0])
	ima, _ := strconv.Atoi(strings.Split(listA[1], "i")[0])
	imb, _ := strconv.Atoi(strings.Split(listB[1], "i")[0])
	r := ra*rb - int(ima)*int(imb)
	fmt.Println(listA, listB, ra, rb, ima, imb, r)
	im := ra*int(imb) + rb*int(ima)
	fmt.Println(im)
	var s []string
	s = append(s, strconv.Itoa(r))
	s = append(s, strconv.Itoa(im)+"i")
	return strings.Join(s, "+")
}

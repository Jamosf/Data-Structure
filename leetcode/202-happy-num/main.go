package main

import "fmt"

func main() {
	ret := isHappy(19)
	fmt.Println(ret)
}

func isHappyExec(n int, m map[int]bool) bool {
	var sum int = 0
	for n%10 != 0 || n/10 != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	if sum == 1 {
		return true
	}
	if _, ok := m[sum]; ok {
		return false
	}
	if sum == 0 {
		return false
	}
	m[sum] = true
	fmt.Println("the sum is ", sum)
	//return false
	return isHappyExec(sum, m)
}

func isHappy(n int) bool {
	m := make(map[int]bool, 1)
	return isHappyExec(n, m)
}

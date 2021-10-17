package _021_fight_for_job

import (
	"fmt"
	"strconv"
	"testing"
)

func areNumbersAscending(s string) bool {
	last := 0
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] >= '0' && s[j] <= '9' {
			j++
		}
		if j > i {
			v, _ := strconv.Atoi(s[i:j])
			fmt.Println(v, last)
			if v <= last {
				return false
			}
			last = v
		}
		i = j + 1
	}
	return true
}

func Test_areNumbersAscending(t *testing.T) {
	fmt.Println(areNumbersAscending("sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"))
	fmt.Println(areNumbersAscending("hello world 5 x 5"))
	fmt.Println(areNumbersAscending("4 5 11 26"))
}

type Bank struct {
	money []int64
	n     int
}

func Constructor_n(balance []int64) Bank {
	return Bank{money: balance, n: len(balance)}
}

func (b *Bank) isValid(account int) bool {
	return account >= 1 && account <= b.n
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if b.isValid(account1) && b.isValid(account2) {
		if b.money[account1-1] >= money {
			b.money[account1-1] -= money
			b.money[account2-1] += money
			return true
		}
	}
	return false
}

func (b *Bank) Deposit(account int, money int64) bool {
	if b.isValid(account) {
		b.money[account-1] += money
		return true
	}
	return false
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if b.isValid(account) && b.money[account-1] >= money {
		b.money[account-1] -= money
		return true
	}
	return false
}

func or(v []int) int {
	ans := 0
	for i := range v {
		ans |= v[i]
	}
	return ans
}

func countMaxOrSubsets(nums []int) int {
	var backtrace func(index int)
	m := make(map[int]int)
	n := len(nums)
	sum := 0
	v := make([]int, 0)
	maxn := 0
	backtrace = func(index int) {
		sum = or(v)
		if sum > 0 {
			m[sum]++
			if sum >= maxn {
				maxn = sum
			}
		}
		if index == n {
			return
		}
		for i := index; i < n; i++ {
			v = append(v, nums[i])
			backtrace(i + 1)
			v = v[:len(v)-1]
		}
	}
	backtrace(0)
	return m[maxn]
}

func Test_countMaxOrSubsets(t *testing.T) {
	fmt.Println(countMaxOrSubsets([]int{3, 1}))
	fmt.Println(countMaxOrSubsets([]int{2, 2, 2}))
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5}))
}

func Test_bin(t *testing.T) {
	sum := 10
	sum |= 51
	fmt.Println(sum)
	sum ^= 51
	fmt.Println(sum)
}

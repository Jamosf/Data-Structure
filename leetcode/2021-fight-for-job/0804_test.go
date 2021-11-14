package ojeveryday

import (
	"fmt"
	"strings"
	"testing"
)

// 第一题
// leetcode27: 移除元素
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			if i < len(nums)-1 {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				nums = nums[:i]
			}
			continue
		}
		i++
	}
	return len(nums)
}

// 第二题
// leetcode125：验证回文串
func isPalindrome(s string) bool {
	puneStr := make([]rune, 0, len(s))
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			puneStr = append(puneStr, v)

		}
	}
	str := strings.ToLower(string(puneStr))
	left, right := 0, len(puneStr)-1
	for left <= right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// leetcode125：验证回文串，解法二
func isPalindrome_(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left <= right {
		if !isDigitalOrChar(s[left]) {
			left++
			continue
		}
		if !isDigitalOrChar(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isDigitalOrChar(v uint8) bool {
	return (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9')
}

func Test_isPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome_("A man, a plan, a canal: Panama"))
}

// 第三题
// leetcode66: 加一
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + carry
		digits[i] = tmp % 10
		carry = tmp / 10
	}
	if carry != 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}

func Test_plusOne(t *testing.T) {
	fmt.Println(plusOne([]int{9}))
}

// 第四题
// leetcode58：最后一个单词长度
func lengthOfLastWord(s string) int {
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if cnt != 0 {
				return cnt
			}
		} else {
			cnt++
		}
	}
	return cnt
}

// 第五题
// leetcode172：阶乘后的零
func trailingZeroes(n int) int {
	cnt := 0
	for n != 0 {
		n /= 5
		cnt += n
	}
	return cnt
}

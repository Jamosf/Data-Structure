package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[字符串]
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
// tag-[字符串]
// leetcode718: 最长重复子数组
func findLength(nums1 []int, nums2 []int) int {
	lenA, lenB := len(nums1), len(nums2)
	ret := 0
	for i := 0; i < lenA; i++ {
		k := 0
		for j := 0; j < min(lenA-i, lenB); j++ {
			if nums1[i+j] == nums2[j] {
				k++
			} else {
				k = 0
			}
			ret = max(ret, k)
		}
	}

	for i := 0; i < lenB; i++ {
		k := 0
		for j := 0; j < min(lenB-i, lenA); j++ {
			if nums1[j] == nums2[j+i] {
				k++
			} else {
				k = 0
			}
			ret = max(ret, k)
		}
	}
	return ret
}

// leetcode187: 重复DNA序列
func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]uint8, len(s)-10)
	for i := 0; i < len(s)-10+1; i++ {
		m[s[i:i+10]]++
	}
	var result []string
	for k, v := range m {
		if v > 1 {
			result = append(result, k)
		}
	}
	return result
}

// leetcode1461: 检查一个字符串是否包含所有长度为K的二进制子串
func hasAllCodes(s string, k int) bool {
	m := make(map[string]uint8, len(s)-k)
	for i := 0; i < len(s)-k+1; i++ {
		m[s[i:i+k]]++
	}
	result := 0
	for _, v := range m {
		if v >= 1 {
			result++
		}
	}
	return result == 1<<k
}
// tag-[字符串]
// leetcode 面试题01.01：字符串压缩
func compressString(S string) string {
	var buff bytes.Buffer
	for i := 0; i < len(S); {
		cnt := 1
		tmp := S[i]
		for i < len(S) {
			if i < len(S)-1 && S[i] != S[i+1] {
				i++
				break
			}
			if i < len(S)-1 {
				cnt++
			}
			i++
		}
		buff.WriteString(string(tmp))
		buff.WriteString(strconv.Itoa(cnt))
	}
	r := buff.String()
	if len(r) > len(S) {
		return S
	}

	return r
}

func Test_compressString(t *testing.T) {
	fmt.Println(compressString("rrrrrLLLLLPPPPPPRRRRRgggNNNNNVVVVVVVVVVDDDDDDDDDDIIIIIIIIIIlllllllAAAAqqqqqqqbbbNNNNffffff"))
}
// tag-[字符串]
// 第一题
// leetcode3: 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	left, right := 0, 0
	maxn := 0
	for left = 0; left < len(s); left++ {
		if left != 0 {
			delete(m, s[left-1])
		}
		for right < len(s) && m[s[right]] == 0 {
			m[s[right]]++
			right++
		}
		maxn = max(maxn, right-left)
	}
	return maxn
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring(" "))
}

// tag-[字符串]
// 第二题
// leetcode567: 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	m := make(map[uint8]int)
	for i := range s1 {
		m[s1[i]]++
	}
	left, right := 0, len(s1)
	for i := left; i <= right; i++ {
		if m[s2[i]] != 0 {
			m[s2[i]]--
		}
	}
	if isMapEmpty(m) {
		return true
	}
	for right < len(s2) {
		left++
		right++
		if m[s2[left-1]] >= 0 {

		}
	}
	return false
}

func isMapEmpty(m map[uint8]int) bool {
	cnt := 0
	for _, v := range m {
		cnt += v
	}
	return cnt == 0
}

// tag-[字符串]
// 第三题
// leetcode387: 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	var m [26]int
	for _, v := range s {
		m[v-'a']++
	}
	for i, v := range s {
		if m[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

// tag-[字符串]
// 第四题
// leetcode383: 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	var m [26]int
	for _, v := range magazine {
		m[v-'a']++
	}
	for _, v := range ransomNote {
		m[v-'a']--
	}
	for _, v := range m {
		if v < 0 {
			return false
		}
	}
	return true
}

// tag-[字符串]
// 第五题
// leetcode242: 有效的字母异位词
func isAnagram(s string, t string) bool {
	var m [26]int
	for _, v := range s {
		m[v-'a']++
	}
	var n [26]int
	for _, v := range t {
		n[v-'a']++
	}
	return m == n
}
// tag-[字符串]
// 第十六题
// leetcode 剑指offer 58-I: 翻转单词顺序
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	stack := make([]string, len(ss))
	for i := range ss {
		stack[len(stack)-i-1] = ss[i]
	}
	var out string
	for i := 0; i < len(stack); i++ {
		if stack[i] != "" {
			if len(out) == 0 {
				out = stack[i]
			} else {
				out += " " + stack[i]
			}
		}
	}
	return out
}

func Test_reverseWords(t *testing.T) {
	fmt.Println(reverseWords("  hello world!  "))
}
// tag-[字符串]
// 第五题
// leetcode 剑指offer58-II: 左旋转字符串
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// tag-[字符串]
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

// tag-[字符串]
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

// tag-[字符串]
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

// tag-[字符串]
// 第一题
// leetcode168: Excel表列名称
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

// tag-[字符串]
// leetcode1980: 找出不同的二进制字符串
func findDifferentBinaryString(nums []string) string {
	m := make(map[string]bool)
	for i := range nums {
		m[nums[i]] = true
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums[i]) - 1; j >= 0; j-- {
			if nums[i][j] == '0' {
				b := []byte(nums[i])
				b[j] = '1'
				if _, ok := m[string(b)]; !ok {
					return string(b)
				}
			}
			if nums[i][j] == '1' {
				b := []byte(nums[i])
				b[j] = '0'
				if _, ok := m[string(b)]; !ok {
					return string(b)
				}
			}
		}
	}
	return ""
}

func Test_findDifferentBinaryString(t *testing.T) {
	fmt.Println(findDifferentBinaryString([]string{"1"}))
}

// tag-[字符串]
// 第七题
// leetcode205：同构字符串
func isIsomorphic(s string, t string) bool {
	return isIsomorphicExec(s, t) && isIsomorphicExec(t, s)
}

func isIsomorphicExec(s string, t string) bool {
	m := make(map[byte]byte)
	for i := range s {
		if v, ok := m[s[i]]; !ok {
			m[s[i]] = t[i]
		} else {
			if v != t[i] {
				return false
			}
		}
	}
	return true
}

func Test_isIsomorphic(t *testing.T) {
	fmt.Println(isIsomorphic("egt", "add"))
}

// tag-[字符串]
// leetcode165: 比较版本号
// 分割字符串
func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	t1 := make([]int, len(s1))
	s2 := strings.Split(version2, ".")
	t2 := make([]int, len(s2))
	for i := range s1 {
		v := strings.TrimLeft(s1[i], "0")
		if v != "" {
			t1[i], _ = strconv.Atoi(v)
		} else {
			t1[i] = 0
		}
	}
	for i := range s2 {
		v := strings.TrimLeft(s2[i], "0")
		if v != "" {
			t2[i], _ = strconv.Atoi(v)
		} else {
			t2[i] = 0
		}
	}
	n := min(len(t1), len(t2))
	for i := 0; i < n; i++ {
		if t1[i] > t2[i] {
			return 1
		} else if t1[i] < t2[i] {
			return -1
		}
	}
	if len(t1) > len(t2) {
		for i := len(t2); i < len(t1); i++ {
			if t1[i] > 0 {
				return 1
			}
		}
	}
	if len(t1) < len(t2) {
		for i := len(t1); i < len(t2); i++ {
			if t2[i] > 0 {
				return -1
			}
		}
	}
	return 0
}

// leetcode165: 比较版本号
// 双指针解法
func compareVersion_(v1 string, v2 string) int {
	n1, n2 := len(v1), len(v2)
	pre1, pre2 := 0, 0
	pos1, pos2 := 0, 0
	for pos1 < n1 && pos2 < n2 {
		for pos1 < n1 && v1[pos1] != '.' {
			pos1++
		}
		for pos2 < n2 && v2[pos2] != '.' {
			pos2++
		}
		for pre1 < pos1 && v1[pre1] == '0' {
			pre1++
		}
		for pre2 < pos2 && v2[pre2] == '0' {
			pre2++
		}
		if idx := compare(v1[pre1:pos1], v2[pre2:pos2]); idx != 0 {
			return idx
		}
		pos1++
		pos2++
		pre1, pre2 = pos1, pos2
	}
	for pos1 < n1 {
		if v1[pos1] > '0' && v1[pos1] != '.' {
			return 1
		}
		pos1++
	}
	for pos2 < n2 {
		if v2[pos2] > '0' && v2[pos2] != '.' {
			return -1
		}
		pos2++
	}
	return 0
}

func compare(s, t string) int {
	if len(s) > len(t) {
		return 1
	}
	if len(s) < len(t) {
		return -1
	}
	if s > t {
		return 1
	}
	if s < t {
		return -1
	}
	return 0
}

// leetcode165: 比较版本号
// leetcode 100%的典范代码
func compareVersion__(version1 string, version2 string) int {
	v1 := NewVersionIterator(version1)
	v2 := NewVersionIterator(version2)
	for {
		r1, ok1 := v1.NextRevision()
		r2, ok2 := v2.NextRevision()
		if r1 > r2 {
			return 1
		}

		if r1 < r2 {
			return -1
		}

		if !ok1 && !ok2 {
			break
		}
	}

	return 0
}

type VersionIterator struct {
	version string
	index   int
}

func NewVersionIterator(version string) *VersionIterator {
	return &VersionIterator{
		version: version,
	}
}

func (v *VersionIterator) NextRevision() (int, bool) {
	if v.index == len(v.version) {
		return 0, false
	}

	revision := 0
	for {
		ch := v.version[v.index]
		v.index++
		if ch == '.' {
			break
		}

		revision = revision*10 + int(ch-'0')
		if v.index == len(v.version) {
			break
		}
	}

	return revision, true
}

func Test_compareVersion(t *testing.T) {
	fmt.Println(compareVersion("1.2", "1.10"))
	fmt.Println(compareVersion_("1.2", "1.10"))
	fmt.Println(compareVersion__("1.2", "1.10"))
}

// tag-[字符串]
// leetcode214: 最短回文串
// 输入：s = "aacecaaa"
// 输出："aaacecaaa"
// 解法：将字符串翻转，然后去掉中间重叠的部分即可
func shortestPalindrome(s string) string {
	n := len(s)
	s1 := []byte(s)
	for i := 0; i < n/2; i++ {
		s1[i], s1[n-i-1] = s1[n-i-1], s1[i]
	}
	ss1 := string(s1)
	// 技巧：逆序遍历，可以最大化的删除重复元素，保证最终得到的字符串最小
	for i := n; i >= 0; i-- {
		if s[:i] == ss1[n-i:] {
			return ss1[:n-i] + s
		}
	}
	return ""
}

func Test_shortestPalindrome(t *testing.T) {
	fmt.Println(shortestPalindrome("aacecaaa"))
}
// tag-[字符串]
// leetcode1974: 使用特殊打字机键入单词的最少次数
func minTimeToType(word string) int {
	n := len(word)
	var pre int = 'a'
	ans := 0
	for i := 0; i < n; i++ {
		ans += minAbs(int(word[i])-pre, pre+26-int(word[i]))
		ans += 1
		pre = int(word[i])
	}
	return ans
}

func Test_minTimeToType(t *testing.T) {
	fmt.Println(minTimeToType("bza"))
}

// tag-[字符串]
// 优秀代码学习，枚举回文串
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

// tag-[字符串]
// leetcode1957: 删除字符串使字符串变好
func makeFancyString(s string) string {
	n := len(s)
	ans := make([]byte, 0)
	ans = append(ans, s[0])
	for i := 1; i < n; {
		if s[i] != s[i-1] {
			ans = append(ans, s[i])
			i++
			continue
		}
		cnt := 0
		for i < n && s[i] == s[i-1] {
			if cnt < 1 {
				ans = append(ans, s[i])
			}
			cnt++
			i++
		}
	}
	return string(ans)
}

func Test_makeFancyString(t *testing.T) {
	fmt.Println(makeFancyString("leeettttccccoooooddddeeee"))
}

// tag-[矩阵]
// leetcode1958: 检查操作是否合法
func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	m, n := len(board), len(board[0])
	var check func(dx, dy int) bool
	check = func(dx, dy int) bool {
		x, y := rMove+dx, cMove+dy
		step := 1
		for x >= 0 && x < m && y >= 0 && y < n {
			if step == 1 {
				if board[x][y] == color || board[x][y] == '.' {
					return false
				}
			} else {
				if board[x][y] == '.' {
					return false
				}
				if board[x][y] == color {
					return true
				}
			}
			step++
			x += dx
			y += dy
		}
		return false
	}
	direct := [8][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	for i := 0; i < 8; i++ {
		if check(direct[i][0], direct[i][1]) {
			return true
		}
	}
	return false
}

// tag-[字符串]
// leetcode1945: 字符串转化后的各位数字之和
func getLucky(s string, k int) int {
	ans := make([]byte, 0)
	for i := range s {
		v := s[i] - 'a' + 1
		if v > 9 {
			ans = append(ans, v/10)
		}
		ans = append(ans, v%10)
	}
	tmp := countSum(string(ans), k)
	ret, _ := strconv.Atoi(tmp)
	return ret
}

func countSum(s string, k int) string {
	if k == 0 {
		return s
	}
	tmp := 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v >= '0' {
			v -= '0'
		}
		tmp += int(v)
	}
	return countSum(strconv.Itoa(tmp), k-1)
}

func Test_getLucky(t *testing.T) {
	fmt.Println(getLucky("dbvmfhnttvr", 5))
}

// tag-[字符串]
// leetcode1946: 子字符串突变后可能得到的最大整数
func maximumNumber(num string, change []int) string {
	b := []byte(num)
	cnt := 0
	for i := 0; i < len(b); i++ {
		v := change[b[i]-'0']
		if int(b[i]-'0') < v {
			b[i] = byte(v + '0')
			cnt++
		} else if int(b[i]-'0') > v {
			if cnt != 0 {
				break
			}
		}
	}
	return string(b)
}

func Test_maximumNumber(t *testing.T) {
	fmt.Println(maximumNumber("334111", []int{0, 9, 2, 3, 3, 2, 5, 5, 5, 5}))
}

// tag-[字符串]
// leetcode2000: 反转单次前缀
func reversePrefix(word string, ch byte) string {
	idx := strings.Index(word, string(ch))
	if idx == -1 {
		return word
	}
	b := []byte(word[:idx+1])
	for i := 0; i <= idx/2; i++ {
		b[i], b[idx-i] = b[idx-i], b[i]
	}
	return string(b) + word[idx+1:]
}

// tag-[字符串]
// leetcode2002: 两个回文子序列长度的最大乘积
func maxProduct(s string) int {
	n := len(s)
	m := map[int]int{}
	for i := 1; i < 1<<n-1; i++ {
		t := make([]byte, 0)
		for idx := 0; idx < n; idx++ {
			if 1<<idx&i == 1<<idx {
				t = append(t, s[n-idx-1])
			}
		}
		if isPlalindrome(t) {
			m[i] = len(t)
		}
	}
	maxn := 0
	for i := 1; i < 1<<n-1; i++ {
		for j := i - 1; j >= 0; j-- {
			if i&j == 0 && m[i] != 0 && m[j] != 0 {
				fmt.Println(i, j, m[i], m[j])
				maxn = max(maxn, m[i]*m[j])
			}
		}
	}
	return maxn
}

func isPlalindrome(b []byte) bool {
	j := len(b) - 1
	for i := 0; i < j; i++ {
		if b[i] != b[j] {
			return false
		}
		j--
	}
	return true
}

func Test_maxProduct(t *testing.T) {
	fmt.Println(maxProduct("leetcodecom"))
}

// tag-[字符串]
// leetcode2011：执行操作后的变量值
func finalValueAfterOperations(operations []string) int {
	ans := 0
	for i := range operations {
		if operations[i] == "X++" || operations[i] == "++X" {
			ans++
		} else {
			ans--
		}
	}
	return ans
}

// tag-[字符串]
func longestSubsequenceRepeatedK(s string, k int) (ans string) {
	n := len(s)
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}
	nxt := make([][26]int, n)
	cnt := [26]int{}
	for i := n - 1; i >= 0; i-- {
		nxt[i] = pos
		pos[s[i]-'a'] = i
		cnt[s[i]-'a']++
	}

	// 计算所有可能出现在答案中的字符，包括重复的
	// 倒着统计，这样下面计算排列时的第一个合法方案就是答案，从而提前退出
	a := []byte{}
	for i := 25; i >= 0; i-- {
		for c := cnt[i]; c >= k; c -= k {
			a = append(a, 'a'+byte(i))
		}
	}

	for m := len(a); m > 0 && ans == ""; m-- { // 从大到小枚举答案长度 m
		permutations(len(a), m, func(ids []int) bool { // 枚举长度为 m 的所有排列
			t := make([]byte, m)
			for i, id := range ids {
				t[i] = a[id]
			}
			i, j := 0, 0
			if t[0] == s[0] {
				j = 1
			}
			for {
				i = nxt[i][t[j%m]-'a']
				if i == n {
					break
				}
				j++
			}
			if j >= k*m {
				ans = string(t)
				return true // 提前退出
			}
			return false
		})
	}
	return
}

// 模板：生成 n 选 r 的排列
func permutations(n, r int, do func(ids []int) bool) {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	if do(ids[:r]) {
		return
	}
	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}
	for {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				tmp := ids[i]
				copy(ids[i:], ids[i+1:])
				ids[n-1] = tmp
				cycles[i] = n - i
			} else {
				j := cycles[i]
				ids[i], ids[n-j] = ids[n-j], ids[i]
				if do(ids[:r]) {
					return
				}
				break
			}
		}
		if i == -1 {
			return
		}
	}
}

// tag-[字符串]
// 第四题
// leetcode621: 任务调度器
func leastInterval(tasks []byte, n int) int {
	f := [26]int{}
	maxn := 0
	// 1. 先找出任务数量最大的
	for i := range tasks {
		v := tasks[i] - 'A'
		f[v]++
		maxn = max(maxn, f[v])
	}
	cnt := 0
	// 2. 计算最后一个桶的数量，只有数量和最大的相同，才有可能占用最后一个桶
	for i := range f {
		if f[i] == maxn {
			cnt++
		}
	}
	// 3. 任务很稀疏时，值为任务数量
	return max(len(tasks), cnt+(n+1)*(maxn-1))
}

// tag-[字符串]
// leetcode49: 字母的异位分词
func groupAnagrams(strs []string) [][]string {
	mk := make(map[[26]int][]string)
	var cacl func(s string) [26]int
	cacl = func(s string) [26]int {
		ans := [26]int{}
		for i := range s {
			ans[s[i]-'a']++
		}
		return ans
	}
	for i := range strs {
		v := cacl(strs[i])
		mk[v] = append(mk[v], strs[i])
	}
	ans := make([][]string, 0)
	for _, v := range mk {
		ans = append(ans, v)
	}
	return ans
}
// tag-[字符串]
// leetcode2027：转换字符串的最少操作次数
// 字符串、贪心
func minimumMoves(s string) int {
	n := len(s)
	idx := 0
	cnt := 0
	for idx < n {
		if s[idx] == 'X' {
			cnt++
			idx += 3
		} else {
			idx++
		}
	}
	return cnt
}

// tag-[字符串]
// 第四题
// leetcode43: 字符串相乘
// 数学，字符串
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	ans := make([]int, m+n)
	var t uint8
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			t = (num2[i] - '0') * (num1[j] - '0')
			ans[i+j+1] += int(t % 10)
			c := 0
			if ans[i+j+1] >= 10 {
				ans[i+j+1] %= 10
				c = 1
			}
			carry := int(t/10) + c
			idx := i + j
			for {
				if ans[idx]+carry >= 10 {
					v := ans[idx] + carry
					ans[idx] = v % 10
					carry = v / 10
					idx--
				} else {
					ans[idx] += carry
					break
				}
			}
		}
	}
	b := make([]byte, m+n)
	for i := range ans {
		b[i] = byte(ans[i] + '0')
	}
	i := 0
	for b[i] == '0' {
		i++
	}
	return string(b[i:])
}

func Test_multiply(t *testing.T) {
	fmt.Println(multiply("1234", "4567"))
}
// tag-[字符串]
// leetcode8
func myAtoi(s string) int {
	ss := strings.TrimLeft(s, " ")
	ans := make([]byte, 0)
	for i := range ss {
		v := ss[i]
		if v >= '0' && v <= '9' {
			ans = append(ans, v)
		}
		if v == '-' || v == '+' {
			if len(ans) == 0 {
				ans = append(ans, v)
			} else {
				break
			}
		}
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || v == '.' || v == ' ' {
			break
		}
	}
	var factor int
	if len(ans) > 0 {
		if ans[0] == '-' {
			factor = -1
			ans = ans[1:]
		} else if ans[0] == '+' {
			factor = 1
			ans = ans[1:]
		} else {
			factor = 1
		}
	}
	res := 0
	t := 1
	for i := range ans {
		res += int(ans[len(ans)-i-1]-'0') * t
		if factor == 1 && (res > math.MaxInt32 || t > math.MaxInt32) {
			return math.MaxInt32
		}
		if factor == -1 && (res < math.MinInt32 || t < math.MinInt32) {
			return math.MinInt32
		}
		t *= 10
	}
	r := factor * res
	if r > math.MaxInt32 {
		return math.MaxInt32
	}
	if r < math.MinInt32 {
		return math.MinInt32
	}
	return r
}

func Test_myAtoi(t *testing.T) {
	fmt.Println(myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
}

// leetcode8 优化解法
func myAtoi1(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	// 丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	// 标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')  // 字节 byte '0' == 48
		if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}

func Test_forsum(t *testing.T) {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
// tag-[字符串]
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

// tag-[字符串]
// leetcode6:字符串（模拟）
func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 {
		return s
	}
	numCols := n / (4 * (numRows - 1)) * 2 * (numRows - 1)
	left := n % (4 * (numRows - 1))
	if left >= 3*numRows-2 {
		numCols += numRows + (left - 3*numRows + 2)
	} else if left > 2*numRows-2 {
		numCols += numRows
	} else if left >= numRows {
		numCols += 1 + (left - numRows)
	} else {
		numCols += 1
	}
	matrix := make([][]byte, numRows)
	for i := range matrix {
		matrix[i] = make([]byte, numCols)
	}
	row, col := 0, 0
	up := false
	for i := range s {
		matrix[row][col] = s[i]
		if row == numRows-1 {
			up = true
		} else if row == 0 {
			up = false
		}
		if !up {
			row++
		} else {
			row--
			col++
		}
	}
	res := make([]byte, 0, n)
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if matrix[i][j] != 0 {
				res = append(res, matrix[i][j])
			}
		}
	}
	return string(res)
}

// tag-[字符串]
// leetcode6:优化解法，无需计算列的个数
func convert_(s string, numRows int) string {
	n := len(s)
	if numRows == 1 {
		return s
	}
	maxRow := min(numRows, n)
	res := make([][]byte, maxRow)
	row := 0
	up := false
	for i := range s {
		res[row] = append(res[row], s[i])
		if row == numRows-1 || row == 0 {
			up = !up
		}
		if !up {
			row++
		} else {
			row--
		}
	}
	ans := make([]string, 0, n)
	for i := 0; i < numRows; i++ {
		ans = append(ans, string(res[i]))
	}
	return strings.Join(ans, "")
}

func Test_convert_(t *testing.T) {
	fmt.Println(convert("A", 2))
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 5))
	fmt.Println(convert("PAYPALISHIRING", 6))
	fmt.Println(convert("PAYPALISHIRING", 7))
	fmt.Println(convert("PAYPALISHIRING", 8))
	fmt.Println(convert("PAYPALISHIRING", 9))
}

// tag-[字符串]
// leetcode481:神奇字符串（模拟）
func magicalString_(n int) int {
	if n < 3 {
		return 1
	}
	t := make([]byte, 3, n)
	t[0] = '1'
	t[1] = '2'
	t[2] = '2'
	cnt := 1
	fast, slow := 2, 1
	for len(t) < n {
		slow++
		if t[slow] == '2' {
			if t[fast] == '2' {
				t = append(t, []byte{'1', '1'}...)
				if len(t) > n {
					cnt += 1
				} else {
					cnt += 2
				}
			}
			if t[fast] == '1' {
				t = append(t, []byte{'2', '2'}...)
			}
			fast += 2
		} else {
			if t[fast] == '2' {
				t = append(t, '1')
				cnt += 1
			}
			if t[fast] == '1' {
				t = append(t, '2')
			}
			fast += 1
		}
	}
	return cnt
}

func Test_magicalString_(t *testing.T) {
	fmt.Println(magicalString_(100))
}

// tag-[字符串]
// leetcode71: 简化路径
func simplifyPath(path string) string {
	ss := strings.Split(path, "/")
	newPath := make([]string, 0)
	for _, s := range ss {
		if s == "." || s == "" {
			continue
		} else if s == ".." {
			if len(newPath) > 0 {
				newPath = newPath[:len(newPath)-1]
			}
		} else {
			newPath = append(newPath, s)
		}
	}
	return "/" + strings.Join(newPath, "/")
}

func Test_simplifyPath(t *testing.T) {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/../"))
}

// tag-[字符串]
// leetcode周赛第一题
func countValidWords(sentence string) int {
	ss := strings.Split(sentence, " ")
	isChar := func(v byte) bool {
		return v >= 'a' && v <= 'z'
	}
	isPunctuation := func(v byte) bool {
		return v == '!' || v == '.' || v == ','
	}
	isValid := func(s string) bool {
		cnt1 := 0
		cnt2 := 0
		for j := range s {
			v := s[j]
			if !isChar(v) && v != '-' && !isPunctuation(v) {
				return false
			}
			if v == '-' {
				if cnt1 > 0 {
					return false
				}
				cnt1++
				if j > 0 && j < len(s)-1 && isChar(s[j-1]) && isChar(s[j+1]) {
					continue
				}
				return false
			}
			if isPunctuation(v) {
				if cnt2 > 0 {
					return false
				}
				cnt2++
				if j != len(s)-1 {
					return false
				}
			}
		}
		return true
	}
	cnt := 0
	for i := range ss {
		if ss[i] == "" || ss[i] == " " {
			continue
		}
		if isValid(ss[i]) {
			cnt++
		}
	}
	return cnt
}

func Test_countValidWords(t *testing.T) {
	fmt.Println(countValidWords("alice and  bob are playing stone-game10"))
	fmt.Println(countValidWords("he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."))
	fmt.Println(countValidWords("!this  1-s b8d!"))
	fmt.Println(countValidWords("cat and  dog"))
	fmt.Println(countValidWords("!this  a-s- bad!"))
}

// tag-[字符串]
// leetcode43: 字符串相乘
// 重刷
func multiply_(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	res := make([]byte, n1+n2)
	for i := range res {
		res[i] = '0'
	}
	for i := n1 - 1; i >= 0; i-- {
		v1 := int(num1[i] - '0')
		carry := 0
		for j := n2 - 1; j >= 0; j-- {
			v2 := int(num2[j] - '0')
			v := v1*v2 + int(res[i+j+1]-'0') + carry
			res[i+j+1] = byte(v%10 + '0')
			carry = v / 10
		}
	}
	return strings.TrimLeft(string(res), "0")
}

func Test_multiply_(t *testing.T) {
	fmt.Println(multiply_("2", "3"))
}

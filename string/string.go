package string

import "strings"

// 无重复字符的最长子串

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	res := 1
	left, right := 0, 1
	for right < length {
		for i := right - 1; i >= left; i-- {
			if s[i] == s[right] {
				left = i + 1
				break
			}
		}
		if nv := right - left + 1; nv > res {
			res = nv
		}
		right++
	}
	return res
}

// 替换后的最长重复字符

// 给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
func characterReplacement(s string, k int) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	res := 1
	m := make(map[byte]int)
	left, right := 0, 0
	m[s[left]]++

	getRest := func(m map[byte]int, sum int) int {
		var max int
		for _, v := range m {
			if max < v {
				max = v
			}
		}
		return sum - max
	}
	for left <= right {
		val := getRest(m, right-left+1)
		if val > k {
			m[s[left]]--
			left++
		} else {
			if nv := right - left + 1; nv > res {
				res = nv
			}
			right++
			if right == length {
				break
			}
			m[s[right]]++
		}
	}
	return res
}

// 最小覆盖子串

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
func minWindow(s string, t string) string {
	isok := func(m map[byte]int) bool {
		for _, v := range m {
			if v > 0 {
				return false
			}
		}
		return true
	}
	m := make(map[byte]int)
	for _, v := range t {
		m[byte(v)]++
	}
	left, right := 0, -1
	l, r, min := left, right, len(s)
	for right < len(s) {
		if isok(m) {
			if v := right - left; min > v {
				min = v
				l = left
				r = right
			}
			if _, ok := m[s[left]]; ok {
				m[s[left]]++
			}
			left++
		} else {
			right++
			if right == len(s) {
				break
			}
			if _, ok := m[s[right]]; ok {
				m[s[right]]--
			}
		}
	}
	return s[l : r+1]
}

// 有效的字母异位词

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
func isAnagram(s string, t string) bool {
	m := make(map[byte]int)
	for _, v := range s {
		m[byte(v)]++
	}

	for _, v := range t {
		m[byte(v)]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// 字母异位词分组

// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
func groupAnagrams(strs []string) [][]string {
	nums := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	m := make(map[int][]string)
	for _, str := range strs {
		key := 1
		for _, v := range str {
			key *= nums[byte(v)-'a']
		}
		m[key] = append(m[key], str)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 有效的括号

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
func isValid(s string) bool {
	m := map[byte]byte{'}': '{', ')': '(', ']': '['}
	var stack []byte
	for _, v := range s {
		c := byte(v)
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')', '}', ']':
			if length := len(stack); length == 0 || stack[length-1] != m[c] {
				return false
			} else {
				stack = stack[:length-1]
			}
		default:
			return false
		}
	}
	return len(stack) == 0
}

// 验证回文串

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
// 说明：本题中，我们将空字符串定义为有效的回文串。
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for {
		for left < right {
			if (s[left] >= 'a' && s[left] <= 'z') || (s[left] >= '0' && s[left] <= '9') {
				break
			} else {
				left++
			}
		}

		for left < right {
			if (s[right] >= 'a' && s[right] <= 'z') || (s[right] >= '0' && s[right] <= '9') {
				break
			} else {
				right--
			}
		}
		if left >= right {
			break
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 最长回文子串

// 给你一个字符串 s，找到 s 中最长的回文子串
func longestPalindrome(s string) string {
	var res string
	length := len(s)
	if length == 0 {
		return res
	}
	res = s[:1]
loop1:
	for i := 0; i < length; i++ {
		left, right := i-1, i+1
		for left >= 0 && right < length {
			if s[left] != s[right] {
				continue loop1
			}
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}
loop2:
	for i := 0; i < length-1; i++ {
		left, right := i, i+1
		for left >= 0 && right < length {
			if s[left] != s[right] {
				continue loop2
			}
			if right-left+1 > len(res) {
				res = s[left : right+1]
			}
			left--
			right++
		}
	}
	return res
}

func countSubstrings(s string) int {
	length := len(s)
	res := length

	for i := 1; i < length-1; i++ {
		left, right := i-1, i+1
		for left >= 0 && right < length {
			if s[left] == s[right] {
				res++
				left--
				right++
			} else {
				break
			}
		}
	}

	for i := 0; i < length-1; i++ {
		left, right := i, i+1
		for left >= 0 && right < length {
			if s[left] == s[right] {
				res++
				left--
				right++
			} else {
				break
			}
		}
	}
	return res
}

package str

import "strings"

//找最长的不重复字符串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	left, right := 0, 1
	var res int

	for right < len(s) {
		for i := right - 1; i >= left; i-- {
			if s[right] == s[i] {
				left = i + 1
				break
			}
		}
		if n := right - left + 1; n > res {
			res = n
		}
		right++
	}
	return res
}

// 给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	left, right, max, maxchar, res := 0, 1, 1, s[0], 1

	m := make(map[byte]int)
	m[s[0]]++

	for right < len(s) {
		m[s[right]]++
		if max < m[s[right]] {
			max = m[s[right]]
			maxchar = s[right]
		}
		if right-left+1-max > k {
			m[s[left]]--
			if s[left] == maxchar {
				max--
			}
			left++
		} else {
			if n := right - left + 1; n > res {
				res = n
			}
		}
		right++
	}
	return res
}

// 给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。
func minWindow(s string, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	left, right, res, l, r, f, hasone := 0, 0, len(s), 0, 0, true, false
	for right < len(s) {
		if _, ok := m[s[right]]; ok && f {
			m[s[right]]--
		}
		f = false

		isok := true
		for _, v := range m {
			if v > 0 {
				isok = false
			}
		}
		if isok {
			hasone = true
			if res > right-left {
				res = right - left
				l = left
				r = right
			}
			if _, ok := m[s[left]]; ok {
				m[s[left]]++
			}
			left++
		} else {
			f = true
			right++
		}
	}

	if !hasone {
		return ""
	}

	return s[l : r+1]
}

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func getMul(s string) int {
	slice := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	res := 1
	for i := 0; i < len(s); i++ {
		res *= slice[s[i]-'a']
	}
	return res
}

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[int][]string)
	for i := 0; i < len(strs); i++ {
		mul := getMul(strs[i])
		m[mul] = append(m[mul], strs[i])
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res

}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
func isValid(s string) bool {
	slice := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			slice = append(slice, s[i])
		} else {
			if len(slice) == 0 {
				return false
			}
			end := slice[len(slice)-1]
			switch s[i] {
			case ')':
				if end != '(' {
					return false
				}
			case '}':
				if end != '{' {
					return false
				}
			case ']':
				if end != '[' {
					return false
				}
			}
			slice = slice[:len(slice)-1]
		}
	}
	if len(slice) != 0 {
		return false
	}
	return true
}

//判断是否是回文串
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var slice []byte
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			slice = append(slice, s[i])
		}
	}
	for i := 0; i < len(slice)/2; i++ {
		if slice[i] != slice[len(slice)-i-1] {
			return false
		}
	}
	return true
}

// //最长的回文串
// func longestPalindrome(s string) string {
// 	length := len(s)
// 	for i := length; i >= 1; i-- {
// 	loop:
// 		for j := 0; j < length-i+1; j++ {
// 			for k := j; k < j+i/2; k++ {
// 				if s[k] != s[j+j+i-1-k] {
// 					continue loop
// 				}
// 			}
// 			return s[j : j+i]
// 		}
// 	}
// 	return ""
// }

//最长的回文串
//dp 解法
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var res, l, r int
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		for j := 0; j < i; j++ {
			if s[j] == s[i] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				if res < i-j {
					res = i - j
					l = j
					r = i
				}
			}
		}
	}
	return s[l : r+1]
}

//数回文字符串的数量
func countSubstrings(s string) int {
	var res int
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		res++
		for j := 0; j < i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				res++
			}
		}
	}
	return res
}

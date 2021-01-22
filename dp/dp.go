package dp

// 爬楼梯

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func climbStairs(n int) int {
	a, b := 1, 2
	for n > 1 {
		a, b = b, a+b
		n--
	}
	return a
}

// 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

// 你可以认为每种硬币的数量是无限的
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			nv := i - v
			if nv < 0 {
				continue
			} else if nv == 0 {
				dp[i] = 1
			} else {
				if dp[nv] > 0 {
					if x := dp[nv] + 1; x < dp[i] || dp[i] == 0 {
						dp[i] = x
					}
				}
			}
		}
	}
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

// 最长递增子序列

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
func lengthOfLIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	var res int
	for i := 0; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				v := dp[j] + 1
				if v > dp[i] {
					dp[i] = v
				}
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

// 最长公共子序列

// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
// 若这两个字符串没有公共子序列，则返回 0。
func longestCommonSubsequence(text1 string, text2 string) int {
	length1 := len(text1) + 1
	length2 := len(text2) + 1
	dp := make([][]int, length1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, length2)
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}
	return dp[length1-1][length2-1]
}

// 单词拆分

// 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
func wordBreak(s string, wordDict []string) bool {
	length := len(s)
	dp := make([]bool, length)
	for i := 0; i < length; i++ {
		for _, v := range wordDict {
			l := len(v)
			if l > i+1 {
				continue
			}
			if s[i+1-l:i+1] == v && (i+1-l == 0 || dp[i-l]) {
				dp[i] = true
			}
		}
	}
	return dp[length-1]
}

// 组合总和 Ⅳ

// 给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 0; i < len(dp); i++ {
		for _, v := range nums {
			if i < v {
				continue
			} else if i == v {
				dp[i]++
			} else {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}

// 打家劫舍

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
func rob2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := make([]int, length)
	if length > 0 {
		dp[0] = nums[0]
	}

	if length > 1 {
		max := nums[0]
		if nums[0] < nums[1] {
			max = nums[1]
		}
		dp[1] = max
	}
	for i := 2; i < length; i++ {
		v := nums[i] + dp[i-2]
		if v > dp[i-1] {
			dp[i] = v
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[length-1]
}

// 打家劫舍 II

// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	v1 := rob2(nums[1:])
	v2 := rob2(nums[:length-1])
	if v1 > v2 {
		return v1
	}
	return v2
}

// 解码方法

// 一条包含字母 A-Z 的消息通过以下方式进行了编码：

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。

// 题目数据保证答案肯定是一个 32 位的整数。
func numDecodings(s string) int {
	length := len(s)
	dp := make([]int, length)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	if length > 1 {
		switch s[0] {
		case '1':
			switch s[1] {
			case '0':
				dp[1] = 1
			default:
				dp[1] = 2
			}
		case '2':
			switch {
			case s[1] == '0':
				dp[1] = 1
			case s[1] >= '1' && s[1] <= '6':
				dp[1] = 2
			default:
				dp[1] = 1
			}
		default:
			switch s[1] {
			case '0':
				return 0
			default:
				dp[1] = 1
			}
		}
	}
	for i := 2; i < length; i++ {
		switch s[i-1] {
		case '0':
			switch s[i] {
			case '0':
				return 0
			default:
				dp[i] = dp[i-1]
			}
		case '1':
			switch s[i] {
			case '0':
				dp[i] = dp[i-2]
			default:
				dp[i] = dp[i-1] + dp[i-2]
			}
		case '2':
			switch {
			case s[i] == '0':
				dp[i] = dp[i-2]
			case s[i] >= '1' && s[i] <= '6':
				dp[i] = dp[i-1] + dp[i-2]
			default:
				dp[i] = dp[i-1]
			}
		default:
			switch s[i] {
			case '0':
				return 0
			default:
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[length-1]
}

// 不同路径

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 跳跃游戏

// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个位置。
func canJump(nums []int) bool {
	length := len(nums)
	dp := make([]bool, length)
	dp[0] = true
	for i := 0; i < length; i++ {
		if dp[i] {
			for j := 1; j <= nums[i] && i+j < length; j++ {
				dp[i+j] = true
			}
		}
	}
	return dp[length-1]
}

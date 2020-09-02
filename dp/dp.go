package dp

//爬楼梯
func ClimbStairs(n int) int {
	a, b := 1, 2
	for n > 1 {
		a, b = b, a+b
		n--
	}
	return a
}

//零钱兑换

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		for _, v := range coins {
			if i-v >= 0 && (i-v == 0 || dp[i-v] > 0) && (dp[i] == 0 || dp[i] > dp[i-v]+1) {
				dp[i] = dp[i-v] + 1
			}
		}
	}
	res := dp[amount]
	if res == 0 {
		return -1
	}
	return res
}

//寻找最长升序子序列
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	res := 1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

//给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
func WordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for _, v := range wordDict {
			if i >= len(v) && s[i-len(v):i] == v && dp[i-len(v)] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

//给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		for _, v := range nums {
			if i-v >= 0 && dp[i-v] > 0 {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
func rob(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(dp); i++ {
		switch {
		case i == 0:
			dp[i] = nums[i]
		case i == 1:
			dp[i] = max(nums[i], nums[i-1])
		case i > 1:
			dp[i] = max(nums[i]+dp[i-2], dp[i-1])
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 思路：先得到去头和去尾两个数组，然后用dp的思想各求最大值，然后对这两个值取最大值
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

// 一条包含字母 A-Z 的消息通过以下方式进行了编码：

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。

//思路：先求前两个数 ，后面用dp思想，搞法略复杂
func NumDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s))

	if len(s) >= 1 {
		if s[0] == 48 {
			return 0
		}
		dp[0] = 1
	}

	if len(s) >= 2 {
		if s[1]-48 == 0 {
			if s[0]-48 > 2 {
				return 0
			}
			dp[1] = 1
		} else if (s[0]-48)*10+(s[1]-48) > 26 {
			dp[1] = 1
		} else {
			dp[1] = 2
		}
	}

	for i := 2; i < len(dp); i++ {
		if s[i-1]-48 == 0 {
			if s[i]-48 == 0 {
				return 0
			}
			dp[i] = dp[i-1]
		} else {
			if s[i]-48 == 0 {
				if s[i-1]-48 <= 2 {
					dp[i] = dp[i-2]
				} else {
					return 0
				}
			} else if (s[i-1]-48)*10+(s[i]-48) > 26 {
				dp[i] = dp[i-1]
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		}
	}
	return dp[len(s)-1]
}

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

// 问总共有多少条不同的路径？
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
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

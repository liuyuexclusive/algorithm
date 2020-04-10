package maxprofit

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

// 注意：你不能在买入股票前卖出股票。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//暴力法
func MaxProfit1(prices []int) int {
	var res int
	for i := 0; i < len(prices); i++ {
		for j := 0; j < len(prices)-i-1; j++ {
			if v := prices[i+j+1] - prices[i]; v > res {
				res = v
			}
		}
	}
	return res
}

//寻找最小值
func MaxProfit2(prices []int) int {
	var min, res int
	if len(prices) == 0 {
		return 0
	}
	min = prices[0]

	for i := 0; i < len(prices); i++ {
		val := prices[i]
		if min > val {
			min = val
		}
		if newVal := val - min; res < newVal {
			res = newVal
		}
	}
	return res
}

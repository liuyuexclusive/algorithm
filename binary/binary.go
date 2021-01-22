package binary

//两整数之和

//不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
func getSum(a int, b int) int {
	for b != 0 {
		temp := a
		a = a ^ b
		b = temp & b << 1
	}
	return a
}

//位1的个数

//编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
func hammingWeight(num uint32) int {
	var res uint32
	for num != 0 {
		res += (num & 1)
		num >>= 1
	}
	return int(res)
}

//比特位计数

//给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
func countBits(num int) []int {
	dp := make([]int, num+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}

//丢失的数字

//给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
func missingNumber(nums []int) int {
	var res int
	for k, v := range nums {
		res ^= k ^ v
	}
	res ^= len(nums)
	return res
}

//颠倒二进制位

//颠倒给定的 32 位无符号整数的二进制位。
func reverseBits(num uint32) uint32 {
	var res uint32
	var n int = 31
	for n >= 0 {
		res += num & 1 << n
		num >>= 1
		n--
	}
	return res
}

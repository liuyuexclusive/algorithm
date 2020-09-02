package binary

//两数之和
func GetSum(a int, b int) int {
	for b != 0 {
		temp := a
		a = a ^ b
		b = temp & b << 1
	}
	return a
}

//位1的个数
func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

//0<=x<=num中所有树1的个数数组
func CountBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i < len(res); i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

//找出缺失的数字
func missingNumber(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res ^= i ^ nums[i]
	}
	res ^= len(nums)

	return res
}

//翻转二进制
func reverseBits(num uint32) uint32 {
	digits := 31
	var res uint32
	for digits >= 0 {
		res += num & 1 << digits
		num >>= 1
		digits--
	}
	return res
}

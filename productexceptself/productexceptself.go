package productexceptself

// ProductExceptSelf 除自己外的乘积 左*右
func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	res, left, right := make([]int, length), make([]int, length), make([]int, length)

	right[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	left[0] = 1
	for i := 1; i < length; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	for i := 0; i < len(res); i++ {
		res[i] = left[i] * right[i]
	}

	return res
}

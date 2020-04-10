package maxsubarray

import "math"

// MaxSubArray 最大连续和
func MaxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	res := math.MinInt64
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if res < sum {
			res = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return res
}

package maxpoduct

// MaxProduct 最大连续乘积
func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res, imax, imin := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		val := nums[i]
		if val < 0 {
			imax, imin = imin, imax
		}
		if newval := imax * val; newval > val {
			imax = newval
		} else {
			imax = val
		}
		if newval := imin * val; newval < val {
			imin = newval
		} else {
			imin = val
		}

		if res < imax {
			res = imax
		}
	}

	return res
}

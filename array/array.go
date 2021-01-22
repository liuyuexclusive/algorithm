package array

// 两数之和

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if i0, ok := m[target-v]; ok {
			return []int{i0, i}
		}
		m[v] = i
	}
	return nil
}

// 买卖股票的最佳时机

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
// 注意：你不能在买入股票前卖出股票。
func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	var res int
	min := prices[0]
	for i := 1; i < length; i++ {
		v := prices[i]
		if nv := v - min; nv > res {
			res = nv
		}
		if v < min {
			min = v
		}
	}
	return res
}

// 存在重复元素

// 给定一个整数数组，判断是否存在重复元素。
// 如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}

//除自身以外数组的乘积

// 给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
func productExceptSelf(nums []int) []int {
	length := len(nums)
	left, right, res := make([]int, length), make([]int, length), make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			left[i] = 1
		} else {
			left[i] = left[i-1] * nums[i-1]
		}
	}
	for i := length - 1; i >= 0; i-- {
		if i == length-1 {
			right[i] = 1
		} else {
			right[i] = right[i+1] * nums[i+1]
		}
		res[i] = left[i] * right[i]
	}
	return res
}

//最大子序和

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	res, left := nums[0], nums[0]
	for i := 1; i < length; i++ {
		if left < 0 {
			left = 0
		}
		v := nums[i]
		if nv := left + v; nv > res {
			res = nv
		}
		left += v
	}
	return res
}

//乘积最大子数组

// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积
func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	res, min, max := nums[0], nums[0], nums[0]

	for i := 1; i < length; i++ {
		v := nums[i]
		if v < 0 {
			min, max = max, min
		}
		if nv := min * v; nv < v {
			min = nv
		} else {
			min = v
		}

		if nv := max * v; nv > v {
			max = nv
		} else {
			max = v
		}
		if max > res {
			res = max
		}
	}

	return res
}

//寻找旋转排序数组中的最小值

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。
// 请找出其中最小的元素。
func findMin(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	for i := 1; i < length; i++ {
		if v := nums[i]; v < nums[i-1] {
			return v
		}
	}
	return nums[0]
}

//搜索旋转排序数组

// 升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
// 请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1
func search(nums []int, target int) int {
	length := len(nums)
	var offset int

	for i := 1; i < length; i++ {
		if nums[i] < nums[i-1] {
			offset = i
		}
	}

	newnums := make([]int, length)

	for i := 0; i < length; i++ {
		newnums[i] = nums[(i+offset)%length]
	}

	left, right, res := 0, length, -1
loop:
	for {
		mid := (left + right) / 2
		midVal := newnums[mid]
		switch {
		case target < midVal:
			right = mid
			if (left+right)/2 == mid {
				break loop
			}
		case target > midVal:
			left = mid
			if (left+right)/2 == mid {
				break loop
			}
		default:
			res = mid
			break loop
		}
	}
	if res == -1 {
		return res
	}
	return (res + offset) % length
}

//三数之和

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
func threeSum(nums []int) [][]int {
	var res [][]int
	m := make(map[[3]int]bool)
	length := len(nums)
	for i := length / 2; i >= 1; i /= 2 {
		for j := 0; j < i; j++ {
			for k := j + i; k < length; k += i {
				for l := k; l > j; l -= i {
					if nums[l] < nums[l-i] {
						nums[l], nums[l-i] = nums[l-i], nums[l]
					} else {
						break
					}
				}
			}
		}
	}

	for i := 0; i < length-2; i++ {
		left, right := i+1, length-1
		for left < right {
			v := nums[i] + nums[left] + nums[right]
			switch {
			case v < 0:
				left++
			case v > 0:
				right--
			default:
				if !m[[3]int{nums[i], nums[left], nums[right]}] {
					res = append(res, []int{nums[i], nums[left], nums[right]})
					m[[3]int{nums[i], nums[left], nums[right]}] = true
				}
				left++
				right--
			}
		}
	}
	return res
}

//盛最多水的容器

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
func maxArea(height []int) int {
	length := len(height)
	if length < 2 {
		return 0
	}
	left, right := 0, length-1
	var res int
	for left < right {
		min := height[left]
		if v := height[right]; v < min {
			min = v
		}
		if nv := (right - left) * min; nv > res {
			res = nv
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return res
}

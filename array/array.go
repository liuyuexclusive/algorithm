package array

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//希尔排序
func hillSort(slice []int) {
	for i := len(slice) / 2; i >= 1; i /= 2 {
		for j := 0; j < i; j++ {
			for k := j + i; k < len(slice); k += i {
				for l := k; l > j; l -= i {
					if slice[l] < slice[l-i] {
						swap(slice, l, l-i)
					} else {
						break
					}
				}
			}
		}
	}
}

//两数之和
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 2)
	for i, v := range nums {
		if i2, ok := m[target-v]; ok {
			res[0] = i2
			res[1] = i
			return res
		}
		m[v] = i
	}
	return res
}

//股票最大利润
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minVal, res := prices[0], 0
	for i := 1; i < len(prices); i++ {
		v := prices[i]
		if v < minVal {
			minVal = v
			continue
		}
		res = max(res, v-minVal)
	}
	return res
}

//判断是否有重复
func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}

//除自己外的数组乘积
func ProductExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	left, right, res := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	left[0] = 1
	right[len(right)-1] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
		right[len(nums)-i-1] = right[len(nums)-i] * nums[len(nums)-i]
	}
	for i := 0; i < len(nums); i++ {
		res[i] = left[i] * right[i]
	}
	return res
}

//最大子序和
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, res := max(0, nums[0]), nums[0]

	for i := 1; i < len(nums); i++ {
		v := nums[i]
		res = max(res, v+sum)
		if v+sum < 0 {
			sum = 0
		} else {
			sum += v
		}
	}
	return res
}

//最大乘积
func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res, m1, m2 := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		v := nums[i]
		if v < 0 {
			m1, m2 = m2, m1
		}
		if m1*v < v {
			m1 = m1 * v
		} else {
			m1 = v
		}
		if m2*v > v {
			m2 = m2 * v
		} else {
			m2 = v
		}
		if res < m2 {
			res = m2
		}
	}
	return res
}

//排序旋转数组寻找最小值
func FindMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]
}

//排序旋转数组查找一个值
func Search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	var offset int
	for i := 1; i < length; i++ {
		if nums[i] < nums[i-1] {
			offset = i
			break
		}
	}
	newNums := make([]int, length)

	for i := 0; i < len(newNums); i++ {
		newNums[i] = nums[(i+offset)%length]
	}

	index := binarySearch(newNums, target)
	if index == -1 {
		return -1
	}
	return (index + offset) % length
}

func binarySearch(slice []int, target int) int {
	left, right := 0, len(slice)
loop:
	for left < right {
		middle := (left + right) / 2
		middleVal := slice[middle]
		switch {
		case middleVal > target:
			right = middle
			if (left+right)/2 == middle {
				break loop
			}
		case middleVal < target:
			left = middle
			if (left+right)/2 == middle {
				break loop
			}
		default:
			return middle
		}
	}
	return -1
}

//三数之和
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	var res [][]int
	dupicated := make(map[[3]int]bool)
	hillSort(nums)
	for i := 0; i < len(nums)-2; i++ {
		v1 := nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			v2 := nums[left]
			v3 := nums[right]
			switch {
			case v1+v2+v3 == 0:
				arr := [3]int{v1, v2, v3}
				if _, ok := dupicated[arr]; !ok {
					res = append(res, []int{v1, v2, v3})
					dupicated[arr] = true
				}
				left++
				right--
			case v1+v2+v3 < 0:
				left++
			case v1+v2+v3 > 0:
				right--
			}
		}
	}
	return res
}

//盛最多水的容器
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var res int
	for left < right {
		v1 := height[left]
		v2 := height[right]
		res = max(res, min(v1, v2)*(right-left))
		if v1 > v2 {
			right--
		} else {
			left++
		}
	}
	return res
}

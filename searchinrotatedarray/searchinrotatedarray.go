package searchinrotatedarray

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

// 你可以假设数组中不存在重复的元素。

// 你的算法时间复杂度必须是 O(log n) 级别。

// 示例 1:

// 输入: nums = [4,5,6,7,0,1,2], target = 0
// 输出: 4
// 示例 2:

// 输入: nums = [4,5,6,7,0,1,2], target = 3
// 输出: -1

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func search(nums []int, target int) int {
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
	newSlice := make([]int, length)
	for i := 0; i < length; i++ {
		newSlice[(i-offset+length)%length] = nums[i]
	}
	index := binarySearch(newSlice, target)
	if index == -1 {
		return index
	}

	return (index + offset) % length
}

func binarySearch(nums []int, target int) int {
	res, left, right := -1, 0, len(nums)

loop:
	for {
		if left >= right {
			break
		}
		middle := (left + right) / 2
		middleVal := nums[middle]

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
		case middleVal == target:
			res = middle
			break loop
		}

	}

	return res
}

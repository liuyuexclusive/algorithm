// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

//

// 示例：

// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

// 满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package threesum

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	//排序
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	p1, p2, p3 := 0, 1, len(nums)-1

	for {
		if nums[p1] > 0 || p1 == len(nums)-2 {
			break
		}
		if p3-p2 <= 0 {
			p1++
			if nums[p1] == nums[p1-1] {
				continue
			}
			p2 = p1 + 1
			p3 = len(nums) - 1
			continue
		}

		val := nums[p1] + nums[p2] + nums[p3]

		switch {
		case val < 0:
			p2++
			for nums[p2] == nums[p2-1] && p2 < p3 {
				p2++
			}
		case val > 0:
			p3--
			for nums[p2] == nums[p3+1] && p2 < p3 {
				p2++
			}
		case val == 0:
			res = append(res, []int{nums[p1], nums[p2], nums[p3]})
			p2++
			for nums[p2] == nums[p2-1] && p2 < p3 {
				p2++
			}
			p3--
			for nums[p2] == nums[p3+1] && p2 < p3 {
				p2++
			}
		}

	}

	return res
}

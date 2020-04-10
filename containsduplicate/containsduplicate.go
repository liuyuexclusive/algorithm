package containsduplicate

// ContainsDuplicate 数组是否包含重复
func ContainsDuplicate(nums []int) bool {
	var res bool
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] > 0 {
			res = true
			break
		}
		m[nums[i]]++
	}
	return res
}

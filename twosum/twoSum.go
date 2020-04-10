package twosum

//暴力法
func TwoSum1(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if i1, i2 := i, i+j+1; nums[i1]+nums[i2] == target {
				res = []int{i1, i2}
				break
			}
		}
	}
	return res
}

//map
func TwoSum2(nums []int, target int) []int {
	var res []int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		v1 := nums[i]
		if v2, ok := m[target-v1]; ok && v2 != i {
			res = []int{i, v2}
			break
		}
	}
	return res
}

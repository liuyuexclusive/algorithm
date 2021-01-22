package interval

import "sort"

// 插入区间

// 给出一个无重叠的 ，按照区间起始端点排序的区间列表。
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		v := intervals[i]
		last := res[len(res)-1]
		if last[1] < v[0] {
			res = append(res, v)
		} else {
			if last[1] < v[1] {
				last[1] = v[1]
			}
		}
	}
	return res
}

// 合并区间

//给出一个区间的集合，请合并所有重叠的区间。
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		v := intervals[i]
		last := res[len(res)-1]
		if last[1] < v[0] {
			res = append(res, v)
		} else {
			if last[1] < v[1] {
				last[1] = v[1]
			}
		}
	}
	return res
}

// 无重叠区间
// 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

// 注意:

// 可以认为区间的终点总是大于它的起点。
// 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	if len(intervals) == 0 {
		return 0
	}
	last := intervals[0]
	var res int
	for i := 1; i < len(intervals); i++ {
		v := intervals[i]
		if v[0] < last[1] {
			res++
		} else {
			last[0] = v[0]
			last[1] = v[1]
		}
	}
	return res
}

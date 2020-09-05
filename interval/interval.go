package interval

import "sort"

// 插入一个区间，融合以后得到新的无重叠区间
func Insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	a, b := newInterval[0], newInterval[1]
	var isInserted bool
	for _, v := range intervals {
		if v[1] >= a && v[0] <= b {
			if v[0] <= newInterval[0] && v[1] >= newInterval[1] {
				return intervals
			} else if v[0] <= newInterval[0] {
				a = v[0]
			} else if v[1] >= newInterval[1] {
				b = v[1]
			}
		} else {
			if v[0] > a && !isInserted {
				res = append(res, []int{a, b})
				isInserted = true
			}
			res = append(res, v)
		}
	}
	if !isInserted {
		res = append(res, []int{a, b})
	}
	return res
}

// 合并区间，得到无重叠区间
func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	for i := 1; i < len(intervals); i++ {
		for j := i; j > 0; j-- {
			if intervals[j][0] < intervals[j-1][0] {
				intervals[j], intervals[j-1] = intervals[j-1], intervals[j]
			} else {
				break
			}
		}
	}
	var res [][]int
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > pre[1] {
			res = append(res, pre)
			pre = intervals[i]
		} else {
			pre = []int{pre[0], max(intervals[i][1], pre[1])}
		}
	}
	res = append(res, pre)
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//抹去区间(最小次数)，得到无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {

	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var res int
	i := 1
	for i < len(intervals) {
		if intervals[i][0] < intervals[i-1][1] {
			intervals = append(intervals[:i], intervals[i+1:]...)
			res++
			continue
		}
		i++
	}
	return res
}

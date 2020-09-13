package heap

type ListNode struct {
	Val  int
	Next *ListNode
}

func rise(slice []*ListNode, k, end int) {
	for {
		i := 2*k + 1
		if i > end {
			break
		}
		if i < end && slice[i].Val > slice[i+1].Val {
			i++
		}
		if slice[k].Val <= slice[i].Val {
			break
		}
		slice[k], slice[i] = slice[i], slice[k]
		k = i
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	for i := len(lists) - 1; i >= 0; i-- {
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
		}
	}
	res := &ListNode{}
	p := res
	end := len(lists) - 1
	for i := end / 2; i >= 0; i-- {
		rise(lists, i, end)
	}
	for end >= 0 {
		p.Next = lists[0]
		p = p.Next
		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists[0] = lists[end]
			end--
		}
		if end == 0 {
			p.Next = lists[0]
			break
		}
		rise(lists, 0, end)
	}
	return res.Next
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var slice []int
	for k := range m {
		slice = append(slice, k)
	}
	end := len(slice) - 1
	for i := end / 2; i >= 0; i-- {
		rise2(slice, i, end, m)
	}
	var res []int
	for i := 0; i < k; i++ {
		slice[0], slice[end] = slice[end], slice[0]
		res = append(res, slice[end])
		end--
		rise2(slice, 0, end, m)
	}
	return res
}

func rise2(slice []int, k, end int, m map[int]int) {
	for {
		i := 2*k + 1
		if i > end {
			break
		}
		if i < end && m[slice[i+1]] > m[slice[i]] {
			i++
		}
		if m[slice[k]] >= m[slice[i]] {
			break
		}
		slice[i], slice[k] = slice[k], slice[i]
		k = i
	}
}

//数据流的中位数
type MedianFinder struct {
	min []int
	max []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func riseMax(slice []int, k, end int) {
	for {
		i := 2*k + 1
		if i > end {
			break
		}
		if i < end && slice[i+1] > slice[i] {
			i++
		}
		if slice[k] >= slice[i] {
			break
		}
		slice[i], slice[k] = slice[k], slice[i]
		k = i
	}
}

func riseMin(slice []int, k, end int) {
	for {
		i := 2*k + 1
		if i > end {
			break
		}
		if i < end && slice[i+1] < slice[i] {
			i++
		}
		if slice[k] <= slice[i] {
			break
		}
		slice[i], slice[k] = slice[k], slice[i]
		k = i
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.max) == 0 {
		this.max = append(this.max, num)
		return
	}
	if this.max[0] < num {
		this.min = append(this.min, num)
		end := len(this.min) - 1
		for i := (end%2 - 2 + end) / 2; i >= 0; i = (i%2 - 2 + i) / 2 {
			riseMin(this.min, i, end)
		}
		if len(this.min) > len(this.max) {
			this.max = append(this.max, this.min[0])
			end2 := len(this.max) - 1
			for i := (end2%2 - 2 + end2) / 2; i >= 0; i = (i%2 - 2 + i) / 2 {
				riseMax(this.max, i, end2)
			}
			this.min[0] = this.min[len(this.min)-1]
			riseMin(this.min, 0, len(this.min)-2)
			this.min = this.min[:len(this.min)-1]
		}
	} else {
		this.max = append(this.max, num)
		end := len(this.max) - 1
		for i := (end%2 - 2 + end) / 2; i >= 0; i = (i%2 - 2 + i) / 2 {
			riseMax(this.max, i, end)
		}
		if len(this.max)-len(this.min) > 1 {
			this.min = append(this.min, this.max[0])
			end2 := len(this.min) - 1
			for i := (end2%2 - 2 + end2) / 2; i >= 0; i = (i%2 - 2 + i) / 2 {
				riseMin(this.min, i, end2)
			}
			this.max[0] = this.max[len(this.max)-1]
			riseMax(this.max, 0, len(this.max)-2)
			this.max = this.max[:len(this.max)-1]
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.max) > len(this.min) {
		return float64(this.max[0])
	} else {
		return (float64(this.max[0]) + float64(this.min[0])) / 2
	}
}

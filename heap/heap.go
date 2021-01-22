package heap

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode {
	return &ListNode{Val: val}
}

func (l *ListNode) Add(val int) *ListNode {
	l.Next = New(val)
	l = l.Next
	return l
}

// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
func mergeKLists(lists []*ListNode) *ListNode {
	for i := len(lists) - 1; i >= 0; i-- {
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
		}
	}
	end := len(lists) - 1
	for i := end / 2; i >= 0; i-- {
		mergeKListsRise(lists, i, end)
	}
	res := &ListNode{}
	n := res
	for end >= 0 {
		n.Next = lists[0]
		n = n.Next
		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists[0], lists[end] = lists[end], lists[0]
			end--
		}
		mergeKListsRise(lists, 0, end)
	}
	return res.Next
}

func mergeKListsRise(slice []*ListNode, k, end int) {
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

type KV struct {
	K int
	V int
}

// 前 K 个高频元素
// 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var slice []*KV
	for k, v := range m {
		slice = append(slice, &KV{K: k, V: v})
	}
	end := len(slice) - 1
	for i := end / 2; i >= 0; i-- {
		topKFrequentRise(slice, i, end)
	}
	var res []int
	for k > 0 {
		res = append(res, slice[0].K)
		slice[0], slice[end] = slice[end], slice[0]
		end--
		topKFrequentRise(slice, 0, end)
		k--
	}
	return res
}

func topKFrequentRise(slice []*KV, k, end int) {
	for {
		i := 2*k + 1
		if i > end {
			break
		}
		if i < end && slice[i].V < slice[i+1].V {
			i++
		}
		if slice[k].V >= slice[i].V {
			break
		}
		slice[k], slice[i] = slice[i], slice[k]
		k = i
	}
}

// 数据流的中位数

// 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

// 例如，

// [2,3,4] 的中位数是 3

// [2,3] 的中位数是 (2 + 3) / 2 = 2.5

// 设计一个支持以下两种操作的数据结构：

// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。
// 示例：

// addNum(1)
// addNum(2)
// findMedian() -> 1.5
// addNum(3)
// findMedian() -> 2
// 进阶:

// 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
// 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
type MedianFinder struct {
	Max []int
	Min []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.Min) == 0 {
		this.AddNumToMax(num)
	} else {
		if this.Max[0] >= num {
			this.AddNumToMax(num)
		} else {
			this.AddNumToMin(num)
		}
	}
	diff := len(this.Max) - len(this.Min)
	if diff == 2 {
		this.AddNumToMin(this.Max[0])
		this.Max[0] = this.Max[len(this.Max)-1]
		this.Max = this.Max[:len(this.Max)-1]
		MedianFinderRiseMax(this.Max, 0, len(this.Max)-1)
	} else if diff == -1 {
		this.AddNumToMax(this.Min[0])
		this.Min[0] = this.Min[len(this.Min)-1]
		this.Min = this.Min[:len(this.Min)-1]
		MedianFinderRiseMin(this.Min, 0, len(this.Min)-1)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Max) > len(this.Min) {
		return float64(this.Max[0])
	} else {
		return float64(this.Max[0]+this.Min[0]) / 2
	}
}

func (this *MedianFinder) AddNumToMin(num int) {
	this.Min = append(this.Min, num)
	length := len(this.Min)
	end := length - 1
	i := end
	for {
		if i%2 == 1 {
			i = (i - 1) / 2
		} else {
			i = (i - 2) / 2
		}
		if i < 0 {
			break
		}
		MedianFinderRiseMin(this.Min, i, end)
		if i == 0 {
			break
		}
	}
}

func (this *MedianFinder) AddNumToMax(num int) {
	this.Max = append(this.Max, num)
	length := len(this.Max)
	end := length - 1
	i := end
	for {
		if i%2 == 1 {
			i = (i - 1) / 2
		} else {
			i = (i - 2) / 2
		}
		if i < 0 {
			break
		}
		MedianFinderRiseMax(this.Max, i, end)
		if i == 0 {
			break
		}
	}
}

func MedianFinderRiseMax(slice []int, k, end int) {
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
		slice[k], slice[i] = slice[i], slice[k]
		k = i
	}
}

func MedianFinderRiseMin(slice []int, k, end int) {
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
		slice[k], slice[i] = slice[i], slice[k]
		k = i
	}
}

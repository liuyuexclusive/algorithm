package mergetwosortedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Add(val int) *ListNode {
	newNode := &ListNode{Val: val}
	for {
		if head.Next == nil {
			head.Next = newNode
			break
		}
		head = head.Next
	}
	return newNode
}

func (head *ListNode) Show() string {
	var res string
	for head != nil {
		res = fmt.Sprintf("%s%d", res, head.Val)
		head = head.Next
	}
	return res
}

func getMin(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode) {
	if l1 == nil && l2 == nil {
		return l1, l2
	} else if l1 == nil && l2 != nil {
		return l2, l1
	} else if l1 != nil && l2 == nil {
		return l1, l2
	} else if l1.Val > l2.Val {
		return l2, l1
	}
	return l1, l2
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res, max := getMin(l1, l2)
	min := res
	for min != nil {
		min.Next, max = getMin(min.Next, max)
		min = min.Next
	}

	return res
}

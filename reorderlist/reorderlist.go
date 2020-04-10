package reorderlist

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

func reorderList(head *ListNode) {
	m := make(map[int]*ListNode)
	index := 0
	temp := head
	for temp != nil {
		m[index] = temp
		index++
		temp = temp.Next
	}
	left, right := 0, index-1

	for left < right {
		m[left].Next = m[right]
		if left+1 == right {
			m[right].Next = nil
		} else if left+2 == right {
			m[right].Next = m[left+1]
			m[left+1].Next = nil
		} else {
			m[right].Next = m[left+1]
		}
		left++
		right--
	}
}

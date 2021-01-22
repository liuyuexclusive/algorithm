package linkedlist

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) show() []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func (head *ListNode) add(val int) *ListNode {
	if head == nil {
		panic("null list")
	}
	for head.Next != nil {
		head = head.Next
	}
	last := &ListNode{Val: val}
	head.Next = last
	return last
}

// 反转链表

// 反转一个单链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

// 环形链表

// 给定一个链表，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
// 如果链表中存在环，则返回 true 。 否则，返回 false 。
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 合并两个有序链表

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next

		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return res.Next
}

// 合并K个升序链表

// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
func mergeKLists(lists []*ListNode) *ListNode {
	temp := &ListNode{}
	res := temp
	sort.Slice(lists, func(i, j int) bool {
		if lists[i] == nil {
			return true
		}
		if lists[j] == nil {
			return false
		}
		return lists[i].Val < lists[j].Val
	})
	for len(lists) > 0 {
		if lists[0] == nil {
			lists = lists[1:]
			continue
		}
		temp.Next = lists[0]
		temp = temp.Next
		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists = lists[1:]
		} else {
			for i := 0; i < len(lists)-1; i++ {
				if lists[i].Val < lists[i+1].Val {
					break
				} else {
					lists[i], lists[i+1] = lists[i+1], lists[i]
				}
			}
		}
	}
	return res.Next
}

// 删除链表的倒数第N个节点
// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var res []*ListNode
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	length := len(res)
	if n == length {
		return res[0].Next
	}
	res[length-n-1].Next = res[length-n].Next
	return res[0]
}

// 重排链表

// 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
// 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
func reorderList(head *ListNode) {
	var slice []*ListNode
	for head != nil {
		slice = append(slice, head)
		head = head.Next
	}
	length := len(slice)
	temp := &ListNode{}
	for i := 0; i < length/2; i++ {
		temp.Next = slice[i]
		temp = temp.Next
		temp.Next = slice[length-i-1]
		temp = temp.Next
	}
	if length%2 == 1 {
		temp.Next = slice[length/2]
		temp = temp.Next
	}
	temp.Next = nil
}

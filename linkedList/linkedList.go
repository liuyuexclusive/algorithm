package linkedList

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

//翻转链表
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

//判断链表有环
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var p *ListNode = &ListNode{}
	res := p
	order := make([]*ListNode, 0)
	if l1 != nil {
		order = append(order, l1)
	}
	if l2 != nil {
		order = append(order, l2)
	}
	for len(order) > 0 {
		sort.Slice(order, func(i, j int) bool {
			return order[i].Val < order[j].Val
		})
		p.Next = order[0]
		order[0] = order[0].Next
		if order[0] == nil {
			order = order[1:]
		}
		p = p.Next
	}
	return res.Next
}

//合并k个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	var p *ListNode = &ListNode{}
	res := p
	orderList := make([]*ListNode, 0)
	for _, v := range lists {
		if v != nil {
			orderList = append(orderList, v)
		}
	}
	sort.Slice(orderList, func(i, j int) bool {
		return orderList[i].Val < orderList[j].Val
	})
	for len(orderList) > 0 {
		p.Next = orderList[0]
		orderList[0] = orderList[0].Next
		if orderList[0] == nil {
			orderList = orderList[1:]
		} else {
			for i := 1; i < len(orderList); i++ {
				if orderList[i].Val < orderList[i-1].Val {
					orderList[i], orderList[i-1] = orderList[i-1], orderList[i]
				} else {
					break
				}
			}
		}
		p = p.Next
	}
	return res.Next
}

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slice := make([]*ListNode, 0)
	for head != nil {
		slice = append(slice, head)
		head = head.Next
	}
	if n == len(slice) {
		return slice[0].Next
	}
	slice[len(slice)-n-1].Next = slice[len(slice)-n].Next
	return slice[0]
}

//重排链表
// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

// You may not modify the values in the list's nodes, only nodes itself may be changed.
func reorderList(head *ListNode) {
	slice := make([]*ListNode, 0)
	for head != nil {
		slice = append(slice, head)
		head = head.Next
	}

	for i := 0; i < len(slice)/2; i++ {
		if len(slice)-i-1 > i {
			if len(slice)-i-1-(i) == 1 {
				slice[len(slice)-i-1].Next = nil
			}
			slice[i].Next = slice[len(slice)-i-1]
		}
		if len(slice)-i-1 > i+1 {
			if len(slice)-i-1-(i+1) == 1 {
				slice[i+1].Next = nil
			}
			slice[len(slice)-i-1].Next = slice[i+1]
		}
	}
}

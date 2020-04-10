package removefromend

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

// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

// 示例：

// 给定一个链表: 1->2->3->4->5, 和 n = 2.

// 当删除了倒数第二个节点后，链表变为 1->2->3->5.

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	res := head
	m, index := make(map[int]*ListNode), 0
	for head != nil {
		m[index] = head
		head = head.Next
		index++
	}

	if index == n {
		return res.Next
	}

	if n == 1 {
		m[index-2].Next = nil
		return res
	}

	m[index-n-1].Next = m[index-n+1]
	return res
}

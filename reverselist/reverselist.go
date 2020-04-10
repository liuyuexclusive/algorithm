package reverselist

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  反转一个单链表。

// 示例:

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-linked-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

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

func (head *ListNode) show() string {
	var res string
	for head != nil {
		res = fmt.Sprintf("%s%d\t", res, head.Val)
		head = head.Next
	}
	return res
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}

	// fmt.Println(preOrder(root))
	// fmt.Println(inOrder(root))
	fmt.Println(postOrder(root))
}

func preOrder(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			root = root.Left
		}
		last := stack[len(stack)-1]
		root = last.Right
		stack = stack[:len(stack)-1]
	}
	return res
}
func inOrder(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := stack[len(stack)-1]
		res = append(res, last.Val)
		root = last.Right
		stack = stack[:len(stack)-1]
	}
	return res
}

func postOrder(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := stack[len(stack)-1]
		if (last.Left == nil && last.Right == nil) || (last.Left == pre && last.Right == nil) || last.Right == pre {
			pre = last
			res = append(res, last.Val)
			stack = stack[:len(stack)-1]
		} else {
			root = last.Right
		}
	}
	return res
}

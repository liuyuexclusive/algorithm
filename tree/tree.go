package tree

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉树，找出其最大深度。

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

// 说明: 叶子节点是指没有子节点的节点。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

// 给定两个二叉树，编写一个函数来检验它们是否相同。
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

//翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 给定一个非空二叉树，返回其最大路径和。
// 本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	traverse(root, &res)
	return res
}

func traverse(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := max(traverse(root.Left, res), 0)
	r := max(traverse(root.Right, res), 0)
	if val := root.Val + l + r; val > *res {
		*res = val
	}
	return root.Val + max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := []*TreeNode{root}
	layer := -1
	m := map[int]int{-1: 0, 0: 1}
	for len(queue) > 0 {
		if m[layer] == 0 {
			layer++
		}
		head := queue[0]
		if head != nil {
			if layer > len(res)-1 {
				res = append(res, make([]int, 0))
			}
			res[layer] = append(res[layer], head.Val)
			queue = append(queue, []*TreeNode{head.Left, head.Right}...)
			m[layer+1] += 2
		}
		queue = queue[1:]
		m[layer]--
	}
	return res
}

//二叉树的序列化
type Codec struct {
}

func ConstructorCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		head := queue[0]
		if head != nil {
			if res == "" {
				res += strconv.Itoa(head.Val)
			} else {
				res += "," + strconv.Itoa(head.Val)
			}
			queue = append(queue, []*TreeNode{head.Left, head.Right}...)
		} else {
			res += "," + "N"
		}
		queue = queue[1:]
	}
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	if len(arr) == 0 {
		return nil
	}
	var queue []*TreeNode
	var res *TreeNode
	for i, v := range arr {
		var node *TreeNode
		if v != "N" {
			val, _ := strconv.Atoi(v)
			node = &TreeNode{Val: val}
			queue = append(queue, node)
			if i == 0 {
				res = node
			}
		} else {

		}
		if i != 0 && i%2 == 0 {
			queue[0].Right = node
			queue = queue[1:]
		} else {
			queue[0].Left = node
		}
	}
	return res
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return issame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func issame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && issame(s.Left, t.Left) && issame(s.Right, t.Right)
}

//用中序和前序遍历还原二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	var i int
	for x := 0; x < len(inorder); x++ {
		if inorder[x] == root.Val {
			i = x
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}

//判断是否是二叉搜索树
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	p := root
	var pre *TreeNode
	for len(stack) > 0 || p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) > 0 {
			tail := stack[len(stack)-1]
			if pre != nil && pre.Val >= tail.Val {
				return false
			}
			pre = tail
			p = tail.Right
			stack = stack[:len(stack)-1]
		}
	}
	return true
}

//第k个最小元素
func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	p := root
	var count int = 1
	for len(stack) > 0 || p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		tail := stack[len(stack)-1]
		if count == k {
			return tail.Val
		}
		p = tail.Right
		stack = stack[:len(stack)-1]
		count++
	}
	return -1
}

//最低公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > q.Val {
		p, q = q, p
	}
	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	} else if root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return lowestCommonAncestor(root.Left, p, q)
	}

}

type WordDictionary struct {
	Root *Node
}

type Node struct {
	Val   byte
	Nodes map[byte]*Node
	End   bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{Root: &Node{Nodes: make(map[byte]*Node)}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this.Root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.Nodes[char]; !ok {
			node.Nodes[char] = &Node{Val: char, Nodes: make(map[byte]*Node)}
			if i == len(word)-1 {
				node.Nodes[char].End = true
			}
		}
		node = node.Nodes[char]
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	node := this.Root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if char == '.' {
			if i == len(word)-1 {
				for _, v := range node.Nodes {
					if v.End {
						return true
					}
				}
				return false
			}
			for _, v := range node.Nodes {
				if (&WordDictionary{Root: v}).Search(word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			v, ok := node.Nodes[char]
			res := ok && (&WordDictionary{Root: v}).Search(word[i+1:])
			if i == len(word)-1 {
				return res && v.End
			}
			return res
		}
	}
	return true
}

// func findWords(board [][]byte, words []string) []string {
// 	var res []string
// 	if len(board) == 0 {
// 		return res
// 	}
// loop:
// 	for _, v := range words {
// 		m := make([][]bool, len(board))
// 		for i := 0; i < len(board); i++ {
// 			m[i] = make([]bool, len(board[i]))
// 		}
// 		for i := 0; i < len(board); i++ {
// 			for j := 0; j < len(board[0]); j++ {
// 				if findWordsValidate(board, m, v, 0, i, j) {
// 					res = append(res, v)
// 					continue loop
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// func findWordsValidate(board [][]byte, m [][]bool, word string, index, i, j int) bool {
// 	if index == len(word) {
// 		return true
// 	}
// 	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || m[i][j] || word[index] != board[i][j] {
// 		return false
// 	}
// 	m[i][j] = true
// 	index++
// 	res := findWordsValidate(board, m, word, index, i+1, j) ||
// 		findWordsValidate(board, m, word, index, i-1, j) ||
// 		findWordsValidate(board, m, word, index, i, j-1) ||
// 		findWordsValidate(board, m, word, index, i, j+1)
// 	if !res {
// 		m[i][j] = false
// 	}
// 	return res
// }

type findWordsNode struct {
	Word  string
	Nodes map[byte]*findWordsNode
}

func findWords(board [][]byte, words []string) []string {
	root := &findWordsNode{Nodes: make(map[byte]*findWordsNode)}
	r := root
	var res []string
	for _, word := range words {
		r = root
		for _, v := range word {
			c := byte(v)
			if _, ok := r.Nodes[c]; !ok {
				r.Nodes[c] = &findWordsNode{Nodes: make(map[byte]*findWordsNode)}
			}
			r = r.Nodes[c]
		}
		r.Word = word
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, root, i, j, &res)
		}
	}
	return res
}

func dfs(board [][]byte, r *findWordsNode, i, j int, res *[]string) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}
	v := board[i][j]
	if v == 'N' {
		return
	}
	c, ok := r.Nodes[v]
	if !ok {
		return
	}

	board[i][j] = 'N'

	if c.Word != "" {
		*res = append(*res, c.Word)
		c.Word = ""
	}
	r = c

	dfs(board, r, i-1, j, res)
	dfs(board, r, i+1, j, res)
	dfs(board, r, i, j-1, res)
	dfs(board, r, i, j+1, res)
	board[i][j] = v
}

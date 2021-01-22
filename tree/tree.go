package tree

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 二叉树的最大深度

// 给定一个二叉树，找出其最大深度。
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// 说明: 叶子节点是指没有子节点的节点。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	max := left
	if left < right {
		max = right
	}
	return max + 1
}

// 相同的树

// 给定两个二叉树，编写一个函数来检验它们是否相同。
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
func isSameTree(p *TreeNode, q *TreeNode) bool {
	f := func(x, y *TreeNode) bool {
		if x == nil && y == nil {
			return true
		} else if x == nil || y == nil {
			return false
		}
		return x.Val == y.Val
	}
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	return f(p, q) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 翻转二叉树

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 二叉树中的最大路径和

// 给你一个二叉树的根节点 root ，返回其最大路径和。
// 本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径 至少包含一个 节点，且不一定经过根节点。
func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	nodeSum(root, &res)
	return res
}

func nodeSum(r *TreeNode, res *int) int {
	if r == nil {
		return 0
	}
	left := max(nodeSum(r.Left, res), 0)
	right := max(nodeSum(r.Right, res), 0)
	if nv := r.Val + left + right; nv > *res {
		*res = nv
	}
	return r.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 二叉树的层序遍历

// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := []*TreeNode{root}
	m := make(map[int]int)
	m[0] = 1
	currentLayer, currentLayerCount := 0, 0
	for len(queue) > 0 {
		head := queue[0]
		currentLayerCount++
		if head != nil {
			if currentLayer >= len(res) {
				res = append(res, make([]int, 0))
			}
			res[currentLayer] = append(res[currentLayer], head.Val)
			queue = append(queue, []*TreeNode{head.Left, head.Right}...)
			m[currentLayer+1] += 2
		}
		if currentLayerCount == m[currentLayer] {
			currentLayer++
			currentLayerCount = 0
		}
		queue = queue[1:]
	}
	return res
}

type Codec struct {
}

func ConstructorCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		head := queue[0]
		if head != nil {
			res = append(res, strconv.Itoa(head.Val))
			queue = append(queue, []*TreeNode{head.Left, head.Right}...)
		} else {
			res = append(res, "null")
		}
		queue = queue[1:]
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	slice := strings.Split(data, ",")
	if len(slice) == 0 || slice[0] == "null" {
		return nil
	}
	rootVal, _ := strconv.Atoi(slice[0])
	res := &TreeNode{Val: rootVal}
	queue := []*TreeNode{res}
	for i := 1; i < len(slice); i++ {
		v := slice[i]
		var node *TreeNode
		if v != "null" {
			num, _ := strconv.Atoi(v)
			node = &TreeNode{Val: num}
		}
		queue = append(queue, node)
		head := queue[0]
		for head == nil {
			queue = queue[1:]
			head = queue[0]
		}
		if i%2 == 1 {
			head.Left = node
		} else {
			head.Right = node
			queue = queue[1:]
		}
	}
	return res
}

// 另一个树的子树

// 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	if s == nil {
		return isSameNdoe(s, t)
	}
	return isSameNdoe(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameNdoe(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isSameNdoe(s.Left, t.Left) && isSameNdoe(s.Right, t.Right)
}

// 从前序与中序遍历序列构造二叉树
// 根据一棵树的前序遍历与中序遍历构造二叉树。
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	midVal := preorder[0]
	var mid int
	for k, v := range inorder {
		if midVal == v {
			mid = k
		}
	}
	node := &TreeNode{Val: midVal}
	node.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	node.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return node
}

// 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	return isValidBSTInOrder(root, &pre)
}

func isValidBSTInOrder(root *TreeNode, pre *int) bool {
	if root == nil {
		return true
	}
	if !isValidBSTInOrder(root.Left, pre) {
		return false
	}
	if *pre >= root.Val {
		return false
	}
	*pre = root.Val
	if !isValidBSTInOrder(root.Right, pre) {
		return false
	}
	return true
}

// 二叉搜索树中第K小的元素

// 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

// 说明：
// 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		k--
		last := stack[len(stack)-1]
		if k == 0 {
			return last.Val
		}
		root = last.Right
		stack = stack[:len(stack)-1]
	}
	return 0
}

// 二叉搜索树的最近公共祖先

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
// 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
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

type TrieNode struct {
	End  bool
	Data map[byte]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Data: make(map[byte]*TrieNode)}
}

/** Initialize your data structure here. */
func ConstructorTrie() Trie {
	return Trie{Root: NewTrieNode()}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.Root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.Data[char]; !ok {
			node.Data[char] = NewTrieNode()
		}
		node = node.Data[char]
	}
	node.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.Root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.Data[char]; !ok {
			return false
		}
		node = node.Data[char]
	}
	if !node.End {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.Root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if _, ok := node.Data[char]; !ok {
			return false
		}
		node = node.Data[char]
	}
	return true
}

type WordDictionary struct {
	Root *WordDictionaryNode
}

type WordDictionaryNode struct {
	End  bool
	Data map[byte]*WordDictionaryNode
}

func NewWordDictionaryNode() *WordDictionaryNode {
	return &WordDictionaryNode{Data: make(map[byte]*WordDictionaryNode)}
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{Root: NewWordDictionaryNode()}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.Root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.Data[char]; !ok {
			node.Data[char] = NewWordDictionaryNode()
		}
		node = node.Data[char]
	}
	node.End = true
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.Root.End
	}
	node := this.Root
	head := word[0]
	if head == '.' {
		for _, n := range node.Data {
			if (&WordDictionary{Root: n}).Search(word[1:]) {
				return true
			}
		}
		return false
	}
	if n, ok := node.Data[head]; !ok {
		return false
	} else {
		return (&WordDictionary{Root: n}).Search(word[1:])
	}
}

// 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
func findWords(board [][]byte, words []string) []string {
	getm := func() [][]bool {
		m := make([][]bool, len(board))
		for i := 0; i < len(m); i++ {
			m[i] = make([]bool, len(board[i]))
		}
		return m
	}
	var res []string
loop:
	for _, word := range words {
		m := getm()
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if findWords2(board, m, i, j, 0, word) {
					res = append(res, word)
					continue loop
				}
			}
		}
	}
	return res
}

func findWords2(board [][]byte, m [][]bool, i, j, index int, word string) bool {
	if index == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || m[i][j] || board[i][j] != word[index] {
		return false
	}
	m[i][j] = true
	index++
	res := findWords2(board, m, i+1, j, index, word) ||
		findWords2(board, m, i-1, j, index, word) ||
		findWords2(board, m, i, j+1, index, word) ||
		findWords2(board, m, i, j-1, index, word)
	if !res {
		m[i][j] = false
	}
	return res
}

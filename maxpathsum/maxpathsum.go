package maxpathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	return maxPathSum(root.Left)+maxPathSum(root.)
}

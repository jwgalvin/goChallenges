package max_path_sum

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxPathSum returns the maximum path sum in a binary tree.
// A path can start and end at any node; each node appears at most once.
// Input: [-10,9,20,null,null,15,7]  →  42  (15→20→7)
func MaxPathSum(root *TreeNode) int {
	// TODO: implement (DFS tracking max gain per subtree)
	return 0
}

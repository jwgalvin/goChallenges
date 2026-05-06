package build_tree

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree constructs a binary tree from preorder and inorder traversals.
// Input: preorder=[3,9,20,15,7], inorder=[9,3,15,20,7]  →  tree rooted at 3
func BuildTree(preorder []int, inorder []int) *TreeNode {
	// TODO: implement (divide-and-conquer using inorder index map)
	return nil
}

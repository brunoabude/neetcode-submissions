/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	k := 0

	for inorder[k] != root.Val {
		k++
	}

	root.Left = buildTree(preorder[1:k+1], inorder[:k])
	root.Right = buildTree(preorder[k+1:], inorder[k+1:])

	return root
}
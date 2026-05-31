/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func checkStructure(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}
	
	return checkStructure(left.Left, right.Left) && checkStructure(left.Right, right.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var dfs func(left, right *TreeNode) bool 

	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if checkStructure(left, right) {
			return true
		}

		return dfs(left.Left, right) || dfs(left.Right, right)
	}

	return dfs(root, subRoot)
}

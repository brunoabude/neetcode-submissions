/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	current := root

	for current != nil {
		if p.Val < current.Val && q.Val < current.Val {
			current = current.Left
			continue
		}

		if p.Val > current.Val && q.Val > current.Val {
			current = current.Right
			continue
		}

		break
	}

	return current
}

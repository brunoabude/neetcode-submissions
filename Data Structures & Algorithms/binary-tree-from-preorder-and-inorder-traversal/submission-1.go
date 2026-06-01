/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func buildTree(preorder []int, inorder []int) *TreeNode {
	indexes := map[int]int{}

	for i, v := range inorder{
		indexes[v] = i
	}

	var dfs func(int, int) *TreeNode

	currPreorderIdx := 0

	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		root := &TreeNode{Val: preorder[currPreorderIdx]}
		currPreorderIdx++

		inorderIndex := indexes[root.Val]

		root.Left = dfs(left, inorderIndex - 1)
		root.Right = dfs(inorderIndex +1, right)

		return root
	}

	return dfs(0, len(inorder)- 1)
}
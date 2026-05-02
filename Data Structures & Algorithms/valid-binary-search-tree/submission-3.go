/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorder(arr *[]int, node *TreeNode)  {
	if node == nil {
		return
	}
	inorder(arr, node.Left)
	*(arr) = append(*(arr), node.Val)
	inorder(arr, node.Right)
}

func isValidBST(root *TreeNode) bool {
	arr := []int{}
	inorder(&arr, root)

	if len(arr) == 1 {
		return true
	}

	for i := range len(arr)-1 {
		if arr[i+1] <= arr[i] {
			return false
		}
	}

	return true
}

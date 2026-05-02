func inorder(arr *[]int, node *TreeNode) {
	if node == nil {
		return
	}

	inorder(arr, node.Left)
	*arr = append(*arr, node.Val)
	inorder(arr, node.Right)
}

func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int, 0, 1000)
	inorder(&arr, root)
	
	return arr[k-1]
}
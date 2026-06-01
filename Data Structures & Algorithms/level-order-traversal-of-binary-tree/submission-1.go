/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}

	if root != nil {
		queue = append(queue, root)
	}

	res := [][]int{}

	for len(queue) > 0 {
		levelLength := len(queue)
		level := make([]int, 0, levelLength)

		for i := range levelLength {
			node := queue[i]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)

		queue = queue[levelLength:]
	}


	return res
}

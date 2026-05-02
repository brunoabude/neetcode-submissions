/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func checkBST(root *TreeNode) (int, int, bool) {
	minimum := root.Val
	maximum := root.Val
	
	//basic check
	if root.Left != nil && root.Left.Val >= root.Val {
		return 0, 0, false
	}

	if root.Right != nil && root.Right.Val <= root.Val {
		return 0, 0, false
	}

    if root.Left != nil {
		foundMin, foundMax, valid := checkBST(root.Left)

		if !valid {
			return 0, 0, false
		}

		if foundMax >= root.Val {
			return 0, 0, false
		}

		if foundMin < minimum {
			minimum = foundMin
		}

		if foundMax > maximum {
			maximum = foundMax
		}
	}

    if root.Right != nil {
		foundMin, foundMax, valid := checkBST(root.Right)

		if !valid {
			return 0, 0, false
		}

		if foundMin <= root.Val {
			return 0, 0, false
		}

		if foundMin < minimum {
			minimum = foundMin
		}

		if foundMax > maximum {
			maximum = foundMax
		}
	}

	return minimum, maximum, true
}

func isValidBST(root *TreeNode) bool {
	_, _, valid := checkBST(root)

	return valid
}

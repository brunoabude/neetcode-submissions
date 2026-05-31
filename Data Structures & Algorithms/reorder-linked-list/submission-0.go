/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorder(node *ListNode, idx, n int) *ListNode {
	// get idx (current idx)
	// get n - (idx-depth) - 1
	if node == nil {
		return nil
	}

	if 2*idx >= n {
		return nil
	}

	tmp := node.Next // idx+1

	nth_node := node

	for range n - 2*idx - 1 {
		nth_node = nth_node.Next
	}

	node.Next = nth_node
	nth_node.Next = reorder(tmp, idx+1, n)

	return node
}

func reorderList(head *ListNode) {
	n := 0

	curr := head

	for curr != nil {
		n++
		curr = curr.Next
	}

	reorder(head, 0, n)
}
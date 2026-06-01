/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

    left, right := dummy, dummy
	left_idx, right_idx := -1, -1

	for right != nil {
		right = right.Next
		right_idx++

		for right_idx - left_idx > n + 1{
			left_idx++
			left = left.Next
		}
	}

	if left.Next != nil {
		left.Next = left.Next.Next
	}

	return dummy.Next
}

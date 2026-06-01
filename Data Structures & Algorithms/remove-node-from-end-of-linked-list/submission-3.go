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

	for range n+1 {
		right = right.Next
	}

	for right != nil {
		right = right.Next
		left = left.Next
	}

	if left.Next != nil {
		left.Next = left.Next.Next
	}

	return dummy.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */



func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    left_ptr := list1
	right_ptr := list2

	// dummy head to help
	var head *ListNode = &ListNode{Val: -1, Next: nil}
	curr := head

	for left_ptr != nil && right_ptr != nil {
		if left_ptr.Val < right_ptr.Val {
			curr.Next = left_ptr
			left_ptr = left_ptr.Next
		} else {
			curr.Next = right_ptr
			right_ptr = right_ptr.Next
		}

		curr = curr.Next
	}


	if left_ptr != nil {
		curr.Next = left_ptr
	} else {
		curr.Next = right_ptr
	}

	return head.Next
}

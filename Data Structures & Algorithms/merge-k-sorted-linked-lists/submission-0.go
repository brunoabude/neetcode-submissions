/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    // "brute" force

	pointers := make([]*ListNode, len(lists))

	for i, p := range lists {
		pointers[i] = p
	}

	dummy := &ListNode{}
	curr := dummy

	for {
		var smallest *ListNode
		pIdx := -1

		for i, v := range pointers {
			if (v != nil && smallest == nil) || (v != nil && v.Val <= smallest.Val) {
				smallest = v
				pIdx = i
			}
		}

		if pIdx == -1 {
			break
		}

		pointers[pIdx] = smallest.Next

		curr.Next = smallest
		curr = curr.Next
	}

	return dummy.Next
}

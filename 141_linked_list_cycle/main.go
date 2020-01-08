package _41_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	curNode, skippedNode := head, head.Next
	for skippedNode != nil && skippedNode.Next != nil {
		if curNode == skippedNode {
			return true
		}
		curNode = curNode.Next
		skippedNode = skippedNode.Next.Next
	}
	return false
}


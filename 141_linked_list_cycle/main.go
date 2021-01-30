package _41_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle/
// O(n) time and memory solution
	func hasCycle(head *ListNode) bool {
	nodeMap := map[*ListNode]bool{}
	curNode := head
	for curNode != nil {
		if nodeMap[curNode] {
			return true
		}
		nodeMap[curNode] = true
		curNode = curNode.Next
	}
	return false
}

// O(1) memory solution
func hasCycle2(head *ListNode) bool {
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


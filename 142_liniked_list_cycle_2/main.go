package _42_liniked_list_cycle_2

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle-ii/
// O(n) time and memory solution
func detectCycle(head *ListNode) *ListNode {
	nodeMap := map[*ListNode]bool{}
	curNode := head
	for curNode != nil {
		if nodeMap[curNode] {
			return curNode
		}
		nodeMap[curNode] = true
		curNode = curNode.Next
	}
	return nil
}


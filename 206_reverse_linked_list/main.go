package _06_reverse_linked_list

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    nodes := make([]*ListNode, 0)
    curNode := head
    for curNode.Next != nil {
        nodes = append(nodes, curNode)
        curNode = curNode.Next
    }

    curNode = nodes[len(nodes) - 1]
    head = curNode
    nodes = nodes[:len(nodes) - 1]
    for len(nodes) > 0 {
        curNode.Next = nodes[len(nodes) - 1]
        curNode = nodes[len(nodes) - 1]
        nodes = nodes[:len(nodes) - 1]
    }
    curNode.Next = nil
    return head
}
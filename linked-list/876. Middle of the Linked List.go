package linked_list

//Given the head of a singly linked list, return the middle node of the linked list.
//If there are two middle nodes, return the second middle node.

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil {
			break
		}
	}

	return slow
}

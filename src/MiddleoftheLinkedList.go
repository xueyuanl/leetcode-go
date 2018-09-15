package src

func main() {

}

func middleNode(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}

	var fast *ListNode = head
	var slow *ListNode = head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast.Next == nil {
		return slow
	} else {
		return slow.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

package algo

import . "basic"

/*The classical linked_list reverse algorithm*/
func reverseList(head *ListNode) *ListNode {

	var newHead *ListNode = nil

	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}

	return newHead
}

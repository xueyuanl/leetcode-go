package list

import . "basic"

/*The classical linked_list reverse algorithm*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	old := head
	new := head
	var tempOld *ListNode
	for new.Next != nil {
		new = new.Next
		old.Next = tempOld
		tempOld = old
		old = new

	}
	new.Next = tempOld
	return new
}

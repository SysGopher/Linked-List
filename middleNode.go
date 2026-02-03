package main


func (ll *LinkedList) GetMiddle () *Node{
	if ll.head == nil {
        return nil
    }
	slow := ll.head
	fast := ll.head
	
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow

}
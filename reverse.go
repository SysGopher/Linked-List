package main


func (ll *LinkedList) Reverse(){
	var prev *Node = nil
	current := ll.head
	var next *Node = nil

	for current != nil {
		next = current.next
		current.next = prev

		prev = current 
		current = next
	}
	ll.head = prev
}
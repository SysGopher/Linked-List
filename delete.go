package main



func (ll *LinkedList) delete(key int){

	if ll.head == nil{
      return
	}

	current := ll.head

	if ll.head.data == key {
		ll.head = ll.head.next
		return 
	}
	for current.next != nil && current.next.data != key {
     current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}
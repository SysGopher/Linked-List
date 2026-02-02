package main



func (ll *LinkedList) insertStartAndEnd(value int){
	newNode := &Node{data: value}

	if ll.head == nil{
		ll.head = newNode
        return
	}
	current := ll.head
	for current.next != nil{
		current = current.next
	}
	current.next = newNode
}


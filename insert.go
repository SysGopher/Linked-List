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


func (ll *LinkedList) insertSpecificPosition(afterValue, newValue int){
	current := ll.head
	for current != nil && current.data != afterValue{
		current = current.next

	}

	if current != nil{
		newNode := &Node{data :newValue}
		newNode.next = current.next
		current.next = newNode
	}
}

package main
import (
     "fmt"
	
)


func (ll *LinkedList) Traverse(){


	current := ll.head
	for current != nil{
		fmt.Printf("%d -> ", current.data)
		current = current.next

		
	}
	fmt.Println("nil")  
    

}


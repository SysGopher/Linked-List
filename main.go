package main




type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}


func main(){
	myList := LinkedList{}
    numbers := []int{10, 20, 30, 40,50}

	for _, num := range numbers {
		myList.insertStartAndEnd(num)
	}
	//myList.insertSpecificPosition()
	//myList.delete(10)
	  //myList.Reverse()
    myList.Traverse()

	//mid := myList.GetMiddle()
	
	
	

	
}
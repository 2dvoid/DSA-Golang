package main

func (l *LinkedList) Append(val int){
	// Create the new Node
	newNode := &Node{
		data: val,
	}
	
	// Length Increment
	l.length++

	// Case 1: List is Empty
	if l.head == nil {
		// Head and Tail will now point to the new Node
		l.head = newNode
		l.tail = newNode
		return
	}
	// Case 2: List not Empty
	// Attach the new Node to the Tail
	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
}

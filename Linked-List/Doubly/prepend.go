package main

func (l *LinkedList) Prepend(val int){
	// Create the new Node
	newNode := &Node{
		data: val,
	}

	// Length Increment
	l.length++

	// Case 1: List Empty
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	// Case 2: List not Empty
	// New Node will be the Head
	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode

}

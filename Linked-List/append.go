package main

func (l *LinkedList) Append(val int) {

	// Create a new Node
	newNode := &Node{
		data: val,
	}

	// Length increment
	l.length++

	// If the list is empty
	if l.head == nil {
		// New Node is the Head and Tail
		l.head = newNode
		l.tail = newNode
		return
	}

	// If not empty
	// Attach the new node the the tail node
	l.tail.link = newNode
	// Update the Tail variable
	l.tail = newNode
}

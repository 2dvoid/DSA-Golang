package main

func (l *LinkedList) Append(val int) {
	// Create a new Node
	newNode := &Node{
		data: val,
	}
	// Check if the provided list is empty or not
	if l.head == nil {
		// If Empty then attach the new node the the head directly
		l.head = newNode
		l.tail = newNode
	} else {
		// If list not empty
		// Attach the new node the the tail node (last node)
		l.tail.link = newNode
		// Update the global tail
		l.tail = newNode
	}

}

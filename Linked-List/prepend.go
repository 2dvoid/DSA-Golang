package main

func (l *LinkedList) Prepend(val int) {
	// Create a new node
	newNode := &Node{
		data: val,
	}
	// If the list is empty (Head points to null)
	if l.head == nil {
		// Point the Head and the Tail to this newly created node
		l.head = newNode
		l.tail = newNode
		return
	}
	// If list is not empty
	// Point the New Node to the Head node
	newNode.link = l.head
	// New Node becomes the Head!
	l.head = newNode
}

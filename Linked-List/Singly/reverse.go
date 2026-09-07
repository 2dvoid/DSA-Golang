package main

func (l *LinkedList) Reverse() {
	var prev *Node
	curr := l.head

	for curr != nil {
		// Holds the link to the rest of the Nodes
		next := curr.link
		// Current will point to it's previous Node
		curr.link = prev

		prev = curr
		curr = next
	}
	// Tail will become the Head and the Head will become the Tail
	l.tail = l.head
	// While curr is nil, prev holds the last Node's address, and it will become the Head
	l.head = prev
}

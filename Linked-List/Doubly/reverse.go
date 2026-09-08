package main

func (l *LinkedList) Reverse() {
	// If the list is empty or has only 1 node, it is already reversed.
	if l.length <= 1 {
		return
	}
	curr := l.head
	for curr != nil {
		tmp := curr.next
		curr.next = curr.prev
		curr.prev = tmp
		curr = curr.prev
	}
	tmp := l.head
	l.head = l.tail
	l.tail = tmp
}

package main

import "fmt"

func (l *LinkedList) Delete(pos int) {

	// Case 1: List is Empty
	if l.head == nil {
		fmt.Println("List is Empty. Nothing to delete.")
		return
	}

	// Case 2: Position out of bound
	if pos < 1 || pos > l.length {
		fmt.Println("Invalid position.")
		return
	}

	// Case 3: Only one node in the list
	if l.length == 1 {
		l.head = nil
		l.tail = nil
		l.length--
		return
	}

	// Case 4: Delete the Head Node
	if pos == 1 {
		l.head = l.head.next
		l.head.prev = nil
		// Length Decrement
		l.length--
		return
	}

	// Case 5: Delete the tail Node
	if pos == l.length {
		l.tail = l.tail.prev
		l.tail.next = nil
		// Length Decrement
		l.length--
		return
	}
	// Case 6: Main Logic
	curr := l.head
	for i := 1; i < pos; i++ {
		curr = curr.next
	}
	curr.prev.next = curr.next
	curr.next.prev = curr.prev

	// Length Decrement
	l.length--
}

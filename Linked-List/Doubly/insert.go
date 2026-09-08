package main

import "fmt"

func (l *LinkedList) Insert(pos int, val int) {

	// Filter out bad data immediately and exit
	if pos < 1 || pos > l.length+1 {
		fmt.Println("Invalid Position")
		return
	}

	// Position 1 means Prepend
	if pos == 1 {
		l.Prepend(val)
		return
	}

	// Position length+1 means Append
	if pos == l.length+1 {
		l.Append(val)
		return
	}

	// The Main Logic: Middle Insertion

	// Create New Node
	newNode := &Node{
		data: val,
	}

	// Length Increment
	l.length++

	curr := l.head
	for i := 1; i < pos-1; i++ {
		curr = curr.next	
	}
	// Hold the address of the next nodes
	tmp := curr.next
	// Attach the new Node to the current node
	curr.next = newNode
	newNode.prev = curr
	// Attach the new node the the next node
	tmp.prev = newNode
	newNode.next = tmp

}

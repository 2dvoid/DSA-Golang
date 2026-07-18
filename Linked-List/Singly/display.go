package main

import "fmt"

func (l *LinkedList) Display() {

	// If the list is empty
	if l.head == nil {
		fmt.Println("List is empty!")
		return
	}

	// If not empty
	// Start from head
	current := l.head
	// Till the last node
	for current != nil {
		// Print the value
		fmt.Printf("%v ", current.data)
		// Current becomes the Next!
		current = current.link
	}

	// Insert an extra new line at the end to look good
	fmt.Println()

}

package main

import "fmt"

func (l *LinkedList) Search(key int) {

	// No need to check if the head is empty or not.
	// Empty head contains nill which will be automatically  rejected in the loop

	current := l.head
	pos := 1
	found := false

	for current != nil {
		if current.data == key {
			fmt.Printf("Match Found at position %d\n", pos)
			found = true
		}
		pos++
		current = current.link
	}

	// If no match found
	if !found {
		fmt.Println("!! No Match Found !!")
	}

}

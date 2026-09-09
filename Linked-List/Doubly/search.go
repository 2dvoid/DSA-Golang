package main

import "fmt"

func (l *LinkedList) Search(key int) {

	// No need to check if the head is empty or not.
	// Empty head contains nill which will be automatically  rejected in the loop

	curr := l.head
	pos := 1
	found := false

	for curr != nil {
		if curr.data == key {
			found = true
			fmt.Printf("Match Found at position: %d\n ", pos)
		}
		pos++
		curr = curr.next
	}

	// If Match not found
	if found == false {
		fmt.Println("No Match Found!")
	}

}

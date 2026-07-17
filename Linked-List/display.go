package main

import "fmt"

func (l *LinkedList) Display() {

	if l.head == nil {
		fmt.Println("List is empty!")
		return
	}

	current := l.head

	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.link
	}

}

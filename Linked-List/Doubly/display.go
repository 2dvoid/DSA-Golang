package main

import "fmt"

func (l *LinkedList) Display(){
	// Case 1: List Empty
	if l.head == nil {
		fmt.Println("List is Empty. Nothing to Display.")
		return
	}
	// Case 2: List not Empty
	tmp := l.head
	for tmp != nil {
		fmt.Printf("%v ", tmp.data)	
		tmp = tmp.next
	}
}

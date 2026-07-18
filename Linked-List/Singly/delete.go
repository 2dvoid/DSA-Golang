package main
import "fmt"

func (l *LinkedList) Delete(pos int) {

	// List Empty
	if l.head == nil {
		fmt.Println("!! List is Empty !!")
		return
	}

	// Invalid Positions
	if pos < 1 || pos > l.length {
		fmt.Printf("!! Invalid Position [Min: 1, Max: %d]\n", l.length)
		return
	}

	// Length Decrement
	l.length--

	// 1 means chop off the Head
	if pos == 1 {
		// Move the Head pointer to it's next node
		// Let the Garbage Collector chop off the  old Head
		l.head = l.head.link

		// Edge Case: Head can also be the only one node in the list
		// In that case tail pointer also needs to be pointed to the head
		if l.head == nil {
			l.tail = nil
		}
		return
	}

	// Main logic:
	current := l.head
	for i := 1; i < pos-1; i++ {
		current = current.link
	}
	// Bypass the deleted node
	current.link = current.link.link

	// If current is now pointing to nil, it means we just deleted the last node.
	if current.link == nil {
		// Point the Tail pointer to the Last Node
		l.tail = current
	}

}

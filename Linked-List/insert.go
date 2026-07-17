package main
import "fmt"

func (l *LinkedList) Insert(pos int, val int) {

	// Filter out bad data immediately and exit
	if pos < 1 || pos > l.length+1 {
		fmt.Printf("!! Invalid Position [Min: 1, Max: %d] !!\n", l.length+1)
		return
	}

	// Position 1 means Prepend
	if pos == 1 {
		l.Prepend(val)
		return
	}

	// Position (length+1) means Append
	if pos == l.length+1 {
		l.Append(val)
		return
	}

	// The Main Logic: Middle Insertion

	// Create New Node
	newNode := &Node{
		data: val,
	}

	// Loop from the Head to the Previous node of the target
	current := l.head
	for i := 1; i < pos-1; i++ {
		current = current.link
	}

	// Link the New Node to the target node
	newNode.link = current.link
	// Link the Current Node to the New Node 
	current.link = newNode
	// Length Increment
	l.length++
}

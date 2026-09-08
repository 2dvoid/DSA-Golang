package main

type Node struct {
	prev *Node
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	length int
}

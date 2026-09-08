package main

type Node struct{
	data int
	link *Node
}

type LinkedList struct{
	head *Node
	tail *Node
	length int
}

package main

func main(){
	list := LinkedList{}

	list.Append(0)
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Reverse()
	list.Display()
}

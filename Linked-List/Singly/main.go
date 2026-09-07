package main

func main(){
	list := LinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(30)
	// info.Insert(1,100)
	// info.Delete(1)
	// info.Display()
	// info.Search(30)
	list.Reverse()
	list.Display()

}

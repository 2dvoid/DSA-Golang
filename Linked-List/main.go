package main

func main(){
	info := LinkedList{}

	info.Append(10)
	info.Append(20)
	info.Append(30)
	info.Append(40)

	info.Display()

	info.Insert(2, 10) // (pos, val)

	info.Delete(2) // (pos, val)

	info.Display()
}

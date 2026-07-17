package main

func main(){
	info := LinkedList{}

	info.Append(10)
	info.Append(20)
	info.Append(30)
	info.Append(40)
	info.Insert(1,100)
	info.Delete(1)
	info.Display()

}

package main

import "fmt"

func main() {
	fmt.Print("Hello \n")

	//p := Person{
	//	Name: "militska",
	//}

	art := Article{
		Name: "Tessa",
	}
	art.SetName("can change")

	printName(&art)
	//fmt.Print(Chat(p))
}

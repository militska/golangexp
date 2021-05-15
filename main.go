package main

import "fmt"

func main() {
	fmt.Print("Hello \n")

	p := Person{
		Name: "militska",
	}

	fmt.Print(Chat(p))
}

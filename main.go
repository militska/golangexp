package main

import "fmt"

type Speaker interface {
	SayName() string
}

type Person struct {
	Name string
}

func (p Person) SayName() string {
	return "Hi, my name is" + p.Name
}

func Chat(s Speaker) string {
	return "Hi, my name is " + s.SayName()
}

func main() {
	fmt.Print("Hello \n")

	p := Person{
		Name: "militska",
	}

	fmt.Print(Chat(p))
}

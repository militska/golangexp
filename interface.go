package main

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

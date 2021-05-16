package main

import "fmt"

type Article struct {
	Name        string
	Description string
	Author      string
}

type Validater interface {
	GetName() string
}

func (a *Article) SetName(name string) {
	a.Name = name
}

func (a *Article) GetName() string {
	return a.Name
}

func printName(v Validater) {
	fmt.Println(v.GetName())
}

func f() {
	ar := Article{Name: "test 1"}
	ar.SetName("change ")
}

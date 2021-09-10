package main

import "fmt"

func main() {
	fmt.Println("hhh")

	a := [4]int{3, 6, -1, 23}

	fmt.Println(a)

	s1 := a[1:3]

	fmt.Println(s1)
	s1[1] = 33

	fmt.Println(a)

}

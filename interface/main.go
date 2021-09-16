package main

import "fmt"

func main() {
	var name interface{}
	var num interface{}
	name = 344
	num = 344
	fmt.Println(name)

	name = "test"
	fmt.Println(name)

	list := []interface{}{true}

	list = append(list, name)
	list = append(list, num)

	fmt.Println(list)

}

package main

import (
	"fmt"
	"sync"
)

//var i = 0

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	var mutex = &sync.Mutex{}

	go func() {
		mutex.Lock()
		m["1"] = "a" // First conflicting access.
		mutex.Unlock()

		c <- true
	}()
	mutex.Lock()
	m["3"] = "b" // Second conflicting access.
	mutex.Unlock()

	<-c

	for k, v := range m {
		fmt.Println(k, v)
	}
}

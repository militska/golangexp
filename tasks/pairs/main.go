package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("01")

	dataS := strings.Split(string(data), "\n")
	countCommand, _ := strconv.Atoi(dataS[0])
	activeCommand := dataS[1:]

	for n := 0; n <= countCommand*2-1; n++ {
		if n%2 != 0 {
			continue
		}

		devs := strings.Split(activeCommand[n+1], " ")

		var newDevs []int
		for _, v := range devs {
			newDev, _ := strconv.Atoi(v)
			newDevs = append(newDevs, newDev)
		}
		newDevs2 := make([]int, len(newDevs))
		copy(newDevs2, newDevs)

		for i, v := range newDevs2 {
			balls := make( map[int]int)

			current := v
			fmt.Print(fmt.Sprintf("%d ", current))
			for i2, _ := range newDevs[i:] {
				balls[i2] = int(math.Abs(float64(current - newDevs[i2])))
			}

			fmt.Println(balls)

			//s1, _ := removeElement(newDevs2, i)
			//newDevs2 = s1
			//
			//fmt.Println(fmt.Sprintf("id %d", i))
			//fmt.Println(s1)

		}


		break
	}
}
func removeElement(s []int, i int) ([]int, error) {
	// s is [1,2,3,4,5,6], i is 2

	//perform bounds checking first to prevent a panic!
	if i >= len(s) || i < 0 {
		return nil, fmt.Errorf("Index is out of range. Index is %d with slice length %d", i, len(s))
	}

	newSlice := make([]int, 0)


	newSlice = append(newSlice, s[:i]...)
	newSlice = append(newSlice, s[i+1:]...)


	return newSlice, nil
}
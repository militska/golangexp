package main

import "fmt"

func main() {
	//str := "ЯЯЕЕЕ"
	str := "aaaЫЫЫbbbcccccaaaaaccЯЯ"
	var inc int
	var key string
	res := make(map[string]int)
	runStr := []rune(str)
	for inc = 0; inc < len([]rune(str)); inc++ {
		key = string(runStr[inc])
		if _, ok := res[key]; ok {
			res[key]++
			continue
		}
		res[key] = 1
	}

	for v, k := range res {
		fmt.Print(fmt.Sprintf("%s%d", v, k))
	}
}

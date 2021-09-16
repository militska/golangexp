package main

import (
	"fmt"
	"sort"
)

func Counter(str string) string {
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
	SortKey := make([]string, 0, len(res))
	for k := range res {
		SortKey = append(SortKey, k)
	}
	sort.Strings(SortKey)
	itog := ""
	for _, k := range SortKey {
		itog += fmt.Sprintf("%s%d", k, res[k])
	}
	return itog
}

func main() {
	//str := "ЯЯЕЕЕ"
	str := "ЯЯaaaЫЫЫbbbcccccaaaaaccЯЯ"

	fmt.Println(Counter(str))
}

package main

import (
	"fmt"
	"sort"
)

func Counter(str string) string {
	var inc int
	res := make(map[string]int)
	runStr := []rune(str)
	for inc = 0; inc < len([]rune(str)); inc++ {
		res[string(runStr[inc])]++
	}
	SortKey := make([]string, 0, len(res))
	for k := range res {
		SortKey = append(SortKey, k)
	}
	sort.Strings(SortKey)
	var itog string
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

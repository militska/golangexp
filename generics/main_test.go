package main

import (
	"fmt"
	"testing"
)

func BenchmarkSum(b *testing.B) {
	var n customInt64
	n = 10
	fmt.Println(MultiplyTen(n))
	fmt.Println(MultiplyTen(5.55))
}

func BenchmarkSum2(b *testing.B) {
	var n customInt64
	n = 10
	fmt.Println(MultiplyTen2(n))
	//fmt.Println(MultiplyTen2(5.55))
}

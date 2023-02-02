package main

import (
	"fmt"
)

type customInt64 int64

type Number interface {
	~int64 | float64
}

func MultiplyTen[T Number](a T) T {

	fmt.Printf("%+v are of type %T \n", a, *new(T))

	return a * 10
}

func MultiplyTen2[T Number](a T) T {

	switch (interface{})(a).(type) {
	case int64, customInt64:
		fmt.Printf("MultiplyTen2 %v: INT64 \n", a)
	case float64:
		fmt.Printf("MultiplyTen2  %v: float64", a)

	}

	return a * 10
}

func main() {
	var n customInt64
	n = 10
	fmt.Println(MultiplyTen(n))
	fmt.Println(MultiplyTen(5.55))
}

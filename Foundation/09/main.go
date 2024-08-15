package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(50, 10, 70, 80, 90, 100, 2042, 5345, 542, 42194721, 493284))
}

func sum(numeros ...int) int {
	valorTotal := 0
	for _, valor := range numeros {
		valorTotal += valor
	}

	return valorTotal
}

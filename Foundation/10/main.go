package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(50, 10, 70, 80, 90, 100, 2042, 5345, 542, 42194721, 493284) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	valorTotal := 0
	for _, valor := range numeros {
		valorTotal += valor
	}

	return valorTotal
}

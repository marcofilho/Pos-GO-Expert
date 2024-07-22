package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, error := sum(50, 10)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("Soma maior que 50")
	}

	return a + b, nil
}

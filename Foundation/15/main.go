package main

import "fmt"

func main() {

	a := 10
	var ponteiro *int = &a

	fmt.Printf("O valor da variável é: %d, e a sua posição na memória é: %x\n", a, ponteiro)

	*ponteiro = 20
	b := &a
	*b = 30

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
}

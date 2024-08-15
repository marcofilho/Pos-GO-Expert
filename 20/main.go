package main

import (
	"fmt"
	"pos-go-expert/matematica"
)

func main() {
	fmt.Println("O resultado da soma é:", matematica.Soma(10, 20))

	carro := matematica.Carro{Marca: "Honda"}
	fmt.Println(carro.Marca)

	var b int = matematica.A + 10
	fmt.Println(b)

	fmt.Println(carro.Andar())
}

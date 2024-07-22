package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	marco := Cliente{
		Nome:  "Marco",
		Idade: 30,
		Ativo: true,
	}

	marco.Ativo = false
	fmt.Printf("O nome é %s, a idade é %d, e está ativo? %t", marco.Nome, marco.Idade, marco.Ativo)
}

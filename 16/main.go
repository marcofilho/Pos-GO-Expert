package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func (c *Cliente) andou() {
	c.nome = "Marco Lima"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {

	marco := Cliente{
		nome: "Marco",
	}

	marco.andou()
	fmt.Printf("O nome do cliente é %v\n", marco.nome)

	conta := Conta{
		saldo: 100,
	}

	conta.simular(100)
	println(conta.saldo)
}

package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (cliente Cliente) Desativar() {
	cliente.Ativo = false
	fmt.Printf("O nome é %s, a idade é %d, mora em %s e está ativo? %t", cliente.Nome, cliente.Idade, cliente.Cidade, cliente.Ativo)
}

func main() {
	marco := Cliente{
		Nome:  "Marco",
		Idade: 30,
		Ativo: true,
	}

	marco.Endereco.Cidade = "Vitória" //sem composicao de structs
	marco.Cidade = "Vitória"          //com composicao de structs
	marco.Desativar()
}

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

func main() {
	marco := Cliente{
		Nome:  "Marco",
		Idade: 30,
		Ativo: true,
	}

	marco.Ativo = false
	marco.Endereco.Cidade = "Vitória" //sem composicao de structs
	marco.Cidade = "Vitória"          //com composicao de structs
	fmt.Printf("O nome é %s, a idade é %d, mora em %s e está ativo? %t", marco.Nome, marco.Idade, marco.Cidade, marco.Ativo)
}

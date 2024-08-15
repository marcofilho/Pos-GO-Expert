package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000
	fmt.Println(salarios["Wes"])

	//sal := make(map[string]int)

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}

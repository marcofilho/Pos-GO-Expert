package main

func main() {

	//for simples
	for i := 0; i < 10; i++ {
		println(i)
	}

	//for com range para arrays ou listas

	numbers := []string{"1", "2", "3", "4", "5"}

	for i, number := range numbers {
		println(i, number)
	}

	//for com blank no indice para arrays ou listas, tambem funciona para esconder o valor ao inves do indice
	for _, number := range numbers {
		println(number)
	}

}

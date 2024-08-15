package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	mapaInteiro := map[string]int{
		"Marco":  1000,
		"Jackie": 2000,
		"Bruna":  3000,
	}

	mapaFloat := map[string]float64{
		"Marco":  1000.0,
		"Jackie": 2000.0,
		"Bruna":  3000.0,
	}

	mapaMyNumber := map[string]MyNumber{
		"Marco":  1000.0,
		"Jackie": 2000.0,
		"Bruna":  3000.0,
	}

	println(Soma(mapaInteiro))
	println(Soma(mapaFloat))
	println(Soma(mapaMyNumber))
	println(Compara(10, 10))

}

package matematica

func Soma[T int | float64](a T, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "O carro está andando"
}

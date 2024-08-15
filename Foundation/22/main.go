package main

func main() {
	a := 1
	b := 2
	c := 3

	//if simples
	if a > b {
		println("a > b")
	} else {
		println("b > a")
	}

	//and operator
	if a > b && c > a {
		println("a > b && c > a")
	}

	//or operator
	if a > b || c > a {
		println("a > b && c > a")
	}

	//switch
	switch a {
	case 1:
		println("a == 1")
	case 2:
		println("a == 2")
	case 3:
		println("a == 3")
	default:
		println("a is not 1, 2 or 3")
	}
}

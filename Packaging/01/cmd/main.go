package main

import (
	"fmt"

	"github.com/marcofilho/Pos-GO-Expert/tree/main/Packaging/01/math"
)

func main() {
	m := math.Math{A: 10, B: 20}
	fmt.Println(m.Add())
}

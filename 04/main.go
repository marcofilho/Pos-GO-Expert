package main

import "fmt"

const a = "Hello, World!"

type ID int

var b bool
var c int
var d string
var e float64
var f ID

func main() {
	fmt.Printf("O tipo da variavel E é %T", f)
}

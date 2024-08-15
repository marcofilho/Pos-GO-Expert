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
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	for i, v := range myArray {
		fmt.Println("O valor do indice é %d e o conteudo é %d\n", i, v)
	}
}

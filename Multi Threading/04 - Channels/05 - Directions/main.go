package main

import "fmt"

func receive(nome string, hello chan<- string) {
	hello <- nome
}

func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}

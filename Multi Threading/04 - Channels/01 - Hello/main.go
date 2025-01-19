package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // empty

	// Thread 2
	go func() {
		canal <- "Hello World!" // full
	}()

	// Thread 1
	msg := <-canal // empty channel
	fmt.Println(msg)
}

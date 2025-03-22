package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 1000; i++ {
		f, err := os.Create(fmt.Sprintf("./tmp/.file%d.txt", i))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		f.WriteString("Hello, World!")
	}
}

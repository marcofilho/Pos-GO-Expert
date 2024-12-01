package main

import (
	"fmt"

	"github.com/marcofilho/Pos-GO-Expert/APIs/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", config)
}

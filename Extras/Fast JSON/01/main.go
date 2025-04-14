package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	var parser fastjson.Parser

	jsonData := `{"name": "John", "age": 30, "city": "New York", "active": true, "arr": [1, 2, 3]}`

	v, err := parser.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\n", v.GetStringBytes("name"))
	fmt.Printf("Age: %d\n", v.GetInt("age"))
	fmt.Printf("City: %s\n", v.GetStringBytes("city"))
	fmt.Printf("Active: %t\n", v.GetBool("active"))
	fmt.Printf("Array: %v\n", v.GetArray("arr"))

	array := v.GetArray("arr")
	for item, value := range array {
		fmt.Printf("index %d: value %s\n", item, value)
	}

}

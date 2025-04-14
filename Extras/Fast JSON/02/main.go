package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	City   string `json:"city"`
	Active bool   `json:"active"`
	Arr    []int  `json:"arr"`
}

func main() {
	var parser fastjson.Parser
	jsonData := `{
		"user": {
			"name": "John",
			"age": 30,
			"city": "New York",
			"active": true,
			"arr": [1, 2, 3]
		}
	}`

	value, err := parser.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	userJSON := value.Get("user").String()

	var user User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		panic(err)
	}

	fmt.Println("User Details:")
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Age: %d\n", user.Age)
	fmt.Printf("City: %s\n", user.City)
	fmt.Printf("Active: %t\n", user.Active)
	fmt.Printf("Array: %v\n", user.Arr)
	for i, value := range user.Arr {
		fmt.Printf("Array[%d]: %d\n", i, value)
	}

}

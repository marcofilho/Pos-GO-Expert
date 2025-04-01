package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	usecase := NewUseCase(db)

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product ID: %d, Product Name: %s\n", product.ID, product.Name)

}

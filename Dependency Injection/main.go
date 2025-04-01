package main

import (
	"database/sql"
	"fmt"

	"github.com/marcofilho/Pos-GO-Expert/DependencyInjection/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	repo := product.NewProductRepository(db)

	usecase := product.NewProductUsecase(repo)
	if err != nil {
		panic(err)
	}

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product ID: %d, Product Name: %s", product.ID, product.Name)

}

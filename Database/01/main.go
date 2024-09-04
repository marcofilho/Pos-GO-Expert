package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func CreateNewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	product := CreateNewProduct("Laptop", 1000)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product inserted successfully")

	product.Price = 1500
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product updated successfully")

	product, err = selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %s, has price: %.2f\n", product.Name, product.Price)

	products, err := selectProducts(db)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Printf("Product: %s, has price: %.2f\n", product.Name, product.Price)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product deleted successfully")
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO Products (id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE Products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT Id, Name, Price FROM Products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func selectProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT Id, Name, Price FROM Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM Products WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

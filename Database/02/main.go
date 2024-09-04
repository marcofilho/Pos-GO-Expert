package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int     `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	//create single record
	db.Create(&Product{
		Name:  "Laptop",
		Price: 1000})

	products := []Product{
		{Name: "Mouse", Price: 10},
		{Name: "Keyboard", Price: 20},
		{Name: "Monitor", Price: 200},
	}

	//create multiple records
	db.Create(&products)

	//read all records
	db.Find(&products)

	for _, product := range products {
		fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
	}
}
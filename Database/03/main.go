package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	gorm.Model
}

type Product struct {
	ID         int      `gorm:"primaryKey" json:"id"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	//create category
	// category := Category{Name: "Electronics"}
	// db.Create(&category)

	// //create product

	// product := Product{
	// 	Name:       "Laptop",
	// 	Price:      1000,
	// 	CategoryID: category.ID,
	// }
	// db.Create(&product)

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Printf("Product: %s, Price: %.2f, Category: %s\n", product.Name, product.Price, product.Category.Name)
	}

}

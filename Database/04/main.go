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
	ID           int          `gorm:"primaryKey" json:"id"`
	Name         string       `json:"name"`
	Price        float64      `json:"price"`
	CategoryID   int          `json:"category_id"`
	Category     Category     `json:"category"`
	SerialNumber SerialNumber `json:"serial_number"`
	gorm.Model
}

type SerialNumber struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Number    string `json:"number"`
	ProductID int    `json:"product_id"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	//create category
	category := Category{Name: "Electronics"}
	db.Create(&category)

	//create product
	product := Product{
		Name:       "Laptop",
		Price:      1000,
		CategoryID: category.ID,
	}
	db.Create(&product)

	//create serial number
	serialNumber := SerialNumber{
		Number:    "123456",
		ProductID: 1,
	}
	db.Create(&serialNumber)

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Printf("Product: %s, Price: %.2f, Category: %s, SerialNumber: %s\n", product.Name, product.Price, product.Category.Name, product.SerialNumber.Number)
	}

}

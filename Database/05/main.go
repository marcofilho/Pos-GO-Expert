package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
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
	// category := Category{Name: "Electronics"}
	// db.Create(&category)

	category := Category{Name: "Kitchen"}
	db.Create(&category)

	//create product
	// product := Product{
	// 	Name:       "Laptop",
	// 	Price:      1000,
	// 	CategoryID: category.ID,
	// }
	// db.Create(&product)

	product2 := Product{
		Name:       "Knife",
		Price:      50.0,
		CategoryID: 2,
	}
	db.Create(&product2)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}
}

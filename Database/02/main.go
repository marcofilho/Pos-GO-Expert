package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int     `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	//create single record
	// db.Create(&Product{
	// 	Name:  "Laptop",
	// 	Price: 1000})

	// products := []Product{
	// 	{Name: "Laptop", Price: 1500},
	// 	{Name: "Mouse", Price: 10},
	// 	{Name: "Keyboard", Price: 20},
	// 	{Name: "Monitor", Price: 200},
	// }

	// //create multiple records
	// db.Create(&products)

	//read all records
	//db.Find(&products)

	// for _, product := range products {
	// 	fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
	// }

	//find records with where conditions

	//var products []Product
	// db.Where("price > ?", 50).Find(&products)
	// for _, product := range products {
	// 	fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
	// }

	//find strings records with like condition

	// db.Where("name LIKE ?", "%k%").Find(&products)
	// for _, product := range products {
	// 	fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
	// }

	// var product Product
	// db.First(&product, 1)
	// product.Name = "New Laptop"
	// db.Save(&product)

	// var product2 Product
	// db.First(&product2, 1)
	// fmt.Println(product2.Name, product2.Price)

	// db.Delete(&product2)
}

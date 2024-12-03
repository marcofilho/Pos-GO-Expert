package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marcofilho/Pos-GO-Expert/APIs/configs"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/entity"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/infra/database"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products", productHandler.GetProducts)
	router.Get("/products/{id}", productHandler.GetProduct)
	router.Put("/products/{id}", productHandler.UpdateProduct)
	router.Delete("/products/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(":8000", router)
}

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/marcofilho/Pos-GO-Expert/APIs/configs"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/entity"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/infra/database"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JwtExpiresIn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/products", func(r chi.Router) {
		router.Use(jwtauth.Verifier(config.TokenAuth))
		router.Use(jwtauth.Authenticator)
		router.Post("/", productHandler.CreateProduct)
		router.Get("/", productHandler.GetProducts)
		router.Get("/{id}", productHandler.GetProduct)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", router)
}

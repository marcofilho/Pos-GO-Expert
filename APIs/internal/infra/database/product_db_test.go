package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupProductTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	return db
}

func TestProductDB(t *testing.T) {
	t.Run("TestNewProduct", func(t *testing.T) {
		db := setupProductTestDB(t)
		product, err := entity.NewProduct("Car", 100.00)
		assert.Nil(t, err)

		productDB := NewProduct(db)
		err = productDB.Create(product)
		assert.NoError(t, err)
		assert.NotEmpty(t, product.ID)
	})

	t.Run("TestFindAllProducts", func(t *testing.T) {
		db := setupProductTestDB(t)
		for i := 1; i < 24; i++ {
			product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
			assert.NoError(t, err)
			db.Create(product)
		}

		productDB := NewProduct(db)
		products, err := productDB.FindAll(1, 10, "asc")
		assert.NoError(t, err)
		assert.Len(t, products, 10)
		assert.Equal(t, "Product 1", products[0].Name)
		assert.Equal(t, "Product 10", products[9].Name)

		products, err = productDB.FindAll(2, 10, "asc")
		assert.NoError(t, err)
		assert.Len(t, products, 10)
		assert.Equal(t, "Product 11", products[0].Name)
		assert.Equal(t, "Product 20", products[9].Name)

		products, err = productDB.FindAll(3, 10, "asc")
		assert.NoError(t, err)
		assert.Len(t, products, 3)
		assert.Equal(t, "Product 21", products[0].Name)
		assert.Equal(t, "Product 23", products[2].Name)
	})

	t.Run("TestFindProductByID", func(t *testing.T) {
		db := setupProductTestDB(t)
		product, err := entity.NewProduct("Car", 100.00)
		assert.NoError(t, err)
		db.Create(product)

		productDB := NewProduct(db)
		product, err = productDB.FindByID(product.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, "Car", product.Name)
		assert.Equal(t, 100.00, product.Price)
	})

	t.Run("TestUpdateProduct", func(t *testing.T) {
		db := setupProductTestDB(t)
		product, err := entity.NewProduct("Car", 100.00)
		assert.NoError(t, err)
		db.Create(product)

		productDB := NewProduct(db)
		product.Name = "Car 2"
		product.Price = 200.00
		err = productDB.Update(product)
		assert.NoError(t, err)

		product, err = productDB.FindByID(product.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, "Car 2", product.Name)
		assert.Equal(t, 200.00, product.Price)
	})

	t.Run("TestDeleteProduct", func(t *testing.T) {
		db := setupProductTestDB(t)
		product, err := entity.NewProduct("Car", 100.00)
		assert.NoError(t, err)
		db.Create(product)

		productDB := NewProduct(db)
		err = productDB.Delete(product.ID.String())
		assert.NoError(t, err)

		_, err = productDB.FindByID(product.ID.String())
		assert.Error(t, err)
	})
}

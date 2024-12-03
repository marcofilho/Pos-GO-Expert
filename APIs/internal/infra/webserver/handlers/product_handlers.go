package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/dto"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/entity"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productInput dto.CreateProductInput

	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := entity.NewProduct(productInput.Name, productInput.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

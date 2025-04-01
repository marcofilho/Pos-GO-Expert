//go:build wire_gen
// +build wire_gen

//go:generate wire

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/marcofilho/Pos-GO-Expert/DependencyInjection/product"
)

// NewUseCase wires the dependencies for ProductUseCase.
func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{} // This is required for compilation.
}

package product

type ProductUsecase struct {
	repo ProductRepository
}

func NewProductUsecase(repo ProductRepository) ProductUsecase {
	return ProductUsecase{
		repo: repo,
	}
}

func (u *ProductUsecase) GetProduct(id int) (Product, error) {
	product, err := u.repo.GetProductByID(id)
	if err != nil {
		return Product{}, nil
	}
	return product, nil
}

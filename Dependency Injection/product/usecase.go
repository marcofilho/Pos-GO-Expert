package product

type ProductUsecase struct {
	repo *ProductRepository
}

func NewProductUsecase(repo *ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		repo: repo,
	}
}

func (u *ProductUsecase) GetProduct(id int) (*Product, error) {
	return u.repo.GetProductByID(id)
}

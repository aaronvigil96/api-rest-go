package product

type ProductService struct {
	repository *ProductRepository
}

func NewProductService(repository *ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) GetProducts() ([]Product, error) {
	return s.repository.GetProducts()
}

func (s *ProductService) GetProductByID(id int) (*Product, error) {
	return s.repository.GetProductByID(id)
}

func (s *ProductService) CreateProduct(product Product) (Product, error) {
	return s.repository.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(id int, product Product) (*Product, error) {
	return s.repository.UpdateProduct(id, product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repository.DeleteProduct(id)
}

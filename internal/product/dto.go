package product

type CreateProductDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateProductDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

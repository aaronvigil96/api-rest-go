package product

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductController struct {
	service *ProductService
}

func NewProductController(service *ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.service.GetProducts()

	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	product, err := c.service.GetProductByID(id)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProductDto CreateProductDTO

	err := json.NewDecoder(r.Body).Decode(&createProductDto)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	product := Product{
		Name:        createProductDto.Name,
		Description: createProductDto.Description,
	}

	createdProduct, err := c.service.CreateProduct(product)

	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(createdProduct)
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var product Product

	err = json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	updatedProduct, err := c.service.UpdateProduct(id, product)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}

	err = c.service.DeleteProduct(id)

	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package routes

import (
	"net/http"
	"prueba/internal/product"
)

func SetupRoutes(productController *product.ProductController) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", productController.GetProducts)
	mux.HandleFunc("GET /products/{id}", productController.GetProductByID)
	mux.HandleFunc("POST /products", productController.CreateProduct)
	mux.HandleFunc("PATCH /products/{id}", productController.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", productController.DeleteProduct)

	return mux
}

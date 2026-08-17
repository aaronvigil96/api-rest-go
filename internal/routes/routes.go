package routes

import (
	"net/http"
	"prueba/internal/auth"
	"prueba/internal/middleware"
	"prueba/internal/product"

	"github.com/rs/cors"
)

func SetupRoutes(productController *product.ProductController, authController *auth.AuthController) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", productController.GetProducts)
	mux.HandleFunc("GET /products/{id}", productController.GetProductByID)

	mux.Handle(
		"POST /products",
		middleware.JWTMiddleware(http.HandlerFunc(productController.CreateProduct)),
	)

	mux.Handle(
		"PATCH /products/{id}",
		middleware.JWTMiddleware(http.HandlerFunc(productController.UpdateProduct)),
	)

	mux.Handle(
		"DELETE /products/{id}",
		middleware.JWTMiddleware(http.HandlerFunc(productController.DeleteProduct)),
	)

	mux.HandleFunc("POST /auth/register", authController.Register)
	mux.HandleFunc("POST /auth/login", authController.Login)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	}).Handler(mux)

	return handler
}

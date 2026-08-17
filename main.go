package main

import (
	"fmt"
	"net/http"
	"prueba/internal/auth"
	"prueba/internal/database"
	"prueba/internal/product"
	"prueba/internal/routes"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	productRepository := product.NewProductRepository(db)
	productService := product.NewProductService(productRepository)
	productController := product.NewProductController(productService)

	authRepository := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepository)
	authController := auth.NewAuthController(authService)

	router := routes.SetupRoutes(productController, authController)

	fmt.Println("Servidor escuchando en http://localhost:8080")

	err = http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}

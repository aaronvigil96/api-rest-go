package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"prueba/internal/auth"
	"prueba/internal/database"
	"prueba/internal/product"
	"prueba/internal/routes"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	email := "admin@admin.com"
	password := "admin123"

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		context.Background(),
		`INSERT INTO users (email, password_hash, role)
	 VALUES ($1, $2, $3)
	 ON CONFLICT (email) DO NOTHING`,
		email,
		string(hash),
		"admin",
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Administrador creado correctamente")

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

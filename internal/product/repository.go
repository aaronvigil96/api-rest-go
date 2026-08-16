package product

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProducts() ([]Product, error) {
	rows, err := r.db.Query(
		context.Background(),
		"SELECT id, name, description FROM products",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id int) (*Product, error) {
	var product Product

	err := r.db.QueryRow(
		context.Background(),
		"SELECT id, name, description FROM products WHERE id = $1",
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) CreateProduct(product Product) (Product, error) {
	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO products (name, description) VALUES ($1, $2) RETURNING id, name, description`,
		product.Name,
		product.Description,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
	)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *ProductRepository) UpdateProduct(id int, product Product) (*Product, error) {
	updatedProduct := &Product{}

	err := r.db.QueryRow(
		context.Background(),
		`UPDATE products SET name = $1, description = $2 WHERE id = $3 RETURNING id, name, description`,
		product.Name,
		product.Description,
		id,
	).Scan(
		&updatedProduct.ID,
		&updatedProduct.Name,
		&updatedProduct.Description,
	)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	_, err := r.db.Exec(
		context.Background(),
		`DELETE FROM products WHERE id = $1`,
		id,
	)

	return err
}

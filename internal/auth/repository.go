package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Register(user User) (User, error) {

	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO users(email, password_hash) VALUES($1, $2) RETURNING id, email, password_hash, created_at, role`,
		user.Email,
		user.Password_Hash,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password_Hash,
		&user.Created_At,
		&user.Role,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *AuthRepository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, email, password_hash, role, created_at FROM users WHERE email = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password_Hash,
		&user.Role,
		&user.Created_At,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *AuthRepository) GetCurrentUser(id int) (UserResponse, error) {
	var user UserResponse

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, email, role, created_at
		 FROM users
		 WHERE id = $1`,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Role,
		&user.Created_At,
	)

	if err != nil {
		return UserResponse{}, err
	}

	return user, nil
}

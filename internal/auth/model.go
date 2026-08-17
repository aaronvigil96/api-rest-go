package auth

import "time"

type User struct {
	ID            int
	Email         string
	Password_Hash string
	Role          string
	Created_At    time.Time
}

type UserResponse struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Created_At time.Time `json:"created_at"`
}

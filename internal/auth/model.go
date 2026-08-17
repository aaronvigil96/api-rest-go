package auth

import "time"

type User struct {
	ID            int
	Email         string
	Password_Hash string
	Role          string
	Created_At    time.Time
}

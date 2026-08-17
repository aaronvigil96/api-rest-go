package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository *AuthRepository
}

func NewAuthService(repository *AuthRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) Register(registerDto RegisterDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(registerDto.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	register := User{
		Email:         registerDto.Email,
		Password_Hash: string(hashedPassword),
	}

	_, err = s.repository.Register(register)

	fmt.Printf("Error en INSERT: %v\n", err)

	return err
}

func (s *AuthService) Login(loginDto LoginDTO) (string, error) {
	user, err := s.repository.FindByEmail(loginDto.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password_Hash), []byte(loginDto.Password))

	if err != nil {
		return "", err
	}

	token, err := GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) GetCurrentUser(id int) (UserResponse, error) {
	return s.repository.GetCurrentUser(id)
}

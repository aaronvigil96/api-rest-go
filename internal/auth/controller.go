package auth

import (
	"encoding/json"
	"net/http"
	"prueba/internal/requestcontext"
)

type AuthController struct {
	service *AuthService
}

func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var registerDto RegisterDTO

	err := json.NewDecoder(r.Body).Decode(&registerDto)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = c.service.Register(registerDto)

	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginDto LoginDTO

	err := json.NewDecoder(r.Body).Decode(&loginDto)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	token, err := c.service.Login(loginDto)

	if err != nil {
		http.Error(w, "Invalid credencials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (c *AuthController) Me(w http.ResponseWriter, r *http.Request) {

	userID, ok := requestcontext.GetUserID(r.Context())

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := c.service.GetCurrentUser(userID)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

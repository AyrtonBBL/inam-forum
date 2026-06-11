package controllers

import (
	"encoding/json"
	"inam-forum/models"
	"inam-forum/services"
	"net/http"
	"os"
)

type AuthController struct {
	authService *services.AuthService
}

// InitAuthController il va initialise le contrôleur d'authentification
func InitAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// RegisterHandler va gèrer la requête HTTP POST /api/register
func (c *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// On force le type de réponse en JSON
	w.Header().Set("Content-Type", "application/json")

	var req models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Données JSON invalides"})
		return
	}

	if req.NomUtilisateur == "" || req.Email == "" || req.MotDePasse == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Tous les champs sont obligatoires"})
		return
	}

	newUser, err := c.authService.Register(req)
	if err != nil {
		w.WriteHeader(http.StatusConflict) // Email déjà pris par exemple
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// LoginHandler y gère la requête HTTP POST /api/login
func (c *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Données JSON invalides"})
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	authResponse, err := c.authService.Login(req, jwtSecret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authResponse)
}

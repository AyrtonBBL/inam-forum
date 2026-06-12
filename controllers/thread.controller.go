package controllers

import (
	"encoding/json"
	"inam-forum/models"
	"inam-forum/services"
	"net/http"
)

type ThreadController struct {
	service *services.ThreadService
}

func InitThreadController(service *services.ThreadService) *ThreadController {
	return &ThreadController{service: service}
}

func (c *ThreadController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Récupération de l'ID utilisateur extrait par le Middleware
	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Action non autorisée"})
		return
	}

	var req models.ThreadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "JSON invalide"})
		return
	}

	if req.Titre == "" || req.Description == "" || req.IDJeu == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Tous les champs sont obligatoires"})
		return
	}

	newThread, err := c.service.CreateThread(req, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Impossible de créer l'annonce"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newThread)
}

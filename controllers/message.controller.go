package controllers

import (
	"encoding/json"
	"inam-forum/models"
	"inam-forum/services"
	"net/http"
)

type MessageController struct {
	service *services.MessageService
}

func InitMessageController(service *services.MessageService) *MessageController {
	return &MessageController{service: service}
}

func (c *MessageController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Récupération sécurisée de l'ID utilisateur via le middleware
	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Action non autorisée"})
		return
	}

	var req models.MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "JSON invalide"})
		return
	}

	if req.Contenu == "" || req.IDFil == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Le contenu et l'ID du fil sont obligatoires"})
		return
	}

	newMsg, err := c.service.CreateMessage(req, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Impossible d'envoyer le message"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMsg)
}

package controllers

import (
	"encoding/json"
	"inam-forum/models"
	"inam-forum/services"
	"net/http"
)

type ReactionController struct {
	service *services.ReactionService
}

func InitReactionController(service *services.ReactionService) *ReactionController {
	return &ReactionController{service: service}
}

func (c *ReactionController) VoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Action non autorisée"})
		return
	}

	var req models.ReactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "JSON invalide"})
		return
	}

	if req.Type == "" || req.MessageID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Le type et l'message_id sont obligatoires"})
		return
	}

	err := c.service.Vote(req, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Réaction enregistrée avec succès"})
}

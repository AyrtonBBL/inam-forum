package controllers

import (
	"encoding/json"
	"inam-forum/models"
	"inam-forum/services"
	"net/http"
)

type AdminController struct {
	adminService *services.AdminService
}

func InitAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
}

func (c *AdminController) BanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Récupération de l'ID de l'admin avc le middleware
	adminID, ok := r.Context().Value("user_id").(string)
	if !ok || adminID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Action non autorisée"})
		return
	}

	var req models.BanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "JSON invalide"})
		return
	}

	if req.UserID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "L'ID de l'utilisateur à modérer est obligatoire"})
		return
	}

	// Appel du service
	err := c.adminService.ModerationUser(req, adminID)
	if err != nil {
		w.WriteHeader(http.StatusForbidden) // 403: Accès refusé
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Statut de modération mis à jour avec succès"})
}

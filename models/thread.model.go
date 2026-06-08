package models

import "time"

// Thread pr un fil de discussion (Annonce de Mate ou Partage d'exploit)
type Thread struct {
	ID            string    `json:"id"`
	Titre         string    `json:"titre"`
	Description   string    `json:"description"`
	Etat          string    `json:"etat"`           // "ouvert", "fermé", "archivé"
	DateCreation  time.Time `json:"date_creation"`
	UserID        string    `json:"user_id"`        // Clé étrangère vers l'auteur
	CategoryID    string    `json:"category_id"`    // Clé étrangère vers le jeu concerné
}
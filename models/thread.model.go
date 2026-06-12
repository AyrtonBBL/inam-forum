package models

import "time"

// Thread représente un fil de discussion 
type Thread struct {
	ID           string    `json:"id"`
	Titre        string    `json:"titre"`
	Description  string    `json:"description"`
	Etat         string    `json:"etat"` // "ouvert", "fermé", "archivé"
	CreatedAt    time.Time `json:"created_at"`
	UserID       string    `json:"user_id"` 
}

// ThreadRequest est la structure reçue en JSON lors de la création d'une annonce
type ThreadRequest struct {
	Titre       string `json:"titre"`
	Description string `json:"description"`
	IDJeu       string `json:"id_jeu"` 
}

package models

import "time"

// Message pr une réponse ou un commentaire dans un fil
type Message struct {
	ID        string    `json:"id"`
	Contenu   string    `json:"contenu"`
	DateEnvoi time.Time `json:"date_envoi"`
	ThreadID  string    `json:"thread_id"` // Le fil auquel il appartient
	UserID    string    `json:"user_id"`   // L'auteur du message
	Score     int       `json:"score"`     // Calculé depuis les likes/dislikes (notre tache FT-6)
}

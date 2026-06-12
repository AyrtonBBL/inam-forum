package models

import "time"

// Message représente une réponse dans un fil de discussion
type Message struct {
	ID        string    `json:"id"`
	Contenu   string    `json:"contenu"`
	DateEnvoi time.Time `json:"date_envoi"`
	ThreadID  string    `json:"thread_id"`
	UserID    string    `json:"user_id"`
	Score     int       `json:"score"` // Pour les Likes / Dislikes
}

// MessageRequest représente la structure du JSON reçu lors de l'envoi d'un message
type MessageRequest struct {
	Contenu  string `json:"contenu"`
	ThreadID string `json:"thread_id"`
}

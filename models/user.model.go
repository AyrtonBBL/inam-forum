package models

import "time"

// cet struct kyky represente "User"  un gamer sur notre plateforme "INAM"
type User struct {
	ID             string    `json:"id"`
	NomUtilisateur string    `json:"nom_utilisateur"` // Doit être unique
	Email          string    `json:"email"`           // Doit être unique
	MotPasseHashe  string    `json:"-"`               // Le "-" masque le mot de passe hashé en SHA-512
	Role           string    `json:"role"`            // "user" ou "admin"
	EstBanni       bool      `json:"est_banni"`       // Géré par l'admin
	CreatedAt      time.Time `json:"created_at"`
}

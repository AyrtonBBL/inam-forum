package models

import "github.com/golang-jwt/jwt/v5"

// ici, c'est les informations sécurisées contenue dans le jeton JWT
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"` // user ou admin
	jwt.RegisteredClaims
}

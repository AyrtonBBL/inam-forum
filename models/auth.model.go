package models

// RegisterRequest est le JSON envoyé par le formulaire d'inscription
type RegisterRequest struct {
	NomUtilisateur string `json:"nom_utilisateur"`
	Email          string `json:"email"`
	MotDePasse     string `json:"mot_passe"` 
}

// LoginRequest est le JSON envoyé par le formulaire de connexion
type LoginRequest struct {
	Identifiant string `json:"identifiant"`
	MotDePasse  string `json:"mot_passe"`
}

// AuthResponse est ce que le serveur renvoie quand on se connecte
type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
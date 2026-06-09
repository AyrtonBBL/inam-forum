package models

// Ayrton, ça c'est le RegisterRequest dedans ya les les données envoyées pendant de l'inscription
type RegisterRequest struct {
	NomUtilisateur string `json:"nom_utilisateur"`
	Email          string `json:"email"`
	MotDePasse     string `json:"mot_de_passe"`
}

// Ca c'est les identifiants pour la connexion
type LoginRequest struct {
	Identifiant string `json:"identifiant"` // nom d'utilisateur ou e-mail
	MotDePasse  string `json:"mot_de_passe"`
}

// la réponse après une connection réussi
type AuthResponse struct {
	Token string `json:"token"` // Le token JWT généré
	User  User   `json:"user"`  // Les infos de l'utilisateur connecté (sans le mot de passe)
}

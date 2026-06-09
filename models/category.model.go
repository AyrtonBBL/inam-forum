package models

// Category pr un jeu vidéo sur la plateforme (ex: Valorant, Apex, fortnite hihihih)
type Category struct {
	ID        string `json:"id"`
	NomJeu    string `json:"nom_jeu"`    // ex: "League of Legends"
	Genre     string `json:"genre"`      // ex: "MOBA"
}
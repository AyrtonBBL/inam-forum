package services

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"inam-forum/models"
	"inam-forum/repositories"
	"time"

	"github.com/google/uuid" // kyk's se truc la y permet de générer des ID uniques au format String
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

// InitAuthService il initialise le service d'authentification
func InitAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Register ca gère la logique de création d'un compte
func (s *AuthService) Register(req models.RegisterRequest) (*models.User, error) {
	//  On verifie si l'email est déjà utilisé
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("cet email est déjà utilisé par un autre gamer")
	}

	// la on doit hasher le mot de passe en SHA-512 pour la sécurité
	hasher := sha512.New()
	hasher.Write([]byte(req.MotDePasse))
	motPasseHashe := hex.EncodeToString(hasher.Sum(nil))

	// on prép le modèle complet de l'utilisateur
	newUser := &models.User{
		ID:             uuid.New().String(), // Génère un ID unique
		NomUtilisateur: req.NomUtilisateur,
		Email:          req.Email,
		MotPasseHashe:  motPasseHashe,
		Role:           "user", // Ca c son rôle par défaut
		EstBanni:       false,
		CreatedAt:      time.Now(),
	}

	// Ici on enregistre en base de données via le Repository
	err = s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

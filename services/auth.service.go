package services

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"inam-forum/models"
	"inam-forum/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

// InitAuthService y initialise le service d'authentification
func InitAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Register il gère la logique de création d'un compte
func (s *AuthService) Register(req models.RegisterRequest) (*models.User, error) {

	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("cet email est déjà utilisé par un autre gamer")
	}

	//  il faut hasher le mot de passe en SHA-512 pour la sécurité
	hasher := sha512.New()
	hasher.Write([]byte(req.MotDePasse))
	motPasseHashe := hex.EncodeToString(hasher.Sum(nil))

	newUser := &models.User{
		ID:             uuid.New().String(),
		NomUtilisateur: req.NomUtilisateur,
		Email:          req.Email,
		MotPasseHashe:  motPasseHashe,
		Role:           "user",
		EstBanni:       false,
		CreatedAt:      time.Now(),
	}

	err = s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *AuthService) Login(req models.LoginRequest, jwtSecret string) (*models.AuthResponse, error) {

	user, err := s.userRepo.GetByEmail(req.Identifiant)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("identifiants incorrects")
	}
	if user.EstBanni {
		return nil, errors.New("ce compte a été suspendu par la modération")
	}

	hasher := sha512.New()
	hasher.Write([]byte(req.MotDePasse))
	motPasseHashe := hex.EncodeToString(hasher.Sum(nil))

	if user.MotPasseHashe != motPasseHashe {
		return nil, errors.New("identifiants incorrects")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: tokenString,
		User:  *user,
	}, nil
}

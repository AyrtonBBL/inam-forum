package services

import (
	"errors"
	"inam-forum/models"
	"inam-forum/repositories"
)

type AdminService struct {
	userRepo *repositories.UserRepository
}

func InitAdminService(userRepo *repositories.UserRepository) *AdminService {
	return &AdminService{userRepo: userRepo}
}

func (s *AdminService) ModerationUser(req models.BanRequest, adminID string) error {
	// vérifier si l'exécuteur est bien un admin
	admin, err := s.userRepo.GetByID(adminID)
	if err != nil || admin == nil {
		return errors.New("administrateur introuvable")
	}

	if admin.Role != "admin" {
		return errors.New("accès refusé : privilèges insuffisants")
	}

	// appliquer le ban
	return s.userRepo.BanUser(req.UserID, req.EstBanni)
}

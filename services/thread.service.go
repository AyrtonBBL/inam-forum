package services

import (
	"inam-forum/models"
	"inam-forum/repositories"
	"time"

	"github.com/google/uuid"
)

type ThreadService struct {
	repo *repositories.ThreadRepository
}

func InitThreadService(repo *repositories.ThreadRepository) *ThreadService {
	return &ThreadService{repo: repo}
}

func (s *ThreadService) CreateThread(req models.ThreadRequest, userID string) (*models.Thread, error) {
	newThread := &models.Thread{
		ID:          uuid.New().String(),
		Titre:       req.Titre,
		Description: req.Description,
		Etat:        "ouvert",
		CreatedAt:   time.Now(),
		UserID:      userID,
	}

	err := s.repo.Create(newThread, req.IDJeu)
	if err != nil {
		return nil, err
	}

	return newThread, nil
}

func (s *ThreadService) GetAllThreads() ([]models.Thread, error) {
	return s.repo.GetAll()
}

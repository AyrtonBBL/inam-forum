package services

import (
	"inam-forum/models"
	"inam-forum/repositories"
	"time"

	"github.com/google/uuid"
)

type MessageService struct {
	repo *repositories.MessageRepository
}

func InitMessageService(repo *repositories.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(req models.MessageRequest, userID string) (*models.Message, error) {
	newMessage := &models.Message{
		ID:        uuid.New().String(),
		Contenu:   req.Contenu,
		DateEnvoi: time.Now(),
		IDFil:     req.IDFil,
		UserID:    userID,
		Score:     0, // Un nouveau message commence à 0 (Likes/Dislikes FT-6)
	}

	err := s.repo.Create(newMessage)
	if err != nil {
		return nil, err
	}

	return newMessage, nil
}

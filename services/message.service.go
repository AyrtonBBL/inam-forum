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
		ThreadID:  req.ThreadID, 
		UserID:    userID,
		Score:     0,
	}

	err := s.repo.Create(newMessage)
	if err != nil {
		return nil, err
	}

	return newMessage, nil
}

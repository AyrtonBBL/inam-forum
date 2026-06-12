package services

import (
	"errors"
	"inam-forum/models"
	"inam-forum/repositories"

	"github.com/google/uuid"
)

type ReactionService struct {
	repo *repositories.ReactionRepository
}

func InitReactionService(repo *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{repo: repo}
}

func (s *ReactionService) Vote(req models.ReactionRequest, userID string) error {
	// Validation du type de vote
	if req.Type != "like" && req.Type != "dislike" {
		return errors.New("le type de réaction doit être 'like' ou 'dislike'")
	}

	reaction := &models.Reaction{
		ID:        uuid.New().String(),
		Type:      req.Type,
		UserID:    userID,
		MessageID: req.MessageID,
	}

	return s.repo.SaveReaction(reaction)
}

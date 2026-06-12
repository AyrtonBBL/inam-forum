package repositories

import (
	"database/sql"
	"inam-forum/models"
)

type MessageRepository struct {
	db *sql.DB
}

func InitMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// Create sa sert a insère un nouveau message lié à un thread et un utilisateur
func (r *MessageRepository) Create(msg *models.Message) error {
	query := `INSERT INTO message (id_message, contenu, date_envoi, id_fil, id_user, score) 
			  VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, msg.ID, msg.Contenu, msg.DateEnvoi, msg.ThreadID, msg.UserID, msg.Score)
	return err
}

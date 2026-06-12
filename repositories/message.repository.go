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
// GetByThreadID y récupère tous les messages associés à une annonce spécifique
func (r *MessageRepository) GetByThreadID(threadID string) ([]models.Message, error) {
	query := `SELECT id_message, contenu, date_envoi, id_fil, id_user, score 
			  FROM message WHERE id_fil = ? ORDER BY date_envoi ASC`
	
	rows, err := r.db.Query(query, threadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.Contenu, &m.DateEnvoi, &m.ThreadID, &m.UserID, &m.Score); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}
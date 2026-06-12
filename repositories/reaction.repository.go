package repositories

import (
	"database/sql"
	"inam-forum/models"
)

type ReactionRepository struct {
	db *sql.DB
}

func InitReactionRepository(db *sql.DB) *ReactionRepository {
	return &ReactionRepository{db: db}
}

// SaveReaction ca ajoute ou remplace un vote, puis met à jour le score du message
func (r *ReactionRepository) SaveReaction(reaction *models.Reaction) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	queryReaction := `INSERT INTO reaction (id_reaction, type_reaction, id_user, id_message) 
					  VALUES (?, ?, ?, ?)
					  ON CONFLICT(id_user, id_message) DO UPDATE SET type_reaction = excluded.type_reaction`
	
	_, err = tx.Exec(queryReaction, reaction.ID, reaction.Type, reaction.UserID, reaction.MessageID)
	if err != nil {
		tx.Rollback()
		return err
	}

	var likes, dislikes int
	queryCount := `SELECT 
					COUNT(CASE WHEN type_reaction = 'like' THEN 1 END),
					COUNT(CASE WHEN type_reaction = 'dislike' THEN 1 END)
				   FROM reaction WHERE id_message = ?`
	
	err = tx.QueryRow(queryCount, reaction.MessageID).Scan(&likes, &dislikes)
	if err != nil {
		tx.Rollback()
		return err
	}

	newScore := likes - dislikes

	queryUpdateMessage := `UPDATE message SET score = ? WHERE id_message = ?`
	_, err = tx.Exec(queryUpdateMessage, newScore, reaction.MessageID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
package repositories

import (
	"database/sql"
	"inam-forum/models"
)

type ThreadRepository struct {
	db *sql.DB
}

func InitThreadRepository(db *sql.DB) *ThreadRepository {
	return &ThreadRepository{db: db}
}

// Create va insère l'annonce et fait la liaison avec le jeu choisi
func (r *ThreadRepository) Create(thread *models.Thread, gameID string) error {
	// Début d'une transaction SQL pour être sûr d'insérer partout ou nulle part
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// et l'insertion de l'annonce principale la
	queryThread := `INSERT INTO fil_discussion (id_fil, titre, description, etat, id_user) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.Exec(queryThread, thread.ID, thread.Titre, thread.Description, thread.Etat, thread.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryMatch := `INSERT INTO correspondre (id_fil, id_jeu) VALUES (?, ?)`
	_, err = tx.Exec(queryMatch, thread.ID, gameID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
	
}

// GetAll récupère tous les fils de discussion de la base de données en gros kyks
func (r *ThreadRepository) GetAll() ([]models.Thread, error) {
	query := `SELECT id_fil, titre, description, etat, date_creation, id_user FROM fil_discussion`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var t models.Thread
		if err := rows.Scan(&t.ID, &t.Titre, &t.Description, &t.Etat, &t.CreatedAt, &t.UserID); err != nil {
			return nil, err
		}
		threads = append(threads, t)
	}
	return threads, nil
}
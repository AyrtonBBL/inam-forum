package repositories

import (
	"database/sql"
	"inam-forum/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func InitCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// GetAll y récupère tous les jeux de la BDD
func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	query := `SELECT id_jeu, nom_jeu, genre FROM categorie_jeu`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.NomJeu, &cat.Genre); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

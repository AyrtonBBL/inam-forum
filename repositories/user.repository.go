package repositories

import (
	"database/sql"
	"errors"
	"inam-forum/models"
)

type UserRepository struct {
	db *sql.DB
}

// InitUserRepository y va initialise le dépôt pour les utilisateurs
func InitUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create ca va insère un nouvel utilisateur dans la base de données
func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO utilisateur (id_user, nom_utilisateur, email, mot_passe_hashe, role, est_banni) 
			  VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, user.ID, user.NomUtilisateur, user.Email, user.MotPasseHashe, user.Role, user.EstBanni)
	return err
}

// GetByEmail y cherche un utilisateur par son email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `SELECT id_user, nom_utilisateur, email, mot_passe_hashe, role, est_banni, created_at 
			  FROM utilisateur WHERE email = ?`

	row := r.db.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.NomUtilisateur, &user.Email, &user.MotPasseHashe, &user.Role, &user.EstBanni, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// BanUser y change le statut d'un utilisateur pour le bannir ou le débannir
func (r *UserRepository) BanUser(userID string, estBanni bool) error {
	val := 0
	if estBanni {
		val = 1
	}
	query := `UPDATE utilisateur SET est_banni = ? WHERE id_user = ?`
	_, err := r.db.Exec(query, val, userID)
	return err
}

// GetByID cherche un utilisateur par son ID (utile pour vérifier le rôle de l'admin)
func (r *UserRepository) GetByID(id string) (*models.User, error) {
	query := `SELECT id_user, nom_utilisateur, email, mot_passe_hashe, role, est_banni, created_at 
			  FROM utilisateur WHERE id_user = ?`
	
	row := r.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.NomUtilisateur, &user.Email, &user.MotPasseHashe, &user.Role, &user.EstBanni, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

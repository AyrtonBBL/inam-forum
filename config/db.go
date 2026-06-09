package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// La c'est pour se connecter à la bdr
func InitDB() *sql.DB {
	driver := GetEnv("DB_DRIVER", "sqlite3")
	source := GetEnv("DB_SOURCE", "./database/inam_forum.db")

	db, err := sql.Open(driver, source)
	if err != nil {
		log.Fatalf("Impossible d'ouvrir la base de données: %v", err)
	}

	// vérification de la connexion
	if err = db.Ping(); err != nil {
		log.Fatalf("La base de données est injoignable: %v", err)
	}

	log.Println("Base de données connectée avec succès !")
	return db
}

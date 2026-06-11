package config

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	
	driver := GetEnv("DB_DRIVER", "sqlite") 
	source := GetEnv("DB_SOURCE", "./database/inam_forum.db")

	db, err := sql.Open(driver, source)
	if err != nil {
		log.Fatalf("Impossible d'ouvrir la base de données: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("La base de données est injoignable: %v", err)
	}

	log.Println("Base de données connectée avec succès !")

	// Lecture du fichier init.sql créé par Ayrton
	migrationSQL, err := os.ReadFile("./migration/init.sql")
	if err != nil {
		log.Fatalf("Impossible de lire le fichier de migration init.sql: %v", err)
	}

	// Exécution du script SQL dans la BDD
	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation des tables SQL: %v", err)
	}

	log.Println("Tables de la base de données initialisées avec succès !")
	return db
}
package main

import (
	"inam-forum/app"
	"log"
	"net/http"
)

func main() {
	//Initialisation
	application := app.InitApp()
	defer application.Close()

	//Lancement du serveu
	log.Printf("Serveur I Need a Mate lancé sur : http://localhost:8080")
	serveErr := http.ListenAndServe(":8080", application.Router)
	if serveErr != nil {
		log.Fatalf("Erreur lors du lancement du serveur - %s", serveErr.Error())
	}
}

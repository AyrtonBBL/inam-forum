package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"inam-forum/config"

	"github.com/gorilla/mux"
)

type App struct {
	Db     *sql.DB
	Router *mux.Router
}

func InitApp() *App {
	config.LoadEnv()
	db := config.InitDB()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Bienvenue sur I Need a Mate </h1><p>Prêt à trouver ton futur mate ?</p>")
	}).Methods("GET")

	_ = router.PathPrefix("/api").Subrouter()

	/* 
	   Les futurs contrôleurs se brancheront sur apiRouter .
	*/

	return &App{
		Db:     db,
		Router: router, 
	}
}

func (a *App) Close() {
	if a.Db != nil {
		a.Db.Close()
		log.Println("Connexion à la base de données fermée proprement.")
	}
}
package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"inam-forum/config"
	"inam-forum/controllers"
	"inam-forum/repositories"
	"inam-forum/routers"
	"inam-forum/services"

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

	// Page d'accueil de bienvenue
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Bienvenue sur I Need a Mate </h1><p>Prêt à trouver ton futur mate ?</p>")
	}).Methods("GET")

	userRepo := repositories.InitUserRepository(db)
	authService := services.InitAuthService(userRepo)
	authController := controllers.InitAuthController(authService)

	categoryRepo := repositories.InitCategoryRepository(db)
	categoryService := services.InitCategoryService(categoryRepo)
	categoryController := controllers.InitCategoryController(categoryService)

	threadRepo := repositories.InitThreadRepository(db)
	threadService := services.InitThreadService(threadRepo)
	threadController := controllers.InitThreadController(threadService)

	apiRouter := router.PathPrefix("/api").Subrouter()

	routers.RegisterAuthRoutes(apiRouter, authController)
	routers.RegisterCategoryRoutes(apiRouter, categoryController)
	routers.RegisterThreadRoutes(apiRouter, threadController)

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

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

	messageRepo := repositories.InitMessageRepository(db)
	messageService := services.InitMessageService(messageRepo)
	messageController := controllers.InitMessageController(messageService)

	reactionRepo := repositories.InitReactionRepository(db)
	reactionService := services.InitReactionService(reactionRepo)
	reactionController := controllers.InitReactionController(reactionService)

	adminService := services.InitAdminService(userRepo)
	adminController := controllers.InitAdminController(adminService)

	apiRouter := router.PathPrefix("/api").Subrouter()

	routers.RegisterAuthRoutes(apiRouter, authController)
	routers.RegisterCategoryRoutes(apiRouter, categoryController)
	routers.RegisterThreadRoutes(apiRouter, threadController)
	routers.RegisterMessageRoutes(apiRouter, messageController)
	routers.RegisterReactionRoutes(apiRouter, reactionController)
	routers.RegisterAdminRoutes(apiRouter, adminController)

	//  AJOUT SÉCURITÉ CORS POUR LE FRONTEND car sinon ca marche pas
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	return &App{
		Db:     db,
		Router: router,
	}
}

func (a *App) Close() {
	if a.Db != nil {
		a.Db.Close()
		log.Println("Connexion à la base de données fermée.")
	}
}

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
	// 1. Chargement de la configuration et de la base de données
	config.LoadEnv()
	db := config.InitDB()

	// 2. Initialisation du routeur principal
	router := mux.NewRouter()

	// --- LE FIX CORS DÉFINITIF POUR GORILLA MUX ---
	// Autoriser les requêtes de vérification (OPTIONS) du navigateur
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		w.WriteHeader(http.StatusOK)
	})

	// Appliquer les headers CORS à toutes les autres requêtes
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
			next.ServeHTTP(w, r)
		})
	})
	// -----------------------------------------------

	// Page d'accueil de bienvenue pour vérifier que le Go tourne
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Bienvenue sur l'API de I Need a Mate !")
	}).Methods("GET")

	// 3. Initialisation de toutes les couches (Repositories, Services, Controllers)
	
	// Utilisateurs & Auth
	userRepo := repositories.InitUserRepository(db)
	authService := services.InitAuthService(userRepo)
	authController := controllers.InitAuthController(authService)

	// Catégories / Salons
	categoryRepo := repositories.InitCategoryRepository(db)
	categoryService := services.InitCategoryService(categoryRepo)
	categoryController := controllers.InitCategoryController(categoryService)

	// Annonces / Threads
	threadRepo := repositories.InitThreadRepository(db)
	threadService := services.InitThreadService(threadRepo)
	threadController := controllers.InitThreadController(threadService)

	// Messages / Réponses
	messageRepo := repositories.InitMessageRepository(db)
	messageService := services.InitMessageService(messageRepo)
	messageController := controllers.InitMessageController(messageService)

	// Réactions / Likes
	reactionRepo := repositories.InitReactionRepository(db)
	reactionService := services.InitReactionService(reactionRepo)
	reactionController := controllers.InitReactionController(reactionService)

	// Administration / Bannissement
	adminService := services.InitAdminService(userRepo)
	adminController := controllers.InitAdminController(adminService)

	// 4. Création du sous-routeur avec le préfixe /api
	apiRouter := router.PathPrefix("/api").Subrouter()

	// 5. Enregistrement de toutes les routes
	routers.RegisterAuthRoutes(apiRouter, authController)
	routers.RegisterCategoryRoutes(apiRouter, categoryController)
	routers.RegisterThreadRoutes(apiRouter, threadController)
	routers.RegisterMessageRoutes(apiRouter, messageController)
	routers.RegisterReactionRoutes(apiRouter, reactionController)
	routers.RegisterAdminRoutes(apiRouter, adminController)

	return &App{
		Db:     db,
		Router: router,
	}
}

// Run démarre le serveur web
func (a *App) Run() {
	fmt.Println("🚀 Serveur démarré avec succès sur le port 8080...")
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
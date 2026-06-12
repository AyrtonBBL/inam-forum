package routers

import (
	"inam-forum/controllers"
	"inam-forum/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAdminRoutes(router *mux.Router, controller *controllers.AdminController) {
	// Route sécurisée : Seul un admin connecté peut POSTER ici pour bannir
	router.Handle("/admin/ban", middleware.AuthMiddleware(http.HandlerFunc(controller.BanHandler))).Methods("POST")
}

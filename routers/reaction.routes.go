package routers

import (
	"inam-forum/controllers"
	"inam-forum/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterReactionRoutes(router *mux.Router, controller *controllers.ReactionController) {
	// Route sécurisée : POST /api/reactions (FT-6)
	router.Handle("/reactions", middleware.AuthMiddleware(http.HandlerFunc(controller.VoteHandler))).Methods("POST")
}
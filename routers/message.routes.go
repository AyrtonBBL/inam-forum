package routers

import (
	"inam-forum/controllers"
	"inam-forum/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterMessageRoutes(router *mux.Router, controller *controllers.MessageController) {
	// Route sécurisée : POST /api/messages
	router.Handle("/messages", middleware.AuthMiddleware(http.HandlerFunc(controller.CreateHandler))).Methods("POST")
	// Route publique pour voir les messages d'un fil 
	router.HandleFunc("/threads/{id}/messages", controller.GetByThreadHandler).Methods("GET")
}


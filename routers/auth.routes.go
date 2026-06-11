package routers

import (
	"inam-forum/controllers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router, authController *controllers.AuthController) {
	
	router.HandleFunc("/register", authController.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", authController.LoginHandler).Methods("POST")
}

package routers

import (
	"inam-forum/controllers"
	"inam-forum/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterThreadRoutes(router *mux.Router, controller *controllers.ThreadController) {
	// 1 utilisateur avec un JWT valide peut poster
	router.Handle("/threads", middleware.AuthMiddleware(http.HandlerFunc(controller.CreateHandler))).Methods("POST")
}

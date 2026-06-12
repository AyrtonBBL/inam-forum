package routers

import (
	"inam-forum/controllers"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(router *mux.Router, controller *controllers.CategoryController) {
	// URL : /api/categories
	router.HandleFunc("/categories", controller.GetAllHandler).Methods("GET")
}

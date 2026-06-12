package controllers

import (
	"encoding/json"
	"inam-forum/services"
	"net/http"
)

type CategoryController struct {
	service *services.CategoryService
}

func InitCategoryController(service *services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := c.service.GetAllCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erreur lors de la récupération des salons"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

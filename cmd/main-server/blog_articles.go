package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func blogArticlesRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/categories", getArticleCategories)

	r.Get("/status-options", getArticleStatusOptions)

	return r
}

func getArticleCategories(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadArticleCategories(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading article categories", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "article categories read successfully", Data: teams}, http.StatusOK)
}

func getArticleStatusOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadArticleStatusOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading article status options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "article status options read successfully", Data: teams}, http.StatusOK)
}

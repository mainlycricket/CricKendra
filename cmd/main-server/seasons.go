package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func seasonsRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createSeason)
	r.Get("/", getSeasons)
	return r
}

func createSeason(w http.ResponseWriter, r *http.Request) {
	var season models.Season

	err := json.NewDecoder(r.Body).Decode(&season)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	err = dbutils.InsertSeason(r.Context(), DB_POOL, &season)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting season", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "season created successfully", Data: nil}, http.StatusCreated)
}

func getSeasons(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadSeasons(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading seasons", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "seasons read successfully", Data: response}, http.StatusOK)
}

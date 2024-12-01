package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func continentsRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createContinent)
	r.Get("/", getContinents)
	return r
}

func createContinent(w http.ResponseWriter, r *http.Request) {
	var continent models.Continent

	err := json.NewDecoder(r.Body).Decode(&continent)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	continentId, err := dbutils.InsertContinent(r.Context(), DB_POOL, &continent)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting continent", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "continent created successfully", Data: continentId}, http.StatusCreated)
}

func getContinents(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadContinents(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading continents", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "continents read successfully", Data: response}, http.StatusOK)
}

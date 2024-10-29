package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func citiesRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createCity)
	return r
}

func createCity(w http.ResponseWriter, r *http.Request) {
	var city models.City

	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	err = dbutils.InsertCity(r.Context(), DB_POOL, &city)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting city", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "city created successfully", Data: nil}, http.StatusCreated)
}

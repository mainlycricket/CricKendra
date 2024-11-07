package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func toursRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createTour)
	r.Get("/", getTours)
	return r
}

func createTour(w http.ResponseWriter, r *http.Request) {
	var tour models.Tour

	err := json.NewDecoder(r.Body).Decode(&tour)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	err = dbutils.InsertTour(r.Context(), DB_POOL, &tour)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting tour", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "tour created successfully", Data: nil}, http.StatusCreated)
}

func getTours(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadTours(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading tours", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "tours read successfully", Data: response}, http.StatusOK)
}

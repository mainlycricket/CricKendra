package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func seriesRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createSeries)
	return r
}

func createSeries(w http.ResponseWriter, r *http.Request) {
	var series models.Series

	err := json.NewDecoder(r.Body).Decode(&series)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	err = dbutils.InsertSeries(r.Context(), DB_POOL, &series)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting series", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "series created successfully", Data: nil}, http.StatusCreated)
}

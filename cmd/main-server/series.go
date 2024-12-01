package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func seriesRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createSeries)
	r.Get("/", getSeries)

	r.Get("/{seriesId}", getSeriesById)
	return r
}

func createSeries(w http.ResponseWriter, r *http.Request) {
	var series models.Series

	err := json.NewDecoder(r.Body).Decode(&series)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	seriesId, err := dbutils.InsertSeries(r.Context(), DB_POOL, &series)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting series", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "series created successfully", Data: seriesId}, http.StatusCreated)
}

func getSeries(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadSeries(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading series", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "series read successfully", Data: response}, http.StatusOK)
}

func getSeriesById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("seriesId")
	int_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid player id", Data: err}, http.StatusBadRequest)
		return
	}

	players, err := dbutils.ReadSeriesById(r.Context(), DB_POOL, int_id)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading series", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched series successfully", Data: players}, http.StatusOK)
}

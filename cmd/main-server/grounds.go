package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func groundsRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createGround)
	r.Get("/", getGrounds)
	return r
}

func createGround(w http.ResponseWriter, r *http.Request) {
	var ground models.Ground

	err := json.NewDecoder(r.Body).Decode(&ground)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	groundId, err := dbutils.InsertGround(r.Context(), DB_POOL, &ground)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting ground", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "ground created successfully", Data: groundId}, http.StatusCreated)
}

func getGrounds(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadGrounds(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading grounds", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "grounds read successfully", Data: response}, http.StatusOK)
}

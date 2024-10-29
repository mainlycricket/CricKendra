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
	return r
}

func createGround(w http.ResponseWriter, r *http.Request) {
	var ground models.Ground

	err := json.NewDecoder(r.Body).Decode(&ground)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	err = dbutils.InsertGround(r.Context(), DB_POOL, &ground)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting ground", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "ground created successfully", Data: nil}, http.StatusCreated)
}

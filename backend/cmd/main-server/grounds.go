package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/dbutils"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
	"github.com/mainlycricket/CricKendra/backend/internal/responses"
	"github.com/mainlycricket/CricKendra/backend/internal/utils"
)

func groundsRouter() *chi.Mux {
	r := chi.NewRouter()

	// auth by controller
	r.Post("/", createGround)

	r.Get("/", getGrounds)

	return r
}

func createGround(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	var ground models.Ground

	if err := json.NewDecoder(r.Body).Decode(&ground); err != nil {
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

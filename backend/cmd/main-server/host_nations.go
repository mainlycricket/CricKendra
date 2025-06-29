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

func hostNationsRouter() *chi.Mux {
	r := chi.NewRouter()

	// auth by controller
	r.Post("/", createHostNation)

	r.Get("/", getHostNations)

	return r
}

func createHostNation(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	var host_nation models.HostNation

	if err := json.NewDecoder(r.Body).Decode(&host_nation); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	hostNationId, err := dbutils.InsertHostNation(r.Context(), DB_POOL, &host_nation)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting host nation", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "host nation created successfully", Data: hostNationId}, http.StatusCreated)
}

func getHostNations(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadHostNations(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading host nations", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "host nations read successfully", Data: response}, http.StatusOK)
}

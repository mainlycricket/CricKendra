package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func hostNationsRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createHostNation)
	return r
}

func createHostNation(w http.ResponseWriter, r *http.Request) {
	var host_nation models.HostNation

	err := json.NewDecoder(r.Body).Decode(&host_nation)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	err = dbutils.InsertHostNation(r.Context(), DB_POOL, &host_nation)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting host nation", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "host nation created successfully", Data: nil}, http.StatusCreated)
}

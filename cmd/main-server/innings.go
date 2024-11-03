package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func inningsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/innings-end-options", getInningsEndOptions)

	return r
}

func getInningsEndOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadInningsEndOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading innings end options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings end options read successfully", Data: teams}, http.StatusOK)
}

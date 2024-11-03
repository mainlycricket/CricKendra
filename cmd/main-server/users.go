package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func usersRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/user-role-options", getUserRoleOptions)

	return r
}

func getUserRoleOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadUserRoleOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading user role options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "user role options read successfully", Data: teams}, http.StatusOK)
}

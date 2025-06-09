package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/dbutils"
	"github.com/mainlycricket/CricKendra/backend/internal/responses"
)

func StatFiltersRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", Get_Stat_Filter_Options)

	return r
}

func Get_Stat_Filter_Options(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Stat_Filter_Options(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading stat-filter options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched stat-filter options successfully", Data: response}, http.StatusOK)
}

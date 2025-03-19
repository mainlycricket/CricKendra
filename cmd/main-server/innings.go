package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func inningsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{inningsNumber}/commentary", getMatchInningsDeliveries)

	return r
}

func getMatchInningsDeliveries(w http.ResponseWriter, r *http.Request) {
	matchIdRaw, inningsNumberRaw := r.PathValue("matchId"), r.PathValue("inningsNumber")

	matchId, err := strconv.ParseInt(matchIdRaw, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid match id", Data: nil}, http.StatusBadRequest)
		return
	}

	inningsNumber, err := strconv.Atoi(inningsNumberRaw)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "invalid innings number", Data: err}, http.StatusBadRequest)
		return
	}

	commentary, err := dbutils.ReadDeliveriesByMatchInnings(r.Context(), DB_POOL, matchId, inningsNumber)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Message: "error while reading innings commentary", Data: err}, http.StatusInternalServerError)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched innings commentary", Data: commentary}, http.StatusOK)
}

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

/* Used in matchesRouter */

func inningsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", createInnings)
	r.Get("/{inningsNumber}/commentary", getMatchInningsDeliveries)

	return r
}

func createInnings(w http.ResponseWriter, r *http.Request) {
	var innings models.Innings

	err := json.NewDecoder(r.Body).Decode(&innings)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	inningsId, err := dbutils.InsertInnings(r.Context(), DB_POOL, &innings)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting innings", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings created successfully", Data: inningsId}, http.StatusCreated)
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

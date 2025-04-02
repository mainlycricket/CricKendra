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

	r.Patch("/{inningsId}/innings-end", updateInningsEnd)
	r.Patch("/{inningsId}/current-batters", updateInningsCurrentBatters)
	r.Patch("/{inningsId}/current-bowlers", updateInningsCurrentBowlers)

	r.Mount("/{inningsNumber}/batting-scorecards", battingScorecardsRouter())
	r.Mount("/{inningsNumber}/bowling-scorecards", bowlingScorecardsRouter())
	r.Mount("/{inningsNumber}/deliveries", deliveriesRouter())

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

func updateInningsEnd(w http.ResponseWriter, r *http.Request) {
	rawInningsId := r.PathValue("inningsId")
	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.InningsEndInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true

	if err := dbutils.UpdateInningsEnd(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating innings end", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings end updated successfully", Data: nil}, http.StatusCreated)
}

func updateInningsCurrentBatters(w http.ResponseWriter, r *http.Request) {
	rawInningsId := r.PathValue("inningsId")
	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.InningsCurrentBattersInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true

	if err := dbutils.UpdateInningsCurrentBatters(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating innings current batters", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings current batters updated successfully", Data: nil}, http.StatusCreated)
}

func updateInningsCurrentBowlers(w http.ResponseWriter, r *http.Request) {
	rawInningsId := r.PathValue("inningsId")
	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.InningsCurrentBowlersInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true

	if err := dbutils.UpdateInningsCurrentBowlers(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating innings current bowlers", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings current bowlers updated successfully", Data: nil}, http.StatusCreated)
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

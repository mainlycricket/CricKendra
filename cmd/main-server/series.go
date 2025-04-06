package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/models"
	"github.com/mainlycricket/CricKendra/internal/responses"
	"github.com/mainlycricket/CricKendra/internal/utils"
)

func seriesRouter() *chi.Mux {
	r := chi.NewRouter()

	// auth by controller
	r.Post("/", createSeries)
	r.Patch("/{seriesId}/final-result", updateSeriesFinalResult)
	r.Post("/{seriesId}/squads", createSeriesSquad)
	r.Post("/{seriesId}/squads/{squadId}/upsert-entries", upsertSeriesSquadEntries)

	r.Get("/", getSeries)
	r.Get("/{seriesId}", getSeriesOverviewById)
	r.Get("/{seriesId}/matches", getSeriesMatchesById)
	r.Get("/{seriesId}/teams", getSeriesTeamsById)
	r.Get("/{seriesId}/squads-list", getSeriesSquadsListById)
	r.Get("/{seriesId}/squads/{squadId}", getSeriesSingleSquadById)

	return r
}

func createSeries(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	var series models.Series

	if err := json.NewDecoder(r.Body).Decode(&series); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	seriesId, err := dbutils.InsertSeries(r.Context(), DB_POOL, &series)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting series", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "series created successfully", Data: seriesId}, http.StatusCreated)
}

func createSeriesSquad(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	rawSeriesId := r.PathValue("seriesId")
	parsedSeriesId, err := strconv.ParseInt(rawSeriesId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid series id", Data: err}, http.StatusBadRequest)
		return
	}

	var squad models.SeriesSquad

	if err := json.NewDecoder(r.Body).Decode(&squad); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	squad.SeriesId.Int64, squad.SeriesId.Valid = parsedSeriesId, true

	squadId, err := dbutils.InsertSeriesSquad(r.Context(), DB_POOL, &squad)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting series squad", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "series squad created successfully", Data: squadId}, http.StatusCreated)
}

func upsertSeriesSquadEntries(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	rawSquadId := r.PathValue("squadId")
	parsedSquadId, err := strconv.ParseInt(rawSquadId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid squad id", Data: err}, http.StatusBadRequest)
		return
	}

	var entries []models.SeriesSquadEntry

	if err := json.NewDecoder(r.Body).Decode(&entries); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	for idx, entry := range entries {
		entry.SquadId.Int64, entry.SquadId.Valid = parsedSquadId, true
		entries[idx] = entry
	}

	if err := dbutils.UpsertSeriesSquadEntries(r.Context(), DB_POOL, entries); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while upserting series squad entries", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "series squad entries upserted successfully", Data: nil}, http.StatusCreated)
}

func updateSeriesFinalResult(w http.ResponseWriter, r *http.Request) {
	_, err := utils.AuthorizeRequest(r, []string{SYSTEM_ADMIN_ROLE})
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
		return
	}

	rawSeriesId := r.PathValue("seriesId")
	parsedSeriesId, err := strconv.ParseInt(rawSeriesId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid series id", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.SeriesFinalResult

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.SeriesId.Int64, input.SeriesId.Valid = parsedSeriesId, true

	if err := dbutils.UpdateSeriesFinalResult(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while upserting series final result", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "series final result updated successfully", Data: nil}, http.StatusOK)
}

func getSeries(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.ReadSeries(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading series", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "series read successfully", Data: response}, http.StatusOK)
}

func getSeriesOverviewById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("seriesId")
	int_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid series id", Data: err}, http.StatusBadRequest)
		return
	}

	seriesOverview, err := dbutils.ReadSeriesOverviewById(r.Context(), DB_POOL, int_id)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading series overview", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched series overview successfully", Data: seriesOverview}, http.StatusOK)
}

func getSeriesMatchesById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("seriesId")
	int_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid series id", Data: err}, http.StatusBadRequest)
		return
	}

	seriesWithMatches, err := dbutils.ReadSeriesMatchesById(r.Context(), DB_POOL, int_id)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading series with matches", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched series with matches successfully", Data: seriesWithMatches}, http.StatusOK)
}

func getSeriesTeamsById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("seriesId")
	int_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid series id", Data: err}, http.StatusBadRequest)
		return
	}

	seriesWithTeams, err := dbutils.ReadSeriesTeamsById(r.Context(), DB_POOL, int_id)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading series with teams", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched series with teams successfully", Data: seriesWithTeams}, http.StatusOK)
}

func getSeriesSquadsListById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("seriesId")
	int_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid series id", Data: err}, http.StatusBadRequest)
		return
	}

	seriesWithSquadsList, err := dbutils.ReadSeriesSquadsListById(r.Context(), DB_POOL, int_id)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading series with squads list", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched series with squads list successfully", Data: seriesWithSquadsList}, http.StatusOK)
}

func getSeriesSingleSquadById(w http.ResponseWriter, r *http.Request) {
	seriesId, squadId := r.PathValue("seriesId"), r.PathValue("squadId")

	intSeriesId, err := strconv.ParseInt(seriesId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid series id", Data: err}, http.StatusBadRequest)
		return
	}

	intSquadId, err := strconv.ParseInt(squadId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid squad id", Data: err}, http.StatusBadRequest)
		return
	}

	seriesWithSquadsList, err := dbutils.ReadSeriesSingleSquadById(r.Context(), DB_POOL, intSeriesId, intSquadId)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading series single squad", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched series single squad successfully", Data: seriesWithSquadsList}, http.StatusOK)
}

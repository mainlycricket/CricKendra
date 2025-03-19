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

func seriesRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", createSeries)
	r.Get("/", getSeries)

	r.Get("/{seriesId}", getSeriesOverviewById)
	r.Get("/{seriesId}/matches", getSeriesMatchesById)
	r.Get("/{seriesId}/teams", getSeriesTeamsById)
	r.Get("/{seriesId}/squads-list", getSeriesSquadsListById)
	r.Get("/{seriesId}/squads/{squadId}", getSeriesSingleSquadById)
	return r
}

func createSeries(w http.ResponseWriter, r *http.Request) {
	var series models.Series

	err := json.NewDecoder(r.Body).Decode(&series)
	if err != nil {
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

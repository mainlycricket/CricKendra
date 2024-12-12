package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func overallBattingStatsRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/batters", Get_Overall_Batting_Batters_Stats)
	r.Get("/teams", Get_Overall_Batting_Teams_Stats)
	r.Get("/oppositions", Get_Overall_Batting_Oppositions_Stats)
	r.Get("/seasons", Get_Overall_Batting_Seasons_Stats)
	r.Get("/years", Get_Overall_Batting_Years_Stats)
	return r
}

// Function Names are in Get_Overall_Batting_x_Stats format, x represents grouping

func Get_Overall_Batting_Batters_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Batters_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall batters stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall batters stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_Teams_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Teams_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall batting team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall batting team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_Oppositions_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Oppositions_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall bowling team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall bowling team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_Seasons_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Seasons_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall batting season stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall batting season stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_Years_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Years_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall batting years stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall batting years stats successfully", Data: response}, http.StatusOK)
}

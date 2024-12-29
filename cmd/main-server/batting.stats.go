package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func BattingStatsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/overall/batters", Get_Overall_Batting_Batters_Stats)
	r.Get("/overall/team-innings", Get_Overall_Batting_TeamInnings_Stats)
	r.Get("/overall/matches", Get_Overall_Batting_Matches_Stats)
	r.Get("/overall/teams", Get_Overall_Batting_Teams_Stats)
	r.Get("/overall/oppositions", Get_Overall_Batting_Oppositions_Stats)
	r.Get("/overall/grounds", Get_Overall_Batting_Grounds_Stats)
	r.Get("/overall/host-nations", Get_Overall_Batting_HostNations_Stats)
	r.Get("/overall/continents", Get_Overall_Batting_Continents_Stats)
	r.Get("/overall/years", Get_Overall_Batting_Years_Stats)
	r.Get("/overall/seasons", Get_Overall_Batting_Seasons_Stats)
	r.Get("/overall/aggregate", Get_Overall_Batting_Aggregate_Stats)

	r.Get("/individual/innings", Get_Individual_Batting_Innings_Stats)
	r.Get("/individual/grounds", Get_Individual_Batting_Grounds_Stats)
	r.Get("/individual/host-nations", Get_Individual_Batting_HostNations_Stats)
	r.Get("/individual/oppositions", Get_Individual_Batting_Oppositions_Stats)
	r.Get("/individual/years", Get_Individual_Batting_Years_Stats)
	r.Get("/individual/seasons", Get_Individual_Batting_Seasons_Stats)

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

func Get_Overall_Batting_TeamInnings_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_TeamInnings_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team innings stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team innings stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_Matches_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Matches_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall matches stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall matches stats successfully", Data: response}, http.StatusOK)
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

func Get_Overall_Batting_Grounds_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Grounds_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall ground stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall ground stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_HostNations_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_HostNations_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall host nations stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall host nations stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_Continents_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Continents_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall continents stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall continents stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_Years_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Years_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall batting years stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall batting years stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_Seasons_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Seasons_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall batting season stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall batting season stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Batting_Aggregate_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Batting_Aggregate_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall batting aggregate stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall batting aggregate stats successfully", Data: response}, http.StatusOK)
}

// Function Names are in Get_Individual_Batting_x_Stats format, x represents grouping

func Get_Individual_Batting_Innings_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Batting_Innings_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual batters-innings stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual batters-innings stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Batting_Grounds_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Batting_Grounds_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual batters-grounds stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual batters-grounds stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Batting_HostNations_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Batting_HostNations_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual batters-host_nations stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual batters-host_nations stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Batting_Oppositions_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Batting_Oppositions_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual batters-oppositions stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual batters-oppositions stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Batting_Years_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Batting_Years_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual batters-years stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual batters-years stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Batting_Seasons_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Batting_Seasons_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual batters-seasons stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual batters-seasons stats successfully", Data: response}, http.StatusOK)
}

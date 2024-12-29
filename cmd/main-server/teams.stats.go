package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func TeamStatsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/overall/teams", Get_Overall_Team_Teams_Stats)
	r.Get("/overall/players", Get_Overall_Team_Players_Stats)
	r.Get("/overall/matches", Get_Overall_Team_Matches_Stats)
	r.Get("/overall/grounds", Get_Overall_Team_Grounds_Stats)
	r.Get("/overall/host-nations", Get_Overall_Team_HostNations_Stats)
	r.Get("/overall/continents", Get_Overall_Team_Continents_Stats)
	r.Get("/overall/years", Get_Overall_Team_Years_Stats)
	r.Get("/overall/seasons", Get_Overall_Team_Seasons_Stats)
	r.Get("/overall/decades", Get_Overall_Team_Decades_Stats)
	r.Get("/overall/aggregate", Get_Overall_Team_Aggregate_Stats)

	r.Get("/individual/matches", Get_Individual_Team_Matches_Stats)
	r.Get("/individual/innings", Get_Individual_Team_Innings_Stats)
	r.Get("/individual/grounds", Get_Individual_Team_Grounds_Stats)
	r.Get("/individual/host-nations", Get_Individual_Team_HostNations_Stats)
	r.Get("/individual/years", Get_Individual_Team_Years_Stats)
	r.Get("/individual/seasons", Get_Individual_Team_Seasons_Stats)

	return r
}

// Function Names are in Get_Overall_Team_x_Stats format, x represents grouping

func Get_Overall_Team_Teams_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_Teams_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Team_Players_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_Players_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Team_Matches_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_Matches_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Team_Grounds_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_Grounds_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Team_HostNations_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_HostNations_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Team_Continents_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_Continents_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Team_Years_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_Years_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Team_Seasons_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_Seasons_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Team_Decades_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_Decades_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Team_Aggregate_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Team_Aggregate_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team stats successfully", Data: response}, http.StatusOK)
}

// Function Names are in Get_Individual_Team_x_Stats format, x represents grouping

func Get_Individual_Team_Matches_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Team_Matches_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual team-matches stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual team-matches stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Team_Innings_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Team_Innings_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual team-innings stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual team-innings stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Team_Grounds_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Team_Grounds_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual team-grounds stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual team-grounds stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Team_HostNations_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Team_HostNations_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual team-host_nations stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual team-host_nations stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Team_Years_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Team_Years_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual team-years stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual team-years stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Team_Seasons_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Team_Seasons_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading individual team-seasons stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched individual team-seasons stats successfully", Data: response}, http.StatusOK)
}

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func BowlingStatsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/overall/summary", Get_Overall_Bowling_Summary_Stats)
	r.Get("/overall/bowlers", Get_Overall_Bowling_Bowlers_Stats)
	r.Get("/overall/team-innings", Get_Overall_Bowling_TeamInnings_Stats)
	r.Get("/overall/matches", Get_Overall_Bowling_Matches_Stats)
	r.Get("/overall/teams", Get_Overall_Bowling_Teams_Stats)
	r.Get("/overall/oppositions", Get_Overall_Bowling_Oppositions_Stats)
	r.Get("/overall/grounds", Get_Overall_Bowling_Grounds_Stats)
	r.Get("/overall/host-nations", Get_Overall_Bowling_HostNations_Stats)
	r.Get("/overall/continents", Get_Overall_Bowling_Continents_Stats)
	r.Get("/overall/series", Get_Overall_Bowling_Series_Stats)
	r.Get("/overall/tournaments", Get_Overall_Bowling_Tournaments_Stats)
	r.Get("/overall/years", Get_Overall_Bowling_Years_Stats)
	r.Get("/overall/seasons", Get_Overall_Bowling_Seasons_Stats)
	r.Get("/overall/decades", Get_Overall_Bowling_Decades_Stats)
	r.Get("/overall/aggregate", Get_Overall_Bowling_Aggregate_Stats)

	r.Get("/individual/innings", Get_Individual_Bowling_Innings_Stats)
	r.Get("/individual/match-totals", Get_Individual_Bowling_MatchTotals_Stats)
	r.Get("/individual/series", Get_Individual_Bowling_Series_Stats)
	r.Get("/individual/tournaments", Get_Individual_Bowling_Tournaments_Stats)
	r.Get("/individual/grounds", Get_Individual_Bowling_Grounds_Stats)
	r.Get("/individual/host-nations", Get_Individual_Bowling_HostNations_Stats)
	r.Get("/individual/oppositions", Get_Individual_Bowling_Oppositions_Stats)
	r.Get("/individual/years", Get_Individual_Bowling_Years_Stats)
	r.Get("/individual/seasons", Get_Individual_Bowling_Seasons_Stats)

	return r
}

// Function Names are in Get_Overall_Bowling_x_Stats format, x represents grouping

func Get_Overall_Bowling_Summary_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Summary_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall bowling summary stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall bowling summary stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Bowlers_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Bowlers_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall bowlers stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall bowlers stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_TeamInnings_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_TeamInnings_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall team innings stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall team innings stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Matches_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Matches_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall matches stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall matches stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Teams_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Teams_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall teams stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall teams stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Oppositions_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Oppositions_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall oppositions stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall oppositions stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Grounds_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Grounds_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall grounds stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall grounds stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_HostNations_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_HostNations_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall host nations stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall host nations stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Continents_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Continents_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall continents stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall continents stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Series_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Series_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall series stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall series stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Tournaments_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Tournaments_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall tournaments stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall tournaments stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Years_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Years_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall years stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall years stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Seasons_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Seasons_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall seasons stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall seasons stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Decades_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Decades_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall decades stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall decades stats successfully", Data: response}, http.StatusOK)
}

func Get_Overall_Bowling_Aggregate_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Overall_Bowling_Aggregate_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading overall aggregate stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched overall aggregate stats successfully", Data: response}, http.StatusOK)
}

// Function Names are in Get_Individual_Bowling_x_Stats format, x represents grouping

func Get_Individual_Bowling_Innings_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Bowling_Innings_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowler-innings stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched bowler-innings stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Bowling_MatchTotals_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Bowling_MatchTotals_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowler match totals stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched bowler match totals stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Bowling_Series_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Bowling_Series_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowler-series stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched bowler-series stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Bowling_Tournaments_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Bowling_Tournaments_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowler-tournaments stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched bowler-tournaments stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Bowling_Grounds_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Bowling_Grounds_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowler-grounds stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched bowler-grounds stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Bowling_HostNations_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Bowling_HostNations_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowler host nations stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched bowler host nations stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Bowling_Oppositions_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Bowling_Oppositions_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowler-oppositions stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched bowler-oppositions stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Bowling_Years_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Bowling_Years_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowler-years stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched bowler-years stats successfully", Data: response}, http.StatusOK)
}

func Get_Individual_Bowling_Seasons_Stats(w http.ResponseWriter, r *http.Request) {
	response, err := dbutils.Read_Individual_Bowling_Seasons_Stats(r.Context(), DB_POOL, r.URL.Query())

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowler-seasons stats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "fetched bowler-seasons stats successfully", Data: response}, http.StatusOK)
}

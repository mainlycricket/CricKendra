package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/internal/dbutils"
	"github.com/mainlycricket/CricKendra/internal/responses"
)

func optionsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/blogs/articles/categories", getArticleCategoryOptions)
	r.Get("/blogs/articles/status", getArticleStatusOptions)

	r.Get("/innings/end", getInningsEndOptions)

	r.Get("/matches/results", getMatchResultOptions)
	r.Get("/matches/types", getMatchTypeOptions)
	r.Get("/matches/formats", getMatchFormats)
	r.Get("/matches/levels", getMatchLevels)

	r.Get("/players/bowling-styles", getBowlingStyleOptions)
	r.Get("/players/dismissal-types", getDismissalTypeOptions)
	r.Get("/players/playing-status", getPlayingStatusOptions)

	r.Get("/users/roles", getUserRoleOptions)

	return r
}

/* Blog Articles */

func getArticleCategoryOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadArticleCategories(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading article categories", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "article categories read successfully", Data: teams}, http.StatusOK)
}

func getArticleStatusOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadArticleStatusOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading article status options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "article status options read successfully", Data: teams}, http.StatusOK)
}

/* Innings */

func getInningsEndOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadInningsEndOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading innings end options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "innings end options read successfully", Data: teams}, http.StatusOK)
}

/* Matches */

func getMatchResultOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadMatchResultOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading match result options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match result options read successfully", Data: teams}, http.StatusOK)
}

func getMatchTypeOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadMatchTypeOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading match type options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match type options read successfully", Data: teams}, http.StatusOK)
}

func getMatchFormats(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadMatchFormats(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading match formats", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match formats read successfully", Data: teams}, http.StatusOK)
}

func getMatchLevels(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadMatchLevels(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading match levels", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "match levels read successfully", Data: teams}, http.StatusOK)
}

/* Players */

func getBowlingStyleOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadBowlingStyleOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading bowling style options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "bowling style options read successfully", Data: teams}, http.StatusOK)
}

func getDismissalTypeOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadDismissalTypeOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading dismissal type options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "dismissal type options read successfully", Data: teams}, http.StatusOK)
}

func getPlayingStatusOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadPlayingStatusOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading playing status options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "playing status options read successfully", Data: teams}, http.StatusOK)
}

/* Users */

func getUserRoleOptions(w http.ResponseWriter, r *http.Request) {
	teams, err := dbutils.ReadUserRoleOptions(r.Context(), DB_POOL)

	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while reading user role options", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "user role options read successfully", Data: teams}, http.StatusOK)
}

package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mainlycricket/CricKendra/backend/internal/dbutils"
	"github.com/mainlycricket/CricKendra/backend/internal/models"
	"github.com/mainlycricket/CricKendra/backend/internal/responses"
	"github.com/mainlycricket/CricKendra/backend/internal/utils"
)

func usersRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", loginUser)
	r.Post("/logout", LogoutUser)

	return r
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var credentials models.LoginInput

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	loginOutput, err := dbutils.LoginUser(r.Context(), DB_POOL, credentials.Email.String)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while logging in", Data: nil}, http.StatusBadRequest)
		return
	}

	if err := utils.ComparePassword(credentials.Password.String, loginOutput.HashedPassword); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid password", Data: nil}, http.StatusUnauthorized)
		return
	}

	token, err := utils.GetSignedToken(loginOutput.UserId, loginOutput.Role)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while signing token", Data: err}, http.StatusInternalServerError)
		return
	}

	cookie := utils.NewTokenCookie(token)
	http.SetCookie(w, cookie)

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "logged in successfully", Data: nil}, http.StatusOK)
}

func LogoutUser(w http.ResponseWriter, _ *http.Request) {
	cookie := &http.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: time.Now(),
		Path:    "/",
	}

	http.SetCookie(w, cookie)

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "logged out successfully", Data: nil}, http.StatusOK)
}

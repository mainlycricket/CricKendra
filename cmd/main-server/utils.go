package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func writeJsonResponse(w http.ResponseWriter, _ *http.Request, response ApiResponse, status int) {
	if err, ok := response.Data.(error); ok {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			response.Message += fmt.Sprintf(`: %v`, pgErr.Message)
		} else {
			response.Message += fmt.Sprintf(`: %v`, err)
		}

		response.Data = nil
		log.Println(err)
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		response = ApiResponse{false, "failed to jsonify response", nil}
		jsonData, _ = json.Marshal(response)
		status = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}

package responses

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

func WriteJsonResponse(w http.ResponseWriter, response ApiResponse, status int) {
	if err, ok := response.Data.(error); ok {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			var pgMessage string
			status, pgMessage = handlePgErr(pgErr)
			response.Message += fmt.Sprintf(`: %v`, pgMessage)
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

func handlePgErr(pgErr *pgconn.PgError) (int, string) {
	var status int
	var message string

	sqlErrCode := pgErr.SQLState()

	switch sqlErrCode[:2] {
	case "42":
		status = http.StatusInternalServerError
		message = pgErr.Message
	case "23":
		status = http.StatusBadRequest
		message = pgErr.Message
	default:
		status = http.StatusBadRequest
		message = pgErr.Message
	}

	return status, message
}

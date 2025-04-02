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

/* Used in inningsRouter */

func deliveriesRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", createDeliveryWithScoringInput)
	r.Patch("/{inningsDeliveryNumber}/commentary", updateDeliveryCommentary)
	r.Patch("/{inningsDeliveryNumber}/advance-info", updateDeliveryAdvanceInfo)

	return r
}

func createDeliveryWithScoringInput(w http.ResponseWriter, r *http.Request) {
	rawMatchId, rawInningsId := r.PathValue("matchId"), r.PathValue("inningsId")

	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	parsedMatchId, err := strconv.ParseInt(rawMatchId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid match id", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.DeliveryScoringInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true
	input.MatchId.Int64, input.MatchId.Valid = parsedMatchId, true

	if err := dbutils.InsertDeliveryWithScoringData(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while inserting delivery with scoring input", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "delivery with scoring input created successfully", Data: nil}, http.StatusCreated)
}

func updateDeliveryCommentary(w http.ResponseWriter, r *http.Request) {
	rawInningsId := r.PathValue("inningsId")
	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	inningsDeliveryNumber := r.PathValue("inningsDeliveryNumber")
	parsedInningsDeliveryNumber, err := strconv.ParseInt(inningsDeliveryNumber, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings delivery number", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.DeliveryCommentaryInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true
	input.InningsDeliveryNumber.Int64, input.InningsDeliveryNumber.Valid = parsedInningsDeliveryNumber, true

	if err := dbutils.UpdateDeliveryCommentary(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating delivery commentary", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "delivery commentary updated successfully", Data: nil}, http.StatusCreated)
}

func updateDeliveryAdvanceInfo(w http.ResponseWriter, r *http.Request) {
	rawInningsId, inningsDeliveryNumber := r.PathValue("inningsId"), r.PathValue("inningsDeliveryNumber")

	parsedInningsId, err := strconv.ParseInt(rawInningsId, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings id", Data: err}, http.StatusBadRequest)
		return
	}

	parsedInningsDeliveryNumber, err := strconv.ParseInt(inningsDeliveryNumber, 10, 64)
	if err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "invalid innings delivery number", Data: err}, http.StatusBadRequest)
		return
	}

	var input models.DeliveryAdvanceInfoInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while decoding json", Data: err}, http.StatusBadRequest)
		return
	}

	input.InningsId.Int64, input.InningsId.Valid = parsedInningsId, true
	input.InningsDeliveryNumber.Int64, input.InningsDeliveryNumber.Valid = parsedInningsDeliveryNumber, true

	if err := dbutils.UpdateDeliveryAdvanceInfo(r.Context(), DB_POOL, &input); err != nil {
		responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "error while updating delivery advance info", Data: err}, http.StatusBadRequest)
		return
	}

	responses.WriteJsonResponse(w, responses.ApiResponse{Success: true, Message: "delivery advance info updated successfully", Data: nil}, http.StatusCreated)
}

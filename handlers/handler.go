package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danyeric123/rewards-api/db"
	"github.com/danyeric123/rewards-api/domain"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	rewardsDB *db.ReceiptDB
}

func NewHandler(rewardsDB *db.ReceiptDB) *Handler {
	return &Handler{rewardsDB: rewardsDB}
}

// HealthCheck is a handler for health checks
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	logrus.Println("Health check")
	w.WriteHeader(http.StatusOK)
}

// ProcessReceipt is a handler for processing receipts
func (h *Handler) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	logrus.Println("Processing receipt")
	var receipt domain.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		logrus.Error("Failed to decode request body: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: "Invalid request body"})
		return
	}
	points, err := receipt.CalculatePoints()

	if err != nil {
		logrus.Error("Failed to calculate points for receipt: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: "Failed to calculate points for receipt"})
		return
	}

	// Save the receipt to the database
	ID, err := h.rewardsDB.SaveReceipt(receipt, points)

	if err != nil {
		logrus.Error("Failed to save receipt to the database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domain.ProcessResponse{ID: ID})
}

// GetPoints is a handler for getting points
func (h *Handler) GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	// Validate that the ID is a valid UUID
	if _, err := uuid.Parse(ID); err != nil {
		logrus.Error("Invalid UUID: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: "Invalid UUID"})
		return
	}
	logrus.Println("Getting points for ID: ", vars["id"])
	points, err := h.rewardsDB.GetPoints(ID)
	if err != nil {
		logrus.Error("Failed to get points: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(domain.ErrorResponse{Message: "Failed to get points"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domain.GetResponse{Points: points})
}

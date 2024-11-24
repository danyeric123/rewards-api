package main

import (
	"net/http"
	"os"

	"github.com/danyeric123/rewards-api/db"
	"github.com/danyeric123/rewards-api/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	db.Connect()

	receiptDB := db.NewReceiptDB(db.DB)

	h := handlers.NewHandler(receiptDB)

	r := mux.NewRouter()
	r.HandleFunc("/healthz", h.HealthCheck).Methods("GET")
	r.HandleFunc("/receipts/process", h.ProcessReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", h.GetPoints).Methods("GET")
	logrus.Info("Starting server on :8080")
	logrus.Fatal(http.ListenAndServe(":8080", r))
}

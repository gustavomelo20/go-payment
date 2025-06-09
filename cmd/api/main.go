package main

import (
	"encoding/json"
	"log"
	"net/http"

	"go-payment/internal/payment"
	"go-payment/internal/payment/model"
)

func main() {
	service := payment.NewService()

	http.HandleFunc("/charge", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var chargeRequest model.ChargeRequest
		if err := json.NewDecoder(r.Body).Decode(&chargeRequest); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		results := service.ProcessPayments(chargeRequest)

		response := model.ChargeResponse{
			ChargeID: chargeRequest.ChargeID,
			Results:  results,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Println("API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

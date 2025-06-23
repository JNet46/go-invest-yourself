package handlers

import (
	"encoding/json"
	"go-invest-yourself/db"
	"net/http"
)

type PaymentRequest struct {
	SenderID   string  `json:"sender_id"`
	ReceiverID string  `json:"receiver_id"`
	Amount     float64 `json:"amount"`
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	json.NewDecoder(r.Body).Decode(&req)

	sender := db.GetUser(req.SenderID)
	receiver := db.GetUser(req.ReceiverID)

	if sender == nil || receiver == nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	if sender.Balance < req.Amount {
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	sender.Balance -= req.Amount
	receiver.Balance += req.Amount

	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

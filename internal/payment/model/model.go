package model

import "time"

type ChargeRequest struct {
	ChargeID string    `json:"charge_id"`
	Payments []Payment `json:"payments"`
}

type Payment struct {
	Method  PaymentMethod     `json:"method"`
	Amount  float64           `json:"amount"`
	Details map[string]string `json:"details"`
}

type PaymentResult struct {
	PaymentID   string            `json:"payment_id"`
	Status      PaymentStatus     `json:"status"`
	Message     string            `json:"message"`
	ProcessedAt time.Time         `json:"processed_at"`
	Details     map[string]string `json:"details"`
}

type ChargeResponse struct {
	ChargeID string          `json:"charge_id"`
	Results  []PaymentResult `json:"results"`
}

package processor

import (
	"fmt"
	"time"

	"go-payment/internal/payment/model"

	"github.com/google/uuid"
)

type CreditCardProcessor struct{}

func NewCreditCardProcessor() *CreditCardProcessor {
	return &CreditCardProcessor{}
}

func (c *CreditCardProcessor) Process(payment model.Payment) model.PaymentResult {
	result := model.PaymentResult{
		PaymentID:   uuid.NewString(),
		ProcessedAt: time.Now(),
		Details:     make(map[string]string),
	}

	cardNumber, okNum := payment.Details["card_number"]
	expiry, okExp := payment.Details["expiry"]
	cvv, okCvv := payment.Details["cvv"]

	if !okNum || cardNumber == "" || !okExp || expiry == "" || !okCvv || cvv == "" {
		result.Status = model.StatusDeclined
		result.Message = "Incomplete credit card details"
		return result
	}

	result.Status = model.StatusApproved
	result.Message = "Credit card payment approved successfully"
	if len(cardNumber) >= 4 {
		result.Details["card_last4"] = cardNumber[len(cardNumber)-4:]
	}
	result.Details["authorization_number"] = fmt.Sprintf("CC-%d", time.Now().Unix())

	return result
}

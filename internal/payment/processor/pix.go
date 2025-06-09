package processor

import (
	"fmt"
	"time"

	"go-payment/internal/payment/model"

	"github.com/google/uuid"
)

type PixProcessor struct{}

func NewPixProcessor() *PixProcessor {
	return &PixProcessor{}
}

func (p *PixProcessor) Process(payment model.Payment) model.PaymentResult {
	result := model.PaymentResult{
		PaymentID:   uuid.NewString(),
		ProcessedAt: time.Now(),
		Details:     make(map[string]string),
	}

	key, ok := payment.Details["pix_key"]
	if !ok || key == "" {
		result.Status = model.StatusDeclined
		result.Message = "Pix key missing or invalid"
		return result
	}

	result.Status = model.StatusApproved
	result.Message = "Pix payment approved successfully"
	result.Details["pix_key"] = key
	result.Details["receipt_number"] = fmt.Sprintf("PIX-%d", time.Now().Unix())

	return result
}

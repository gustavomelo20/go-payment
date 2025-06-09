package processor

import "go-payment/internal/payment/model"

type Processor interface {
	Process(payment model.Payment) model.PaymentResult
}

package payment

import (
	"go-payment/internal/payment/model"
	"go-payment/internal/payment/processor"
	"sync"
	"time"
)

type Service struct {
	processors map[model.PaymentMethod]processor.Processor
}

func NewService() *Service {
	return &Service{
		processors: map[model.PaymentMethod]processor.Processor{
			model.MethodPix:        processor.NewPixProcessor(),
			model.MethodCreditCard: processor.NewCreditCardProcessor(),
		},
	}
}

func (s *Service) ProcessPayments(charge model.ChargeRequest) []model.PaymentResult {
	var wg sync.WaitGroup
	resultsChan := make(chan model.PaymentResult, len(charge.Payments))

	for _, payment := range charge.Payments {
		wg.Add(1)
		go func(p model.Payment) {
			defer wg.Done()
			proc, ok := s.processors[p.Method]
			if !ok {
				resultsChan <- model.PaymentResult{
					PaymentID:   "",
					Status:      model.StatusDeclined,
					Message:     "Payment method not supported",
					ProcessedAt: time.Time{},
					Details:     nil,
				}
				return
			}

			result := proc.Process(p)
			resultsChan <- result
		}(payment)
	}

	wg.Wait()
	close(resultsChan)

	var results []model.PaymentResult
	for res := range resultsChan {
		results = append(results, res)
	}

	return results
}

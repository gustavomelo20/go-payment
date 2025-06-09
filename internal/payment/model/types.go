package model

type PaymentMethod string

const (
	MethodPix        PaymentMethod = "PIX"
	MethodCreditCard PaymentMethod = "CREDIT_CARD"
)

type PaymentStatus string

const (
	StatusApproved   PaymentStatus = "APPROVED"
	StatusDeclined   PaymentStatus = "DECLINED"
	StatusPending    PaymentStatus = "PENDING"
	StatusProcessing PaymentStatus = "PROCESSING"
)

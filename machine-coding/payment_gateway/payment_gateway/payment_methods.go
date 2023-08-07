package payment_gateway

type PaymentMethod string

const (
	UPI        PaymentMethod = "UPI"
	CreditCard PaymentMethod = "CreditCard"
	DebitCard  PaymentMethod = "DebitCard"
	NetBanking PaymentMethod = "NetBanking"
)

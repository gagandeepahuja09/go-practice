package payment_gateway

type PaymentMethod string

const (
	UPI        PaymentMethod = "UPI"
	Card       PaymentMethod = "Card"
	NetBanking PaymentMethod = "NetBanking"
)

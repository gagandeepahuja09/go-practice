package bank

type PaymentDetails struct {
	method string
}

type Bank interface {
	MakePayment(pd PaymentDetails) error
}

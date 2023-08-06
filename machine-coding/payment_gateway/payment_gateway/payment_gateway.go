package payment_gateway

import "go-practice.com/machine-coding/payment_gateway/bank"

type PaymentGateway struct {
	methodsSupported []string
	banksSupported   []bank.Bank
	// routing logic
	// first filter on supported banks as per the method
	// after that have percentage allocation
}

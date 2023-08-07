package payment_gateway

import (
	"errors"
	"log"
)

var pg PaymentGateway

const (
	ErrorBankPercentageNotEq100 = "total percentage across all banks should exactly be equal to 100"
)

type BankPercentage struct {
	Percentage int
	BankName   Bank
}

type PaymentGateway struct {
	clientMethodsSupported map[*Client][]PaymentMethod
	methodBankDistribution map[PaymentMethod][]BankPercentage
}

func init() {
	pg = PaymentGateway{
		clientMethodsSupported: map[*Client][]PaymentMethod{},
		methodBankDistribution: map[PaymentMethod][]BankPercentage{},
	}
}

type PaymentOptions struct {
	method        PaymentMethod
	methodDetails interface{}
}

func (p *PaymentGateway) MakePayment(c Client, po PaymentOptions) error {
	return nil
}

func SetMethodBankDistribution(method PaymentMethod, bankPercentages []BankPercentage) error {
	totalPercentage := 0
	for _, bp := range bankPercentages {
		totalPercentage += bp.Percentage
	}
	if totalPercentage != 100 {
		return errors.New(ErrorBankPercentageNotEq100)
	}
	pg.methodBankDistribution[method] = bankPercentages
	return nil
}

func ShowRouterDistribution() {
	log.Println("ROUTER_DISTRIBUTION: ", pg.methodBankDistribution)
}

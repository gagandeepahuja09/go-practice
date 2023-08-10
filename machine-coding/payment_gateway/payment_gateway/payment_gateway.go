package payment_gateway

import (
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/exp/slices"
)

var (
	pg    PaymentGateway
	Icici iciciBank
	Hdfc  hdfcBank
)

const (
	ErrorBankPercentageNotEq100 = "total percentage across all banks should exactly be equal to 100"
	ErrorMethodNotSupported     = "following payment method is not supported"
	ErrorNoBankFound            = "no bank found to support the following request"
)

type BankPercentage struct {
	Percentage int
	Bank       Bank
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
	Icici = iciciBank{}
	Hdfc = hdfcBank{}
}

type NetBankingDetails struct {
	Username string
	Password string
}

type CardDetails struct {
	CardNo      string
	Cvv         int
	ExpiryMonth int
	ExpiryYear  int
}

type UpiDetails struct {
	Vpa string
}

type PaymentOptions struct {
	Amount        int
	Method        PaymentMethod
	MethodDetails interface{}
}

func getAppropriateBankFromDistribution(bankPercentages []BankPercentage) (Bank, error) {
	// ideal to use a math/rand library
	rand100 := time.Now().Unix() % 100
	fmt.Println("RAND_100: ", rand100)
	totalPercentage := 0
	for _, bp := range bankPercentages {
		totalPercentage += bp.Percentage
		if totalPercentage > int(rand100) {
			return bp.Bank, nil
		}
	}
	return nil, errors.New(ErrorNoBankFound)
}

func SetMethodBankDistribution(method PaymentMethod, bankPercentages []BankPercentage) error {
	totalPercentage := 0
	for _, bp := range bankPercentages {
		totalPercentage += bp.Percentage
		pg.methodBankDistribution[method] = append(pg.methodBankDistribution[method], BankPercentage{
			Percentage: totalPercentage,
			Bank:       bp.Bank,
		})
	}
	if totalPercentage != 100 {
		pg.methodBankDistribution[method] = nil
		return errors.New(ErrorBankPercentageNotEq100)
	}
	return nil
}

func ShowRouterDistribution() {
	log.Println("ROUTER_DISTRIBUTION: ", pg.methodBankDistribution)
}

func MakePayment(c *Client, po PaymentOptions) error {
	if !HasClient(c) {
		return errors.New(ErrClientNotFound)
	}
	if !slices.Contains(pg.clientMethodsSupported[c], po.Method) {
		return errors.New(ErrorMethodNotSupported)
	}
	bank, err := getAppropriateBankFromDistribution(pg.methodBankDistribution[po.Method])
	fmt.Printf("BANK_FOUND_AND_ERROR: %+v %v\n", bank, err)
	if bank == &Icici {
		fmt.Println("ICICI_BANK")
	} else {
		fmt.Println("HDFC_BANK")
	}

	return getPaymentResponseFromBank(bank, po.Amount, po.Method, po.MethodDetails)
}

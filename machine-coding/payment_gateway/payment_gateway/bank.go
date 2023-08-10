package payment_gateway

import (
	"errors"
	"fmt"
)

type BankName string

const (
	HDFC  BankName = "HDFC"
	ICICI BankName = "ICICI"
	SBI   BankName = "SBI"
)

const (
	ErrorBankGatewayDown      = "bank gateway is currently down for this method. please try alternative method or bank"
	ErrorInvalidMethodDetails = "invalid method details for method: %s"
)

type Bank interface {
	handleNbPayment(amount int, nb NetBankingDetails) error
	handleUpiPayment(amount int, upi UpiDetails) error
	handleCardPayment(amount int, card CardDetails) error
}

type hdfcBank struct {
}
type iciciBank struct {
}

func (h *hdfcBank) handleNbPayment(amount int, nb NetBankingDetails) error {
	return nil
}
func (h *hdfcBank) handleUpiPayment(amount int, upi UpiDetails) error {
	return nil
}
func (h *hdfcBank) handleCardPayment(amount int, card CardDetails) error {
	return nil
}

func (ic *iciciBank) handleNbPayment(amount int, nb NetBankingDetails) error {
	return nil
}
func (ic *iciciBank) handleUpiPayment(amount int, upi UpiDetails) error {
	return errors.New(ErrorBankGatewayDown)
}
func (ic *iciciBank) handleCardPayment(amount int, card CardDetails) error {
	return errors.New(ErrorBankGatewayDown)
}

func handleNbPayment(bank Bank, amount int, details interface{}) error {
	if nbDetails, ok := details.(NetBankingDetails); ok {
		return bank.handleNbPayment(amount, nbDetails)
	}
	return fmt.Errorf(ErrorInvalidMethodDetails, NetBanking)
}
func handleUpiPayment(bank Bank, amount int, details interface{}) error {
	if upiDetails, ok := details.(UpiDetails); ok {
		return bank.handleUpiPayment(amount, upiDetails)
	}
	return fmt.Errorf(ErrorInvalidMethodDetails, UPI)
}
func handleCardPayment(bank Bank, amount int, details interface{}) error {
	if cardDetails, ok := details.(CardDetails); ok {
		return bank.handleCardPayment(amount, cardDetails)
	}
	return fmt.Errorf(ErrorInvalidMethodDetails, Card)
}

func getPaymentResponseFromBank(bank Bank, amount int, methodName PaymentMethod, methodDetails interface{}) error {
	switch methodName {
	case Card:
		return handleCardPayment(bank, amount, methodDetails)
	case UPI:
		return handleUpiPayment(bank, amount, methodDetails)
	case NetBanking:
		return handleNbPayment(bank, amount, methodDetails)
	default:
		return errors.New(ErrorMethodNotSupported)
	}
}

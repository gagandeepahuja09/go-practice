package main

import (
	"fmt"

	"go-practice.com/machine-coding/payment_gateway/payment_gateway"
)

func setBankDistributionHelper() {
	errBankD := payment_gateway.SetMethodBankDistribution(payment_gateway.UPI, []payment_gateway.BankPercentage{
		{
			Percentage: 50,
			Bank:       &payment_gateway.Hdfc,
		},
		{
			Percentage: 60,
			Bank:       &payment_gateway.Icici,
		},
	})
	fmt.Printf("errBankD: %v\n", errBankD)

	errBankD = payment_gateway.SetMethodBankDistribution(payment_gateway.UPI, []payment_gateway.BankPercentage{
		{
			Percentage: 80,
			Bank:       &payment_gateway.Hdfc,
		},
		{
			Percentage: 20,
			Bank:       &payment_gateway.Icici,
		},
	})
	fmt.Printf("errBankD: %v\n", errBankD)

	errBankD = payment_gateway.SetMethodBankDistribution(payment_gateway.Card, []payment_gateway.BankPercentage{
		{
			Percentage: 10,
			Bank:       &payment_gateway.Hdfc,
		},
		{
			Percentage: 90,
			Bank:       &payment_gateway.Icici,
		},
	})
	fmt.Printf("errBankD: %v\n", errBankD)

	payment_gateway.ShowRouterDistribution()
}

func main() {
	cFk := payment_gateway.AddClient("fk")
	fmt.Println("HasClient fk: ", payment_gateway.HasClient(cFk))

	payment_gateway.RemoveClient(cFk)
	fmt.Println("HasClient fk: ", payment_gateway.HasClient(cFk))

	cZomato := payment_gateway.AddClient("Zomato")

	errFk := payment_gateway.AddSupportForMethods(cFk, []payment_gateway.PaymentMethod{
		payment_gateway.Card,
	})
	fmt.Printf("errFk: %v\n", errFk)

	errZomato := payment_gateway.AddSupportForMethods(cZomato, []payment_gateway.PaymentMethod{
		payment_gateway.Card,
		payment_gateway.UPI,
	})
	fmt.Printf("errZomato: %v\n", errZomato)

	payment_gateway.ListSupportedMethods(cZomato)

	setBankDistributionHelper()

	err := payment_gateway.MakePayment(cZomato, payment_gateway.PaymentOptions{
		Method:        payment_gateway.Card,
		MethodDetails: payment_gateway.CardDetails{},
	})
	fmt.Printf("MakePayment_err: %v\n", err)

	err = payment_gateway.MakePayment(cZomato, payment_gateway.PaymentOptions{
		Method:        payment_gateway.NetBanking,
		MethodDetails: payment_gateway.NetBankingDetails{},
	})
	fmt.Printf("MakePayment_err_NB: %v\n", err)

	err = payment_gateway.MakePayment(cZomato, payment_gateway.PaymentOptions{
		Method:        payment_gateway.UPI,
		MethodDetails: payment_gateway.UpiDetails{},
	})
	fmt.Printf("MakePayment_err_UPI: %v\n", err)
}

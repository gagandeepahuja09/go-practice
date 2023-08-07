package main

import (
	"fmt"

	"go-practice.com/machine-coding/payment_gateway/payment_gateway"
)

func setBankDistributionHelper() {
	errBankD := payment_gateway.SetMethodBankDistribution(payment_gateway.UPI, []payment_gateway.BankPercentage{
		{
			Percentage: 50,
			BankName:   payment_gateway.HDFC,
		},
		{
			Percentage: 60,
			BankName:   payment_gateway.ICICI,
		},
	})
	fmt.Printf("errBankD: %v\n", errBankD)

	errBankD = payment_gateway.SetMethodBankDistribution(payment_gateway.UPI, []payment_gateway.BankPercentage{
		{
			Percentage: 80,
			BankName:   payment_gateway.HDFC,
		},
		{
			Percentage: 20,
			BankName:   payment_gateway.ICICI,
		},
	})
	fmt.Printf("errBankD: %v\n", errBankD)

	errBankD = payment_gateway.SetMethodBankDistribution(payment_gateway.CreditCard, []payment_gateway.BankPercentage{
		{
			Percentage: 10,
			BankName:   payment_gateway.HDFC,
		},
		{
			Percentage: 90,
			BankName:   payment_gateway.ICICI,
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
		payment_gateway.CreditCard,
	})
	fmt.Printf("errFk: %v\n", errFk)

	errZomato := payment_gateway.AddSupportForMethods(cZomato, []payment_gateway.PaymentMethod{
		payment_gateway.CreditCard,
		payment_gateway.UPI,
	})
	fmt.Printf("errZomato: %v\n", errZomato)

	payment_gateway.ListSupportedMethods(cZomato)

	setBankDistributionHelper()
}

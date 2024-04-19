package designpatterns

import "fmt"

// Define a family of algorithms, encapsulate each one and make them interchangeable

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardStrategy struct{}

func (ccs *CreditCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paying with Credit Card the amount of %f", amount)
}

type PayPalStrategy struct{}

func (pps *PayPalStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paying with PayPal the amount of %f", amount)
}

type PaymentContext struct {
	PaymentStrategy PaymentStrategy
}

func (pc *PaymentContext) SetPaymentStrategy(ps PaymentStrategy) {
	pc.PaymentStrategy = ps
}

func (pc *PaymentContext) Pay(amount float64) string {
	return pc.PaymentStrategy.Pay(amount)
}

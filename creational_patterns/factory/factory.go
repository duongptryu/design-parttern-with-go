package factory

import (
	"fmt"
)

type PaymentOption int

const (
	Cash PaymentOption = iota + 1
	DebitCard
	CreditCard
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m PaymentOption) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	default:
		return nil, fmt.Errorf("Payment method %v not recognized\n", m)
	}
}

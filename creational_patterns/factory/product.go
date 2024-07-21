package factory

import "fmt"

type CashPM struct {
}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

type DebitCardPM struct {
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

type CreditCardPM struct {
}

func (c *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using new credit card implementation\n", amount)
}

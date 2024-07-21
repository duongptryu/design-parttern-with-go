package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method type Cash must exist", err)
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("Thee cash payment method message wasn't correct")
	}
	t.Log("LOG: ", msg)
}

func TestCreatePaymentMethodDebitCash(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method type DebitCard must exist", err)
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("Thee cash payment method message wasn't correct")
	}
	t.Log("LOG: ", msg)
}

func TestCreatePaymentMethodCreditCard(t *testing.T) {
	payment, err := GetPaymentMethod(CreditCard)
	if err != nil {
		t.Fatal("A payment method type CreditCard must exist", err)
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using new credit card") {
		t.Error("Thee cash payment method message wasn't correct")
	}
	t.Log("LOG: ", msg)
}

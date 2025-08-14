package payment

import (
	"fmt"
)

type Payment interface {
	pay(amount int)
}

type Cash struct{}
type Card struct{}
type Upi struct{}

func (c Cash) pay(amount int) {
	fmt.Println("amount is paid by cash", amount)
}
func (c Card) pay(amount int) {
	fmt.Println("amount is paid by Card", amount)
}
func (u Upi) pay(amount int) {
	fmt.Println("amountis is paid by UPI", amount)
}
func MakePayment(p Payment, amount int) {
	p.pay(amount)
}

type PaymentOrder struct {
	Method Payment
	Amount int
}

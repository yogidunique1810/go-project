package main

// path = module name + folder
import "github.com/yogidunique1810/go-project/payment"

func main() {
	allPayments := []payment.PaymentOrder{
		{payment.Cash{}, 300},
		{payment.Card{}, 500},
		{payment.Upi{}, 800},
	}
	for _, paymentId := range allPayments {
		payment.MakePayment(paymentId.Method, paymentId.Amount)
	}
}

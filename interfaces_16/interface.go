package main

import "fmt"

// first we made the payment type struct
// then we added a method inisde it
// then we made a razorpay struct
// then we added the method pay()

// suppose we want razor pay again , as our gateway , we need to make changes in our payment struct again
// for testing purposes also we need to make changes

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

type paypal struct {
}

func (p paypal) pay(amount float32) {
	fmt.Println("making the payment through paypal", amount)
}
func (p paypal) refund(amount float32, account string){
	fmt.Println("the refund for the account",account,"is ",amount)
}
func (p payment) makePayment(amount float32) {
	// razorPaypaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)
	// razorPaypaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {
}

// Open close princple it violates

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// suppose we want to implement the stripe

// suppose i want to make a fake payment gateway
type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake payment gateway for testing purposes", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment through using stripe", amount)
}
func main() {
	
	// stripePaymentGw := stripe{}
	// fakePayGw := fakepayment{}
	// testPayment := payment{
	// 	gateway: fakePayGw,
	// }
	// testPayment.makePayment(200)
	// newPayment := payment{
	// 	gateway: stripePaymentGw,
	// }
	// newPayment.makePayment(100)

	paypalGw := paypal{}
	newPayment := payment{
		gateway:paypalGw,
	}
	
	newPayment.gateway.refund(1000,"sabit")

	// fakePayGw := fakepayment{}
	// fakeTest := payment{
	// 	gateway:fakePayGw,
	// }
	// fakeTest.makePayment(200)

	// stripeGw := stripe{}
	// payThroughstripe := payment{
	// 	gateway: stripeGw,
	// }
	// payThroughstripe.makePayment(300)

}

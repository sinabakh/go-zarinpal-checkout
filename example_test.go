package zarinpal_test

import (
	"log"

	"github.com/sinabakh/go-zarinpal-checkout"
)

func ExampleNewZarinpal() {
	zarinPay, err := zarinpal.NewZarinpal("XXXX-XXXX-XXXX-XXXX", false)
	if err != nil {
		log.Fatal(err)
	}
	zarinPayTest, err := zarinpal.NewZarinpal("XXXX-XXXX-XXXX-XXXX", true)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(zarinPay)
	log.Println(zarinPayTest)
}

func ExampleZarinpal_NewPaymentRequest() {
	zarinPay, err := zarinpal.NewZarinpal("XXXX-XXXX-XXXX-XXXX", true)
	if err != nil {
		log.Fatal(err)
	}
	paymentURL, authority, statusCode, err := zarinPay.NewPaymentRequest(100, "http://localhost:3000", "Test", "", "")
	if err != nil {
		if statusCode == -3 {
			log.Println("Amount is not accepted in banking system")
		}
		log.Fatal(err)
	}
	log.Println(authority)  // Save authority in DB
	log.Println(paymentURL) // Send user to paymentURL
}

func ExampleZarinpal_PaymentVerification() {
	zarinPay, err := zarinpal.NewZarinpal("XXXX-XXXX-XXXX-XXXX", true)
	if err != nil {
		log.Fatal(err)
	}
	authority := "XXXXXXXXX" // Read authority from your storage (DB) or callback request
	amount := 1000           // The amount of payment in Tomans
	verified, refID, statusCode, err := zarinPay.PaymentVerification(amount, authority)
	if err != nil {
		if statusCode == 101 {
			log.Println("Payment is already verified")
		}
		log.Fatal(err)
	}
}

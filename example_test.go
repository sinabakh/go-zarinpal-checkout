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
	zarinPay, err := zarinpal.NewZarrinpal("XXXX-XXXX-XXXX-XXXX", true)
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

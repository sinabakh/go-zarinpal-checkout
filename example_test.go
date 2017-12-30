package zarinpal_test

import (
	"encoding/json"
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
		log.Fatal("error:", err, "statusCode:", statusCode)
	}
	log.Println("verified:", verified, "refID:", refID)
}

func ExampleZarinpal_UnverifiedTransactions() {
	zarinPay, err := zarinpal.NewZarinpal("XXXX-XXXX-XXXX-XXXX", true)
	if err != nil {
		log.Fatal(err)
	}
	authorities, statusCode, err := zarinPay.UnverifiedTransactions()
	if err != nil {
		log.Fatal("statusCode:", statusCode, "error:", err)
	}
	marshaledJSON, _ := json.Marshal(authorities)
	log.Println(string(marshaledJSON))
	// Output:
	// [{"Authority":"XXXX","Amount":100,"Channel":"WebGate","CallbackURL":"http://localhost:3000","Referer":"/","Email":"","CellPhone":"","Date":"2017-12-27 22:12:59"}]
}

func ExampleZarinpal_RefreshAuthority() {
	zarinPay, err := zarinpal.NewZarinpal("XXXX-XXXX-XXXX-XXXX", true)
	if err != nil {
		log.Fatal(err)
	}
	statusCode, err := zarinPay.RefreshAuthority("XXXX", 2000)
	if err != nil {
		log.Fatal("statusCode:", statusCode, "error:", err)
	}
}

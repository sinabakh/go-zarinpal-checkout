// Package zarinpal provides simple methods to work
// with Zarinpal (https://www.zarinpal.com/) checkout gateway
package zarinpal

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Zarinpal is the base struct for zarinpal payment
// gateway, one shall not create or manipulate instances
// if this struct manually and just use provided methods
// to woek with it.
type Zarinpal struct {
	MerchantID      string
	Sandbox         bool
	APIEndpoint     string
	PaymentEndpoint string
}

type paymentRequestReq struct {
	MerchantID  string
	Amount      int
	CallbackURL string
	Description string
	Email       string
	Mobile      string
}

type paymentRequestResp struct {
	Status    int
	Authority string
}

// NewZarinpal creates a new instance of zarinpal payment
// gateway with provided configs. It also tries to validate
// provided configs.
func NewZarinpal(merchantID string, sandbox bool) (*Zarinpal, error) {
	if len(merchantID) != 36 {
		return nil, errors.New("MerchantID must be 36 characters")
	}
	apiEndPoint := "https://www.zarinpal.com/pg/rest/WebGate/"
	paymentEndpoint := "https://www.zarinpal.com/pg/StartPay/"
	if sandbox == true {
		apiEndPoint = "https://sandbox.zarinpal.com/pg/rest/WebGate/"
		paymentEndpoint = "https://sandbox.zarinpal.com/pg/StartPay/"
	}
	return &Zarinpal{
		MerchantID:      merchantID,
		Sandbox:         sandbox,
		APIEndpoint:     apiEndPoint,
		PaymentEndpoint: paymentEndpoint,
	}, nil
}

// NewPaymentRequest gets a payment url from Zarinpal.
// amount is in Tomans (not Rials) format.
// email and mobile are optional.
//
// If error is not nil, you can check statusCode for
// specific error handler based on Zarinpal error codes.
// If statusCode is not 0, it means Zarinpal raised an error
// on their end and you can check the error code and its reason
// based on their documentation placed in
// https://github.com/ZarinPal-Lab/Documentation-PaymentGateway/archive/master.zip
func (zarinpal *Zarinpal) NewPaymentRequest(amount int, callbackURL, description, email, mobile string) (paymentURL, authority string, statusCode int, err error) {
	if amount < 1 {
		err = errors.New("amount must be a positive number")
		return
	}
	if callbackURL == "" {
		err = errors.New("callbackURL should not be empty")
		return
	}
	if description == "" {
		err = errors.New("description should not be empty")
		return
	}
	paymentRequest := paymentRequestReq{
		MerchantID:  zarinpal.MerchantID,
		Amount:      amount,
		CallbackURL: callbackURL,
		Description: description,
		Email:       email,
		Mobile:      mobile,
	}
	var resp paymentRequestResp
	err = zarinpal.request("PaymentRequest.json", &paymentRequest, &resp)
	if err != nil {
		return
	}
	if resp.Status == 100 {
		authority = resp.Authority
		paymentURL = zarinpal.PaymentEndpoint + resp.Authority
	} else {
		statusCode = resp.Status
		err = errors.New(strconv.Itoa(resp.Status))
	}
	return
}

func (zarinpal *Zarinpal) request(method string, data interface{}, res interface{}) error {
	reqBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", zarinpal.APIEndpoint+method, bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, res)
	if err != nil {
		err = errors.New("Zarinpal invalid json response")
		return err
	}
	return nil
}

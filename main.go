// Package zarinpal provides simple methods to work
// with Zarinpal (https://www.zarinpal.com/) checkout gateway
package zarinpal

import (
	"errors"
)

type zarinpal struct {
	MerchantID  string
	Sandbox     bool
	APIEndpoint string
}

// NewZarrinpal creates a new instance of zarinpal payment
// gateway with provided configs. It also tries to validate
// provided configs.
func NewZarrinpal(merchantID string, sandbox bool) (*zarinpal, error) {
	if len(merchantID) != 36 {
		return nil, errors.New("MerchantID must be 36 characters")
	}
	apiEndPoint := "https://www.zarinpal.com/pg/rest/WebGate/"
	if sandbox == true {
		apiEndPoint = "https://sandbox.zarinpal.com/pg/rest/WebGate/"
	}
	return &zarinpal{
		MerchantID:  merchantID,
		Sandbox:     sandbox,
		APIEndpoint: apiEndPoint,
	}, nil
}

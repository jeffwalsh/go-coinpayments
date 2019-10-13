package coinpayments

import (
	"net/url"
)

// WithdrawalRequest is what we sent to the API
type WithdrawalRequest struct {
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	MerchantID  string `json:"merchant_id"`
	PBNTag      string `json:"pbntag"`
	AutoConfirm int    `json:"auto_confirm"`
}

// WithdrawalResult is a result from the API for a Withdrawal command
type WithdrawalResult struct {
	Amount         string `json:"amount"`
	ID string `json:"id"`
	Status int `json:"status"` // 0 or 1. 0 = transfer created, waiting for email conf. 1 = transfer created with no email conf. 
}

// WithdrawalResponse is the response we expect from the API server.
type WithdrawalResponse struct {
	ErrorResponse
	Result *WithdrawalResult `json:"result"`
}

// CallCreateTransfer calls the create_Withdrawal command on the API
func (c *Client) CallCreateTransfer(req *WithdrawalRequest) (*WithdrawalResult, error) {

	// add in data specific to this Withdrawal, then forward the request to the call method
	data := url.Values{}
	data.Add("amount", req.Amount)
	data.Add("currency", req.Currency)
	data.Add("merchant", req.MerchantID)
	data.Add("pbntag", req.PBNTag)

	// make the actual call and unmarshal the response into our WithdrawalResponse struct
	var response WithdrawalResponse
	if err := c.Call(CmdCreateTransfer, data, &response); err != nil {
		return nil, err
	}

	// return the entire result to be used
	return response.Result, nil
}

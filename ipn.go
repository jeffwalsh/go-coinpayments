package coinpayments

import (
	"io"
	"io/ioutil"
	"net/url"
)

// IPNDepositResponse is a representation of a response received when any update is happening on a deposit
type IPNDepositResponse struct {
	Address    string `json:"address"`
	TxnID      string `json:"txn_id"`
	Status     string `json:"status"`
	StatusText string `json:"status_text"`
	Currency   string `json:"currency"`
	Confirms   string `json:"confirms"`
	Amount     string `json:"amount"`
	AmountI    string `json:"amounti"` // amount in satoshis
	Fee        string `json:"fee"`
	FeeI       string `json:"feei"` // fee in satoshis
	DestTag    string `json:"dest_tag"`
}

// HandleIPNDeposit takes a http request and returns the post parameters. io.Reader is typically fulfilled by the `req.Body` if used in a HTTP request to handle
// the IPN methods.
// IE: cps.HandleIPNDeposit(req.Body)
func (c *Client) HandleIPNDeposit(reader io.Reader) (*IPNDepositResponse, error) {

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	return &IPNDepositResponse{
		Address:    values.Get("address"),
		TxnID:      values.Get("txn_id"),
		Status:     values.Get("status"),
		StatusText: values.Get("status_text"),
		Currency:   values.Get("currency"),
		Confirms:   values.Get("confirms"),
		Amount:     values.Get("amount"),
		AmountI:    values.Get("amounti"),
		Fee:        values.Get("fee"),
		FeeI:       values.Get("feei"),
		DestTag:    values.Get("dest_tag"),
	}, nil
}

// IPNAPIResponse is the response we expect back from the server when the command is "api"
type IPNAPIResponse struct {
	Status           string `json:"status"`
	StatusText       string `json:"status_text"`
	TxnID            string `json:"txn_id"`
	Currency1        string `json:"currency1"`
	Currency2        string `json:"currency2"`
	Amount1          string `json:"amount1"`
	Amount2          string `json:"amount2"`
	Fee              string `json:"fee"`
	BuyerName        string `json:"buyer_name"`
	Email            string `json:"email"`
	ItemName         string `json:"item_name"`
	ItemNumber       string `json:"item_number"`
	Invoice          string `json:"invoice"`
	Custom           string `json:"custom"`
	SendTX           string `json:"send_tx"` // the tx id of the payment to the merchant. only included when 'status' >= 100 and the payment mode is set to ASAP or nightly or if the payment is paypal passthru
	ReceivedAmount   string `json:"received_amount"`
	ReceivedConfirms string `json:"received_confirms"`
}

// HandleIPNAPI handles the IPN API on input and gives a response
func (c *Client) HandleIPNAPI(reader io.Reader) (*IPNAPIResponse, error) {

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}
	return &IPNAPIResponse{
		Status:           values.Get("status"),
		StatusText:       values.Get("status_text"),
		TxnID:            values.Get("txn_id"),
		Currency1:        values.Get("currency1"),
		Currency2:        values.Get("currency2"),
		Amount1:          values.Get("amount1"),
		Amount2:          values.Get("amount2"),
		Fee:              values.Get("fee"),
		BuyerName:        values.Get("buyer_name"),
		Email:            values.Get("email"),
		ItemName:         values.Get("item_name"),
		ItemNumber:       values.Get("item_number"),
		Invoice:          values.Get("invoice"),
		Custom:           values.Get("custom"),
		SendTX:           values.Get("send_tx"),
		ReceivedAmount:   values.Get("received_amount"),
		ReceivedConfirms: values.Get("received_confirms"),
	}, nil
}

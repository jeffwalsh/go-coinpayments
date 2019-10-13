package coinpayments

import (
	"net/url"
)

// TransactionRequest is what we sent to the API
type TransactionRequest struct {
	// required
	Amount     string `json:"amount"`
	Currency1  string `json:"currency1"`
	Currency2  string `json:"currency2"`
	BuyerEmail string `json:"buyer_email"`

	// optional
	Address    string `json:"address,omitempty"`
	BuyerName  string `json:"buyer_name,omitempty"`
	ItemName   string `json:"item_name,omitempty"`
	ItemNumber string `json:"item_number,omitempty"`
	Invoice    string `json:"invoice,omitempty"`
	Custom     string `json:"custom,omitempty"`
	IPNURL     string `json:"ipn_url,omitempty"`
	SuccessURL string `json:"success_url,omitempty"`
	CancelURL  string `json:"cancel_url,omitempty"`
}

// TransactionResult is a result from the API for a transaction command
type TransactionResult struct {
	Amount         string `json:"amount"`
	Address        string `json:"address"`
	TxnID          string `json:"txn_id"`
	ConfirmsNeeded string `json:"confirms_needed"`
	Timeout        uint32 `json:"timeout"`
	StatusURL      string `json:"status_url"`
	CheckoutURL    string `json:"checkout_url"`
	QRCodeURL      string `json:"qrcode_url"`
}

// TransactionResponse is the response we expect from the API server.
type TransactionResponse struct {
	ErrorResponse
	Result *TransactionResult `json:"result"`
}

// CallCreateTransaction calls the create_transaction command on the API
func (c *Client) CallCreateTransaction(req *TransactionRequest) (*TransactionResult, error) {

	// add in data specific to this transaction, then forward the request to the call method
	data := url.Values{}
	data.Add("amount", req.Amount)
	data.Add("currency1", req.Currency1)
	data.Add("currency2", req.Currency2)
	data.Add("buyer_email", req.BuyerEmail)

	if req.Address != "" {
		data.Add("address", req.Address)
	}
	if req.BuyerName != "" {
		data.Add("buyer_name", req.BuyerName)
	}
	if req.ItemName != "" {
		data.Add("item_name", req.ItemName)
	}
	if req.ItemNumber != "" {
		data.Add("item_number", req.ItemNumber)
	}
	if req.Invoice != "" {
		data.Add("invoice", req.Invoice)
	}
	if req.Custom != "" {
		data.Add("custom", req.Custom)
	}
	if req.IPNURL != "" {
		data.Add("ipn_url", req.IPNURL)
	}
	if req.SuccessURL != "" {
		data.Add("success_url", req.SuccessURL)
	}
	if req.CancelURL != "" {
		data.Add("cancel_url", req.CancelURL)
	}

	// make the actual call and unmarshal the response into our TransactionResponse struct
	var response TransactionResponse
	if err := c.Call(CmdCreateTransaction, data, &response); err != nil {
		return nil, err
	}

	// return the entire result to be used
	return response.Result, nil
}

// CallbackAddressRequest is what the server expects for the get_callback_address command
type CallbackAddressRequest struct {
	Currency string `json:"currency"`
	IPNURL   string `json:"ipn_url"`
}

// CallbackAddressResult is a result from the API for a transaction command
type CallbackAddressResult struct {
	Address string `json:"address"`  // address to deposit selected coin into our wallet
	PubKey  string `json:"pubkey"`   // NXT only
	DestTag string `json:"dest_tag"` // for coins needing destination tag
}

// CallbackAddressResponse is the response given by the API to get_callback_address command
type CallbackAddressResponse struct {
	ErrorResponse
	Result *CallbackAddressResult `json:"result"`
}

// CallGetCallbackAddress calls the get_callback_address command on the api
func (c *Client) CallGetCallbackAddress(req *CallbackAddressRequest) (*CallbackAddressResponse, error) {

	// add in data specific to this command, then forward the request to the call method
	data := url.Values{}
	data.Add("currency", req.Currency)
	data.Add("ipn_url", req.IPNURL)

	// make the actual call and unmarshal the response into our TransactionResponse struct
	var response CallbackAddressResponse
	if err := c.Call(CmdGetCallbackAddress, data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// DepositAddressRequest is what the API expects on the get_deposit_address command
type DepositAddressRequest struct {
	Currency string `json:"currency"`
}

// CallGetDepositAddress calls the get_deposit_address command on the API, which has the same response as
// the get_callback_address command, and thus uses it's response to unmarshal.
func (c *Client) CallGetDepositAddress(req *DepositAddressRequest) (*CallbackAddressResponse, error) {

	// add in data specific to this command, then forward the request to the call method
	data := url.Values{}
	data.Add("currency", req.Currency)
	var response CallbackAddressResponse
	if err := c.Call(CmdGetDepositAddress, data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// TxInfoRequest is what the get_tx_info command expects
type TxInfoRequest struct {
	TxID string `json:"txid"`
	Full string `json:"full"`
}

// TxInfoResponse is the response we receive from the API. The result field will not be populated on error.
type TxInfoResponse struct {
	ErrorResponse
	Result map[string]interface{} `json:"result"`
}

// CallGetTxInfo calls the get_tx_info command on the API
func (c *Client) CallGetTxInfo(req *TxInfoRequest) (*TxInfoResponse, error) {
	// add in data specific to this command, then forward the request to the call method
	data := url.Values{}
	data.Add("txid", req.TxID)
	data.Add("full", req.Full)
	var response TxInfoResponse
	if err := c.Call(CmdGetTxInfo, data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// TxListRequest is what the get_tx_info command expects
type TxListRequest struct {
	Limit string `json:"limit"`
	Start string `json:"start"`
	Newer string `json:"newer"`
	All   string `json:"all"`
}

// TxListResponse is the response we receive from the API. The result field will not be populated on error.
type TxListResponse struct {
	ErrorResponse
	Result []string `json:"result"`
}

// CallGetTxList calls the get_tx_list command on the API
func (c *Client) CallGetTxList(req *TxListRequest) (*TxListResponse, error) {
	// default 25
	if req.Limit == "" {
		req.Limit = "25"
	}

	// add in data specific to this command, then forward the request to the call method
	data := url.Values{}
	data.Add("limit", req.Limit)
	data.Add("start", req.Start)
	data.Add("newer", req.Newer)
	data.Add("all", req.All)

	var response TxListResponse
	if err := c.Call(CmdGetTxList, data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

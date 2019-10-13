package coinpayments

import (
	"net/url"
)

// BasicInfoResult is the result we receive back as part of the basic info command response
type BasicInfoResult struct {
	Username   string `json:"username"`
	MerchantID string `json:"merchant_id"`
	Email      string `json:"email"`
	PublicName string `json:"public_name"`
}

// BasicInfoResponse is a response we get back from the basic info command
type BasicInfoResponse struct {
	ErrorResponse
	Result *BasicInfoResult `json:"result"`
}

// GetError implements the Response interface
func (resp *BasicInfoResponse) GetError() string {
	return resp.Error
}

// CallGetBasicInfo calls the get_basic_info command on the API, returning info about our merchant account via the API keys.
func (c *Client) CallGetBasicInfo() (*BasicInfoResult, error) {

	var resp BasicInfoResponse
	if err := c.Call(CmdGetBasicInfo, url.Values{}, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

// RatesRequest holds the params the API expects for the rates command
type RatesRequest struct {
	Short    string `json:"short"`
	Accepted string `json:"accepted"`
}

// RatesResult is the result we receive back as part of the rates command response
type RatesResult struct {
	IsFiat     int8   `json:"is_fiat"`
	RateBTC    string `json:"rate_btc"`
	LastUpdate string `json:"last_update"`
}

// RatesResponse is a response we get back from the rates command
type RatesResponse struct {
	ErrorResponse
	Result map[string]RatesResult `json:"result"`
}

// CallRates calls the get_basic_info command on the API, returning info about our merchant account via the API keys.
func (c *Client) CallRates(req *RatesRequest) (map[string]RatesResult, error) {

	data := url.Values{}
	data.Add("short", req.Short)
	data.Add("accepted", req.Accepted)
	var resp RatesResponse
	if err := c.Call(CmdRates, data, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

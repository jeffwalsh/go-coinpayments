package coinpayments

import (
	"net/url"
)

// BalancesRequest is what we sent to the API
type BalancesRequest struct {
	All string `json:"all"`
}

// BalancesResult is a result from the API for a balances command
type BalancesResult struct {
	Balance  int    `json:"balance"`  // balance as integer in satoshis
	Balancef string `json:"balancef"` // balance as floating point number
}

// BalancesResponse is the response we expect from the API server.
type BalancesResponse struct {
	ErrorResponse
	Result map[string]BalancesResult `json:"result"`
}

// CallBalances calls the balances command on the API
func (c *Client) CallBalances(req *BalancesRequest) (map[string]BalancesResult, error) {

	// add in data specific to this Balances, then forward the request to the call method
	data := url.Values{}
	data.Add("all", req.All)

	// make the actual call and unmarshal the response into our BalancesResponse struct
	var response BalancesResponse
	if err := c.Call(CmdBalances, data, &response); err != nil {
		return nil, err
	}

	// return the entire result to be used
	return response.Result, nil
}

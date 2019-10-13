package coinpayments

import "net/url"

// ConvertLimitRequest is used to hold our request parameters for getting conversion limits
type ConvertLimitRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// ConvertLimitResponse hold's a response to the Get Conversion Limits API call
type ConvertLimitResponse struct {
	Error  string `json:"error"`
	Result struct {
		Min string `json:"min"`
		Max string `json:"max"`
	} `json:"result"`
}

// CallGetConversionLimits calls the get conversion limits comand on the API
func (c *Client) CallGetConversionLimits(req *ConvertLimitRequest) (*ConvertLimitResponse, error) {

	data := url.Values{}
	data.Add("from", req.From)
	data.Add("to", req.To)

	var response ConvertLimitResponse
	if err := c.Call(CmdGetConversionLimits, data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

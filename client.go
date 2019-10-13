package coinpayments

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// These variables come from the Coinpayments API itself.
var (
	baseURL           = "https://www.coinpayments.net/api.php"
	version           = "1"  // the version of the API - right now, it is only at "1" with no other options.
	successResponse   = "ok" // successResponse is what the Error field of the Response object will be populated with if there were no problems.
	successStatusCode = "200 OK"
	formatJSON        = "json" // the string the API expects as the format parameter
)

// Errors
var (
	ErrCommandDoesntExist = errors.New("command is not supported by api client")
)

// Commands supported by the API
var (
	CmdCreateTransaction   = "create_transaction"
	CmdGetBasicInfo        = "get_basic_info"
	CmdRates               = "rates"
	CmdBalances            = "balances"
	CmdGetCallbackAddress  = "get_callback_address"
	CmdGetDepositAddress   = "get_deposit_address"
	CmdGetTxInfo           = "get_tx_info"
	CmdGetTxInfoMulti      = "get_tx_info_multi"
	CmdGetTxList           = "get_tx_ids"
	CmdGetConversionLimits = "convert_limits"
	CmdCreateTransfer      = "create_transfer"
)

// Reader is our example implementation of a Reader.
// If you don't have a valid Reader to pass in to the various Call methods like a `req.Body`,
// you can create an instance of this Reader and pass in your data as a byte slice.
// ie: client.CallGetDepositAddress(&coinpayments.Reader{data: []byte("{\"currency\":\"USD\"}")}`
type Reader struct {
	Data      []byte
	readIndex int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.readIndex >= int64(len(r.Data)) {
		err = io.EOF
		return
	}

	n = copy(p, r.Data[r.readIndex:])
	r.readIndex += int64(n)
	return
}

// ErrorResponse is an error we get back from coinpayments. All responses must have this implemented as well.
type ErrorResponse struct {
	Error string `json:"error"`
}

// HTTPClient is an interface we rely on to send create requests
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client represents our wrapper for the CoinPayments API
type Client struct {
	commands             []string
	baseURL              string
	httpClient           HTTPClient
	privateKey           string
	publicKey            string
	IPNSecret            string
	IPNURL               string
	BTCForwardingAddress string
	ETHForwardingAddress string
}

// SupportedCommands returns a slice of strings with all the available commands
func SupportedCommands() []string {
	return []string{CmdCreateTransaction, 
		CmdGetBasicInfo, 
		CmdRates, 
		CmdBalances, 
		CmdGetCallbackAddress, 
		CmdGetDepositAddress, 
		CmdGetTxInfo, 
		CmdGetTxInfoMulti, 
		CmdGetTxList, 
		CmdCreateTransfer, 
		CmdGetConversionLimits,
	}
}

// NewClient returns a new Client struct, initialized and ready to go
func NewClient(cfg *Config, httpClient HTTPClient) (*Client, error) {
	if cfg.PrivateKey == "" || cfg.PublicKey == "" {
		return nil, errors.New("private and/or public key missing from config struct")
	}

	if httpClient == nil {
		return nil, errors.New("No http client provided")
	}
	// build out our list of necessary commands the API has been instructed to use so far
	commands := make([]string, 3)
	commands = append(commands, SupportedCommands()...)
	cp := &Client{commands: commands, baseURL: baseURL, httpClient: httpClient, privateKey: cfg.PrivateKey, publicKey: cfg.PublicKey, IPNSecret: cfg.IPNSecret, IPNURL: cfg.IPNURL,
		BTCForwardingAddress: cfg.BTCForwardingAddress, ETHForwardingAddress: cfg.ETHForwardingAddress}
	return cp, nil
}

func (c *Client) unmarshal(reader io.Reader, req interface{}) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &req)
}

// computeHMAC returns our hmac because on the private key of our account
func (c *Client) computeHMAC(data string) (string, error) {
	hash := hmac.New(sha512.New, []byte(c.privateKey))
	if _, err := hash.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// Call sends a request with the given cmd and data, and then unmarshals the response into the given responseStruct.
// If you don't want to pass a Reader to the individual functions, you can use Call to call any command directly.
// You just need to build out your own data with an url.Values, and create an instance of the responsestruct expected by your command.
// ie: data := url.Values{}
// data.Add("currency", "USD")
// resp := coinpayments.DepositAddressResponse{}
// err := c.Call(coinpayments.CmdGetDepositAddress, data, &resp)
func (c *Client) Call(cmd string, data url.Values, responseStruct interface{}) error {
	if !stringExistsInSlice(c.commands, cmd) {
		return ErrCommandDoesntExist
	}
	return c.call(cmd, data, responseStruct)
}

// call sends a request with the given cmd and data, and then unmarshals the response into the given responseStruct.
func (c *Client) call(cmd string, data url.Values, responseStruct interface{}) error {
	data.Add("key", c.publicKey)
	data.Add("version", version)
	data.Add("cmd", cmd)
	data.Add("format", formatJSON)

	dataString := data.Encode()
	// generate hmac hash of data and private key
	hash, err := c.computeHMAC(dataString)
	if err != nil {
		return err
	}
	// create the request using the url and url values
	req, err := http.NewRequest("POST", c.baseURL, strings.NewReader(dataString))
	if err != nil {
		return err
	}

	// add necessary headers for API request
	req.Header.Add("HMAC", hash)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(dataString)))

	// do the actual request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.Status != successStatusCode {
		ret := fmt.Sprintf("failed to make api call: expected status %v, got %s", successStatusCode, resp.Status)
		return errors.New(ret)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	// unmarshal only the error first, as the server returns invalid json if there is an error. it comes in the form of
	// {"error":"error", "result":[]}, which cannot unmarshal the invalid array to a struct.
	cpError := ErrorResponse{}
	if err := json.Unmarshal(body, &cpError); err != nil {
		return err
	}

	// check the error to see if it was OK
	if cpError.Error != successResponse {
		return errors.New(cpError.Error)
	}

	// return the unmarshalled response and an error if it occurred
	return json.Unmarshal(body, responseStruct)
}

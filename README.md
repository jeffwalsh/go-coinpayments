# Golang Coinpayments wrapper library

![Tryvium Travels LTD logo](assets/tryvium-logo.png)

> NOTE: This project started as a fork of jeffwalsh/go-coinpayments project, but rapidly evolved to its own. While we thank and credit original creator for the job so far, we are going separate ways for the time being.
>
> This is the only coinpayments golang wrapper library officially supported and maintained by Tryvium Travels LTD

Coinpayments Golang wrapper. To use, instantiate a `*coinpayments.Client` by calling the `coinpayments.NewClient` function. You will need to pass in a valid instance of
the `*coinpayments.Config` struct, consisting of your private and public keys, as well as an instance of your desired http client.  
  
This looks like:
``` go
clientConfig := &coinpayments.Config{
  PublicKey: "yourpublickey",
  PrivateKey: "yourprivatekey",
}
httpClient := &http.Client{
  Timeout: 10 * time.Second,
}
client, err := coinpayments.NewClient(clientConfig, httpClient)
if err != nil {
  // handle error...
}
```
  
Once instantiated, you can use the client to make API calls. Each call has their own Request struct, and returns it's own Response struct.

An example of the `create_transaction` command being called:
``` go
resp, err := client.CallCreateTransaction(
  &coinpayments.TransactionRequest{
    Amount: "1", // the amount is the value of the fiat currency you wish to receive
    Currency1: "USD", // currency that amount is specified in (USD, CAD, JPY, etc)
    Currency2: "BTC", // currency in crypto that you wish to receive (BTC, ETC, LTC, etc)
    BuyerEmail: "test@email.com", // our client requires the buyer email address
  }
)
if err != nil {
  // handle error...
}
```

# Usage 

You can use Call to call any command directly.
You just need to build out your own data with an `url.Values`, and create an instance of the responsestruct expected by your command, then pass in the relevant
command from the coinpayments package, ie, `coinpayments.CmdCreateTransaction`.

Here is an example of getting the deposit address:  

``` go
var data url.Values
data.Add("currency", "USD")

var resp coinpayments.DepositAddressResponse
err := c.Call(coinpayments.CmdGetDepositAddress, data, &resp)
```

# Testing the library

You need to export two environment variables for the tests to run - your public key, and private key.  
``` bash
export COINPAYMENTS_PUBLIC_KEY=yourpublickeyhere
export COINPAYMENTS_PRIVATE_KEY=yourprivatekeyhere

go test ./...
```

If you want to run the PBN Tag tests and have purchased a PBN Tag, you can also export it, and those tests will run. Otherwise, they will be ignored.

``` bash
export COINPAYMENTS_PBN_TAG=yourpbntaghere
```

# How to contribute

First of all, thank you for that ❤️

If you want to contribute, just make a PR against the `main` branch of this repo.

The PR name must state which issue is solved and a description of your changes.

> Please do not close 2 issues with the same commit, do 2 separate commits or more.

# Implemented Methods

## Information Commands
- [x] Get Basic Account Info
- [x] Get Exchange Rates / Supported Coins
- [x] Get Coin Balances
- [x] Get Deposit Address

## Receiving Payments
- [x] Create Transaction
- [x] Callback Addresses
- [x] Get TX Info
- [x] Get TX List

## Withdrawals / Transfers
- [x] Create Transfer
- [ ] Create Withdrawal / Mass Withdrawal
- [ ] Convert Coins
- [ ] Get Withdrawal History
- [ ] Get Withdrawal Info
- [ ] Get Conversion Info

## $PayByName ( PBN )
- [ ] Get Profile Information
- [ ] Get Tag List
- [ ] Update Tag Profile
- [ ] Claim Tag

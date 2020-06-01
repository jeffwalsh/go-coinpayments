# coinpayments

Coinpayments Golang wrapper. To use, instantiate a `*coinpayments.Client` by calling the `coinpayments.NewClient` function. You will need to pass in a valid instance of
the `*coinpayments.Config` struct, consisting of your private and public keys, as well as an instance of your desired http client.  
  
This looks like: 
`client, err := coinpayments.NewClient(&coinpayments.Config{ PublicKey: "yourpublickey", PrivateKey: "yourprivatekey"}, &http.Client{ Timeout: 10 * time.Second})`.
  
Once instantiated, you can use the client to make API calls. Each call has their own Request struct, and returns it's own Response struct.

An example of the `create_transaction` command being called:  

  `client.CallCreateTransaction(&coinpayments.TransactionRequest{Amount: "1", Currency1: "USD", Currency2: "BTC", BuyerEmail: "test@email.com")`, where the amount is the value of the fiat currency you wish to receive.  
  currency1 is the type of currency that amount is specified in (USD, CAD, JPY, etc).  
  currency2 is the type of currency in crypto that you wish to receive (BTC, ETC, LTC, etc).  
  buyerEmail is optional with the API but mandatory for our client.  

# Call
You can use Call to call any command directly.
You just need to build out your own data with an url.Values, and create an instance of the responsestruct expected by your command. Then pass in the relevant
command from the coinpayments package, ie, `coinpayments.CmdCreateTransaction`.

Here is an example of getting the deposit address:  
ie: 
```
data := url.Values{}
data.Add("currency", "USD")
resp := coinpayments.DepositAddressResponse{}
err := c.Call(coinpayments.CmdGetDepositAddress, data, &resp)
```

# tests

You need to export two environment variables for the tests to run - your public key, and private key.  
```
export COINPAYMENTS_PUBLIC_KEY=yourpublickeyhere
export COINPAYMENTS_PRIVATE_KEY=yourprivatekeyhere

go test ./...
```

If you want to run the PBN Tag tests and have purchased a PBN Tag, you can also export it, and those tests will run. Otherwise, they will be ignored.

```
export COINPAYMENTS_PBN_TAG=yourpbntaghere
```


# Implemented Methods

## Information Commands
[ x ] - Get Basic Account Info
[ x ] - Get Exchange Rates / Supported Coins
[ x ] - Get Coin Balances
[ x ] - Get Deposit Address

## Receiving Payments
[ x ] - Create Transaction
[ x ] - Callback Addresses
[ x ] - Get TX Info
[ x ] - Get TX List

## Withdrawals / Transfers
[ x ] - Create Transfer
[ ] - Create Withdrawal / Mass Withdrawal
[ ] - Convert Coins
[ ] - Get Withdrawal History
[ ] - Get Withdrawal Info
[ ] - Get Conversion Info

## $PayByName ( PBN )
[ ] - Get Profile Information
[ ] - Get Tag List
[ ] - Update Tag Profile
[ ] - Claim Tag

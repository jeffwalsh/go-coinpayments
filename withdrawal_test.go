package coinpayments_test

import (
	"os"
	"testing"

	"github.com/jeffwalsh/coinpayments"
)

func TestCallCreateTransfer(t *testing.T) {
	// only run thse tests if we are provided a PBN tag.
	pbnTag, ok := os.LookupEnv("COINPAYMENTS_PBN_TAG")
	if ok {
		client, err := testClient()
		if err != nil {
			t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
		}

		_, err = client.CallCreateTransfer(&coinpayments.WithdrawalRequest{Amount: "1", Currency: "BTC", MerchantID: "merchantID", PBNTag: pbnTag, AutoConfirm: 0})
		if err != nil {
			t.Fatalf("Could not call create Withdrawal: %s", err.Error())
		}
	}

}

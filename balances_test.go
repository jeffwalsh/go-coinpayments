package coinpayments_test

import (
	"testing"

	"github.com/tryvium-travels/go-coinpayments"
)

func TestCallBalances(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallBalances(&coinpayments.BalancesRequest{All: "1"})
	if err != nil {
		t.Fatalf("Could not call balances %s", err.Error())
	}
}

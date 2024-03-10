package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet:= Wallet{}

	wallet.Deposit(7)

	result:= wallet.Balance()
	expected:= 7

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}
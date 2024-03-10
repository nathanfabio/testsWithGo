package pointerserrors

import "testing"


func TestWallet(t *testing.T) {
	wallet:= Wallet{}
	wallet.Deposit(Bitcoin(7))

	result:= wallet.Balance()
	expected:= Bitcoin(7)

	if result != expected {
		t.Errorf("result %s, expected %s", result, expected)
	}
}
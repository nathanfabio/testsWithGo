package pointerserrors

import "testing"


func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet:= Wallet{}
		wallet.Deposit(Bitcoin(7))
		
		checkBalance(t, wallet, Bitcoin(7))
	})

	t.Run("Withdraw with sufficient balance", func(t *testing.T) {
		wallet:= Wallet{Bitcoin(52)}
		err:= wallet.Withdraw(Bitcoin(40))

		checkBalance(t, wallet, Bitcoin(12))
		checkNoError(t, err)
	})

	t.Run("Insufficient balance", func(t *testing.T) {
		initialBalance:= Bitcoin(5)
		wallet:= Wallet{initialBalance}

		err:= wallet.Withdraw(Bitcoin(7))

		checkBalance(t, wallet, initialBalance)
		checkError(t, err, ErrorInsufficientBalance)
	})
}

func checkBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		result:= wallet.Balance()

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
}

func checkError(t *testing.T, result error, expected error) {
	t.Helper()
	if result == nil {
		t.Fatal("expected an error but none occurred")
	}

	if result != expected {
		t.Errorf("result error %s, expected error %s", result, expected)
	}
}

func checkNoError(t *testing.T, result error) {
	t.Helper()
	if result != nil {
		t.Fatal("unexpected error received")
	}
}
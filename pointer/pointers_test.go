package pointer

import (
	"testing"
)

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
    if got != nil {
        t.Fatal("got an error but didnt want one")
    }
}

func TestWallet(t *testing.T) {
	// assertError := func(t *testing.T, err error) {
	// 	if err == nil {
	// 		t.Error("wanted an error but didnt get one")
	// 	}
	// }
	t.Run("Deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(Bitcoin(100))
		want := Bitcoin(100)
		assertBalance(t, wallet, want)
	
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
	
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
	
}

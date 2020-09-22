package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(15))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrorInsufficientFunds)
	})
}

func TestBitcoin_String(t *testing.T) {
	got := Bitcoin(10).String()
	want := "10 BTC"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertBalance(t *testing.T, wallet Wallet, balance Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != balance {
		t.Errorf("got %s want %s", got, balance)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Wanted an error but didn't get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got an error %q but didn't want one", err)
	}
}

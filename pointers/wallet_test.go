package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		want := Bitcoin(10.0)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20.0)}
		err := wallet.Withdraw(Bitcoin(10.0))
		want := Bitcoin(10.0)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20.0)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100.0))

		assertError(t, err, ErrBalanceTooLow)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn`t want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("We expect an error at this stage")
	}
	if got != want {
		t.Errorf("want Error %q, but got %q", got, want)
	}
}

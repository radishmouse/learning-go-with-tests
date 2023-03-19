package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("can deposit to wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("can withdraw from wallet", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Errorf("expected error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error but got one")
	}
}

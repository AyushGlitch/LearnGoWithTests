package pointers

import "testing"

func PointersTest (t *testing.T) {
	bitcoin := Bitcoin(10)
	wallet := Wallet{bitcoin}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(20)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
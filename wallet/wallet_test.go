package wallet

import "testing"

func TestWallet_Deposit(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(BitCoin(10))

	got := wallet.Balance()
	want := BitCoin(10)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestWallet_Withdraw(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(BitCoin(10))

	err := wallet.Withdraw(BitCoin(10))

	if err != nil {
		t.Error(err)
	}

	got := wallet.Balance()

	want := BitCoin(0)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestWallet_Withdraw2(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(BitCoin(10))

	err := wallet.Withdraw(BitCoin(20))

	if err == nil {
		t.Error("Wanted error")
	}

	if wallet.balance != BitCoin(10) {
		t.Errorf("got %s, want %s", wallet.balance, BitCoin(10))
	}
}

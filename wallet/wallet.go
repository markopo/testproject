package wallet

import (
	"errors"
	"fmt"
)

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount BitCoin) error {

	if amount > w.balance {
		return errors.New("cannot withdraw more than a balance")
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

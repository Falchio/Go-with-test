package pointers

import "errors"

var ErrBalanceTooLow = errors.New("Balance too low")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrBalanceTooLow
	}
	w.balance -= amount
	return nil
}

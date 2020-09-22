package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin represents a bitcoin
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet represents a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit a given amount of bitcoins
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrorInsufficientFunds indicates insufficient funds for withdraw operation
var ErrorInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw a given amount of bitcoins
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrorInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns a wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

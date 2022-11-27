package tests

import (
	"errors"
	"fmt"
	"log"
)

// Error Definition of Wallet
var (
	ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
)

// Bitcoin in wallet
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet for Bitcoin
type Wallet struct {
	balance Bitcoin
}

// Deposit of the Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	log.Printf("address of balance in Deposit: %p", &w.balance)
	w.balance += amount
}

// Ballence of the Wallet
func (w *Wallet) Ballence() Bitcoin {
	return w.balance
}

// Withdraw of the Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	log.Printf("current balance in Wallet: %s", w.balance)
	return nil
}

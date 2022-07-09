package subsystem

import (
	"errors"
	"log"
)

type Wallet struct {
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) CreditBalance(amount int) {
	w.balance += amount
	log.Println("wallet balance added successfully")
	return
}

func (w *Wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("not sufficient balance")
	}

	log.Println("sufficient wallet balance")
	w.balance -= amount
	return nil
}

func NewWallet() *Wallet {
	return &Wallet{balance: 0}
}

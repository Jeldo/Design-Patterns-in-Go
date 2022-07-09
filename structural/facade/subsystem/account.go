package subsystem

import (
	"errors"
	"log"
)

type Account struct {
	name string
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) CheckAccount(name string) error {
	if a.name != name {
		return errors.New("incorrect account name")
	}

	log.Println("account verified")
	return nil
}

func NewAccount(name string) *Account {
	return &Account{name: name}
}

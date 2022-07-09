package facade

import (
	"DesignPatterns/structural/facade/subsystem"
	"log"
)

type WalletFacade struct {
	account      *subsystem.Account
	wallet       *subsystem.Wallet
	securityCode *subsystem.SecurityCode
	notification *subsystem.Notification
	ledger       *subsystem.Ledger
}

func (f *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	log.Println("start add money to wallet")
	if err := f.account.CheckAccount(accountID); err != nil {
		return err
	}

	if err := f.securityCode.CheckCode(securityCode); err != nil {
		return err
	}

	f.wallet.CreditBalance(amount)
	f.notification.SendWalletCreditNotification()
	f.ledger.NewEntry(accountID, "credit", amount)
	return nil
}

func (f *WalletFacade) DeductMoneyToWallet(accountID string, securityCode int, amount int) error {
	log.Println("start debit money from wallet")
	if err := f.account.CheckAccount(accountID); err != nil {
		return err
	}

	if err := f.securityCode.CheckCode(securityCode); err != nil {
		return err
	}

	if err := f.wallet.DebitBalance(amount); err != nil {
		return err
	}

	f.notification.SendWalletDebitNotification()
	f.ledger.NewEntry(accountID, "debit", amount)
	return nil
}
func NewWalletFacade(accountID string, code int) *WalletFacade {
	return &WalletFacade{
		account:      subsystem.NewAccount(accountID),
		wallet:       subsystem.NewWallet(),
		securityCode: subsystem.NewSecurityCode(code),
		notification: subsystem.NewNotification(),
		ledger:       subsystem.NewLedger(),
	}
}


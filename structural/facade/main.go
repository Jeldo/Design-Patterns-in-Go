package main

import (
	"DesignPatterns/structural/facade/facade"
	"log"
)

func main() {
	accountID := "abc"
	securityCode := 1234

	walletFacade := facade.NewWalletFacade(accountID, securityCode)

	if err := walletFacade.AddMoneyToWallet(accountID, securityCode, 10); err != nil {
		log.Printf("[ERROR] %s\n", err.Error())
		return
	}

	if err := walletFacade.DeductMoneyToWallet(accountID, securityCode, 5); err != nil {
		log.Printf("[ERROR] %s\n", err.Error())
		return
	}
}

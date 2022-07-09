package subsystem

import "log"

type Ledger struct{}

func NewLedger() *Ledger {
	return &Ledger{}
}

func (l *Ledger) NewEntry(accountId, txnType string, amount int) {
	log.Printf("created ledger entry for accountID=%s, txnType %s, amount %d\n", accountId, txnType, amount)
	return
}

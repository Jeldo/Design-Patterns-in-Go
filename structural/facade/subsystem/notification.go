package subsystem

import "log"

type Notification struct{}

func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) SendWalletCreditNotification() {
	log.Println("sending wallet credit notification")
}

func (n *Notification) SendWalletDebitNotification() {
	log.Println("sending wallet debit notification")
}

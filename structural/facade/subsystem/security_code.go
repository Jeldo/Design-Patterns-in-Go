package subsystem

import (
	"errors"
	"log"
)

type SecurityCode struct {
	code int
}

func (c *SecurityCode) Code() int {
	return c.code
}

func (c *SecurityCode) CheckCode(code int) error {
	if c.code != code {
		return errors.New("incorrect security code")
	}

	log.Println("security code verified")
	return nil
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

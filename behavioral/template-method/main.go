package main

import (
	"DesignPatterns/behavioral/template-method/email"
	"DesignPatterns/behavioral/template-method/otp"
	"DesignPatterns/behavioral/template-method/sms"
	"fmt"
)

func main() {
	smsOTP := &sms.SMS{}
	o := otp.NewOTP(smsOTP)
	err := o.GenerateAndSendOTP(4)
	fmt.Println("sms err=", err)

	emailOTP := &email.Email{}
	o = otp.NewOTP(emailOTP)
	err = o.GenerateAndSendOTP(5)
	fmt.Println("email err=", err)
}

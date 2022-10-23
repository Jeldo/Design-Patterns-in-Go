package main

import (
	"DesignPatterns/behavioral/template_method/email"
	"DesignPatterns/behavioral/template_method/otp"
	"DesignPatterns/behavioral/template_method/sms"
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

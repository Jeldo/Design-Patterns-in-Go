package sms

import (
	"fmt"
)

type SMS struct {
	//otp.OTPInterface  // duck typing: struct doesn't have to implement interface
}

func (s SMS) GenerateRandomOTP(otpLength int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s SMS) SaveOTPCache(otp string) {
	fmt.Printf("SMS saving otp: %s to cache\n", otp)
}

func (s SMS) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s SMS) SendNotification(message string) error {
	fmt.Println("SMS: sending sms: " + message)
	return nil
}

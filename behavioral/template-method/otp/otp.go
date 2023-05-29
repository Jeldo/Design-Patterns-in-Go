package otp

type OTPInterface interface {
	GenerateRandomOTP(otpLength int) string
	SaveOTPCache(otp string)
	GetMessage(otp string) string
	SendNotification(message string) error
}

type OTPImplementation struct {
	OTPInterface
}

func NewOTP(oi OTPInterface) OTPImplementation {
	return OTPImplementation{OTPInterface: oi}
}

func (o OTPImplementation) GenerateAndSendOTP(otpLength int) error {
	otp := o.GenerateRandomOTP(otpLength)
	o.SaveOTPCache(otp)
	message := o.GetMessage(otp)
	return o.SendNotification(message)
}

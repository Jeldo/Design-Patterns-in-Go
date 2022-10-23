package otp

type OTPInterface interface {
	GenerateRandomOTP(otpLength int) string
	SaveOTPCache(otp string)
	GetMessage(otp string) string
	SendNotification(message string) error
}

type OTPImpl struct {
	OTPInterface
}

func NewOTP(o OTPInterface) OTPImpl {
	return OTPImpl{OTPInterface: o}
}

func (o OTPImpl) GenerateAndSendOTP(otpLength int) error {
	otp := o.GenerateRandomOTP(otpLength)
	o.SaveOTPCache(otp)
	message := o.GetMessage(otp)
	return o.SendNotification(message)
}

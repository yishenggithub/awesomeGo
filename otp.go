package main

type iOtp interface {
	genRandomOtp(int) string
	saveOtpCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type otp struct {
	iOtp iOtp
}

// type otp struct {
// }

// func (o *otp) genAndSendOTP(iOtp iOtp, otpLength int) error {
//  otp := iOtp.genRandomOTP(otpLength)
//  iOtp.saveOTPCache(otp)
//  message := iOtp.getMessage(otp)
//  err := iOtp.sendNotification(message)
//  if err != nil {
//      return err
//  }
//  iOtp.publishMetric()
//  return nil
// }

func (o *otp) genAndSendOtp(otpLength int) error {
	otp := o.iOtp.genRandomOtp(otpLength)
	o.iOtp.saveOtpCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.publishMetric()
	return nil
}

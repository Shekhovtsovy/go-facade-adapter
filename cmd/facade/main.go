package main

import (
	"app/internal/facade"
	"app/internal/facade/driver/email"
	"app/internal/facade/driver/sms"
)

func main()  {
	// Typical facade usage
	smsDriver := sms.NewDriver()
	smsSender := facade.NewSender(smsDriver)
	smsSender.Send(1, "Hello!")

	// In case we have incorrect driver interface but for some reason can't correct the driver. Need to implement an adapter
	emailDriver := email.NewDriver()
	emailDriverAdapter := email.NewAdapter(emailDriver)
	emailSender := facade.NewSender(emailDriverAdapter)
	emailSender.Send(1, "Hello!")
}
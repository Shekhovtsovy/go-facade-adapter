package email

// Adapter is an interface of Email adapter which provides methods sending email to clients via Sender
type Adapter interface {
	Send(cid int, msg string) bool
}

type adapter struct {
	adapter Adapter
	driver  Driver
}

// Send Email to client
func (a *adapter) Send(cid int, msg string) bool {
	return a.driver.Deliver(msg, cid)
}

// NewAdapter returns a new Email Adapter for the Sender
func NewAdapter(a Driver) Adapter {
	return &adapter{
		driver: a,
	}
}

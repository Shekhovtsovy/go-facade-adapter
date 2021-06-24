package email

import "fmt"

// Driver is an interface of Email driver which provides methods sending email to clients
type Driver interface {
	Deliver(msg string, cid int) error
}

type driver struct {
}

// Send Email to client
func (d *driver) Deliver(msg string, cid int) error {
	fmt.Printf("Sending Email to client with id=%d: `%s`\n", cid, msg)
	return nil
}

// NewDriver returns a new Email Driver
func NewDriver() Driver {
	return &driver{}
}

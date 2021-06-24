package sms

import "fmt"

// Driver is an interface of SMS driver which provides methods sending SMS to clients
type Driver interface {
	Send(cid int, msg string) error
}

type driver struct {
}

// Send SMS to client
func (d *driver) Send(cid int, msg string) error {
	fmt.Printf("Sending SMS to client with id=%d: `%s`\n", cid, msg)
	return nil
}

// NewDriver returns a new SMS Driver
func NewDriver() Driver {
	return &driver{}
}

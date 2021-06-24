package facade

import "fmt"

type driver interface {
	Send(cid int, msg string) error
}

// Sender is an interface which provides methods sending messages to clients
type Sender interface {
	Send(cid int, msg string) error
}

type sender struct {
	driver driver
}

// Send message to client
func (s *sender) Send(cid int, msg string) error {
	fmt.Printf("Sending to client with id=%d: `%s`\n", cid, msg)
	if err := s.driver.Send(cid, msg); err != nil {
		return err
	}
	return nil
}

// NewSender returns a new Sender
func NewSender(driver driver) Sender {
	return &sender{
		driver: driver,
	}
}

package facade

import "fmt"

// Sender is an interface which provides methods sending messages to clients
type Sender interface {
	Send(clientId int, message string) bool
}

type driver interface {
	Send(cid int, msg string) bool
}

type sender struct {
	driver driver
}

// Send message to client
func (s *sender) Send(cid int, msg string) bool {
	fmt.Printf("Sending to client with id=%d: `%s`\n", cid, msg)
	s.driver.Send(cid, msg)
	return true
}

// NewSender returns a new Sender
func NewSender(d driver) Sender {
	return &sender{
		driver: d,
	}
}

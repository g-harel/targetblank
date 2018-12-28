package mock

import (
	"github.com/g-harel/targetblank/api/internal/email"
)

var sent []*Email

// Email represents a sent email.
type Email struct {
	To   string
	Sub  string
	Body string
}

// Sender is a mocked email.Client object.
type Sender struct{}

// NewSender creates a new mocked email.Client.
func NewSender() email.ISender {
	return &Sender{}
}

// Send records that an attemp was made to send an email.
func (s *Sender) Send(to, sub, body string) error {
	sent = append(sent, &Email{to, sub, body})
	return nil
}

// LastSentTo returns the last sent email.
func LastSentTo(a string) *Email {
	for i := range sent {
		e := sent[len(sent)-i-1]
		if e.To == a {
			return e
		}
	}
	return nil
}

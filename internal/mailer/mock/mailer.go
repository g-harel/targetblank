package mock

import (
	"bytes"
	"html/template"
)

var sent []*Email

// Email represents a sent email.
type Email struct {
	To   string
	Sub  string
	Body string
}

// Send records that an attempt was made to send an email.
func Send(to, sub, body string, data interface{}) error {
	tmpl, err := template.New("body").Parse(body)
	if err != nil {
		return err
	}

	b := bytes.Buffer{}
	err = tmpl.Execute(&b, data)
	if err != nil {
		return err
	}

	sent = append(sent, &Email{
		To:   to,
		Sub:  sub,
		Body: b.String(),
	})
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

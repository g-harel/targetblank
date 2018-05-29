package email

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const from = "\"targetblank\" <noreply@targetblank.org>"

// ISender abstracts the send action into an interface.
type ISender interface {
	Send(to, sub, body string) error
}

// Client represents an email client
type Client struct {
	ses *ses.SES
}

// NewSender creates a new email client.
func NewSender() ISender {
	return &Client{ses.New(session.New(aws.NewConfig().WithRegion("us-east-1")))}
}

// Send sends an email to a single recipient using the given content.
func (c *Client) Send(to, sub, body string) error {
	input := &ses.SendEmailInput{
		Source: aws.String(from),
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(sub),
			},
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
		},
	}

	_, err := c.ses.SendEmail(input)

	return err
}

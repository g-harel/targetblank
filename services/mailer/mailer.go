package mailer

import (
	"bytes"
	"html/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

// Must be verified by Amazon SES.
const senderEmail = "\"targetblank\" <noreply@targetblank.org>"

// Send sends an email to a single recipient using the given values.
// The body string is rendered as a template using the provided data.
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

	_, err = client.SendEmail(&ses.SendEmailInput{
		Source: aws.String(senderEmail),
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
					Data:    aws.String(b.String()),
				},
			},
		},
	})

	return err
}
